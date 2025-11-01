package sqls

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/**
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   uint8
	Grade string
}

func testSqlPhase(db *gorm.DB) {
	db.Where("1 = 1").Delete(&Student{})

	for i := 13; i < 22; i++ {
		stu := &Student{Name: "学生" + strconv.Itoa(i), Age: uint8(i), Grade: "三年级"}
		db.Create(stu)
	}

	stu := &Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Create(stu)

	s := []Student{}
	db.Where("age > ?", 18).Find(&s)
	fmt.Println(s)

	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	db.Find(&s)
	fmt.Println(s)

	db.Where("age < ?", 15).Delete(&s)
	db.Find(&s)
	fmt.Println(s)
}

func testSqlPhase2(db *gorm.DB) {
	ctx := context.Background()

	stuDb := gorm.G[Student](db)

	stuDb.Where("1 = 1").Delete(ctx)

	for i := 13; i < 22; i++ {
		stu := &Student{Name: "学生" + strconv.Itoa(i), Age: uint8(i), Grade: "三年级"}
		stuDb.Create(ctx, stu)
	}

	stu := &Student{Name: "张三", Age: 20, Grade: "三年级"}
	stuDb.Create(ctx, stu)

	s, _ := stuDb.Where("age > ?", 18).Find(ctx)
	fmt.Println(s)

	stuDb.Where("name = ?", "张三").Update(ctx, "grade", "四年级")

	s, _ = stuDb.Find(ctx)
	fmt.Println(s)

	stuDb.Where("age < ?", 15).Delete(ctx)
	s, _ = stuDb.Find(ctx)
	fmt.Println(s)
}

/*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）
和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance int
}

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountId uint
	ToAccountId   uint
	Amount        int
}

func testTrans(db *gorm.DB) {
	ctx := context.Background()

	tran := gorm.G[Transaction](db)
	acc := gorm.G[Account](db)

	tran.Where("1 = 1").Delete(ctx)
	acc.Where("1 = 1").Delete(ctx)

	acc.Create(ctx, &Account{ID: 1, Balance: 150})
	acc.Create(ctx, &Account{ID: 2, Balance: 50})

	err := db.Transaction(func(tx *gorm.DB) error {
		a, _ := gorm.G[Account](tx).Where("id = ?", 1).First(ctx)
		fmt.Println("account A: ", a)
		if a.Balance < 100 {
			return errors.New("余额不足")
		}
		b, _ := gorm.G[Account](tx).Where("id = ?", 2).First(ctx)
		fmt.Println("account B: ", b)

		b.Balance += 100
		a.Balance -= 100

		gorm.G[Account](tx).Where("id = ?", 1).Update(ctx, "balance", a.Balance)
		gorm.G[Account](tx).Where("id = ?", 2).Update(ctx, "balance", b.Balance)

		a, _ = gorm.G[Account](tx).Where("id = ?", 1).First(ctx)
		b, _ = gorm.G[Account](tx).Where("id = ?", 2).First(ctx)
		fmt.Println("account A after trans: ", a)
		fmt.Println("account B after trans: ", b)

		gorm.G[Transaction](tx).Create(ctx, &Transaction{FromAccountId: a.ID, ToAccountId: b.ID, Amount: 100})

		t, _ := gorm.G[Transaction](tx).Where("from_account_id = ? and to_account_id = ?", 1, 2).First(ctx)
		fmt.Println("trans: ", t)

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		a, _ := gorm.G[Account](tx).Where("id = ?", 1).First(ctx)
		fmt.Println("account A: ", a)
		if a.Balance < 100 {
			return errors.New("余额不足")
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestGorm() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	testSqlPhase(db)
	testSqlPhase2(db)
	testTrans(db)
}
