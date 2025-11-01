package sqls

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Emploee struct {
	ID         uint   `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

type Book struct {
	ID     uint    `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func testEmployeeQuery(db *sqlx.DB) {

	createEmployeeTable(db)

	cleanEmployeeData(db)

	initEmployeeData(db)

	query := `select id, name, department, salary from employees where department = ?`
	var emps []Emploee
	err := db.Select(&emps, query, "技术部")
	if err != nil {
		fmt.Println("query employees fail", err)
		return
	} else {
		fmt.Println("所有部门为 技术部 的员工信息: ", emps)
	}

	var emp Emploee

	query = `select id, name, department, salary from employees order by salary desc limit 1`
	err = db.Get(&emp, query)
	if err != nil {
		fmt.Println("query employee fail", err)
		return
	} else {
		fmt.Println("工资最高的员工信息: ", emp)
	}
}

func createEmployeeTable(db *sqlx.DB) {
	schema := `
		create table if not exists employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			department TEXT,
			salary INTEGER 
		)
	`
	_, err := db.Exec(schema)
	if err != nil {
		fmt.Println("create table employees fail, ", err)
	}
}

func cleanEmployeeData(db *sqlx.DB) {
	sql := `delete from employees`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println("delete employee data fail", err)
	}
}

func initEmployeeData(db *sqlx.DB) {
	emps := []Emploee{
		{Name: "员工1", Department: "市场部", Salary: 7000},
		{Name: "员工2", Department: "人力部", Salary: 7000},
		{Name: "员工3", Department: "技术部", Salary: 9000},
		{Name: "员工4", Department: "技术部", Salary: 10000},
		{Name: "员工5", Department: "销售部", Salary: 6000},
		{Name: "员工6", Department: "销售部", Salary: 6000},
	}

	sql := `insert into employees(name, department,salary) values (?, ?, ?)`

	for _, emp := range emps {
		_, err := db.Exec(sql, emp.Name, emp.Department, emp.Salary)
		if err != nil {
			fmt.Println("insert employee fail, emp:", emp, ", err: ", err)
		}
	}
}

func testBookQuery(db *sqlx.DB) {

	createBookTable(db)

	cleanBookData(db)

	initBookData(db)

	query := `select id, title, author, price from books where price > ?`
	var books []Book
	err := db.Select(&books, query, 50.0)
	if err != nil {
		fmt.Println("query books fail", err)
		return
	} else {
		fmt.Println("所有价格大于50的书籍信息: ", books)
	}
}

func cleanBookData(db *sqlx.DB) {
	sql := `delete from books`

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println("delete books data fail", err)
	}
}

func initBookData(db *sqlx.DB) {
	books := []Book{
		{Title: "图书1", Author: "作者1", Price: 70.0},
		{Title: "图书2", Author: "作者2", Price: 30.0},
		{Title: "图书3", Author: "作者3", Price: 50.0},
		{Title: "图书4", Author: "作者4", Price: 43.5},
		{Title: "图书5", Author: "作者5", Price: 68.5},
		{Title: "图书6", Author: "作者6", Price: 120.0},
	}

	sql := `insert into books(title, author, price) values (?, ?, ?)`

	for _, book := range books {
		_, err := db.Exec(sql, book.Title, book.Author, book.Price)
		if err != nil {
			fmt.Println("insert books fail, book:", book, ", err: ", err)
		}
	}
}

func createBookTable(db *sqlx.DB) {
	schema := `
		create table if not exists books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			author TEXT,
			price FLOAT 
		)
	`
	_, err := db.Exec(schema)
	if err != nil {
		fmt.Println("create table books fail, ", err)
	}
}

func TestSqlx() {
	db, err := sqlx.Connect("sqlite3", "test_sqlx.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	testEmployeeQuery(db)
	testBookQuery(db)
}
