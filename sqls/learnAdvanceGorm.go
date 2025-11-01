package sqls

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"name"`
	Age     uint8  `gorm:"age"`
	Gender  uint8  `gorm:"gender"`
	PostNum uint
	Posts   []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title         string
	Content       string `gorm:"type:text"`
	UserID        uint
	CommentStatus string
	CommentNum    uint
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	PostID  uint
}

func testQuery(db *gorm.DB) {
	ctx := context.Background()
	post, err := gorm.G[Post](db).
		Preload("Comments", func(db gorm.PreloadBuilder) error {
			db.Select("id", "content", "post_id")
			return nil
		}).
		Where("user_id = ?", 1).
		Select("id", "title", "content", "user_id", "comment_status", "comment_num").
		First(ctx)
	if err != nil {
		fmt.Println("query user fail", err)
		return
	}
	fmt.Println("用户1的所有文章及文章的评论：", post)

	p, err := gorm.G[Post](db).Select("id", "title", "content", "user_id").Order("comment_num desc").Limit(1).Find(ctx)
	if err != nil {
		fmt.Println("query post fail", err)
		return
	}
	fmt.Println("评论数最多的文章：", p)
}

func testHook(db *gorm.DB) {
	ctx := context.Background()

	u, err := gorm.G[User](db).Where("id = ?", 1).First(ctx)
	if err != nil {
		fmt.Println("query user fail", err)
		return
	}
	fmt.Println("新增文章前，用户1的文章数量为: ", u.PostNum)

	gorm.G[Post](db).Create(ctx, &Post{Title: "新增的文章1", Content: "新增的文章内容", UserID: 1})

	u, err = gorm.G[User](db).Where("id = ?", 1).First(ctx)
	if err != nil {
		fmt.Println("query user fail", err)
		return
	}
	fmt.Println("新增文章后，用户1的文章数量为: ", u.PostNum)

	p, err := gorm.G[Post](db).Where("id = ?", 2).First(ctx)
	if err != nil {
		fmt.Println("query post fail", err)
		return
	}
	fmt.Println("删除评论前，文章2的评论数量为: ", p.CommentNum, ", 评论状态为: ", p.CommentStatus)

	m, err := gorm.G[Comment](db).Where("id = ?", 2).First(ctx)
	if err != nil {
		fmt.Println("query comment fail", err)
		return
	}
	db.Delete(&m)

	p, err = gorm.G[Post](db).Where("id = ?", 2).Preload("Comments", nil).First(ctx)
	if err != nil {
		fmt.Println("query post fail", err)
		return
	}
	fmt.Println("删除评论后，文章2的评论数量为: ", p.CommentNum, ", 评论状态为: ", p.CommentStatus, ", 评论: ", p.Comments)
}

func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	ctx := context.Background()
	gorm.G[User](db).Where("id = ?", p.UserID).Update(ctx, "post_num", gorm.Expr("post_num + ?", 1))

	return
}

func (m *Comment) AfterDelete(db *gorm.DB) (err error) {
	ctx := context.Background()
	p, err := gorm.G[Post](db).Where("id = ?", m.PostID).First(ctx)
	if err != nil {
		fmt.Println("query Post fail", err)
		return
	}

	if p.CommentNum <= 1 {
		gorm.G[Post](db).Where("id = ?", m.PostID).
			Select("comment_status", "comment_num").
			Updates(ctx, Post{CommentStatus: "无评论", CommentNum: 0})
	} else {
		gorm.G[Post](db).Where("id = ?", m.PostID).
			Update(ctx, "comment_num", gorm.Expr("comment_num - ?", 1))
	}
	return
}

func initData(db *gorm.DB) {
	ctx := context.Background()

	db.Migrator().DropTable(&User{})
	db.Migrator().DropTable(&Post{})
	db.Migrator().DropTable(&Comment{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	users := []User{
		{
			Name:    "用户1",
			Age:     21,
			Gender:  1,
			PostNum: 2,
			Posts: []Post{
				{
					Title:         "标题1",
					Content:       "内容1",
					CommentNum:    1,
					CommentStatus: "已评论",
					Comments: []Comment{
						{
							Content: "评论1",
						},
					},
				},
				{
					Title:         "标题2",
					Content:       "内容2",
					CommentNum:    1,
					CommentStatus: "已评论",
					Comments: []Comment{
						{
							Content: "标题2的评论1",
						},
					},
				},
			},
		},
		{
			Name:    "用户2",
			Age:     22,
			Gender:  2,
			PostNum: 2,
			Posts: []Post{
				{
					Title:         "标题1",
					Content:       "内容1",
					CommentNum:    1,
					CommentStatus: "已评论",
					Comments: []Comment{
						{
							Content: "评论1",
						},
					},
				},
				{
					Title:         "标题2",
					Content:       "内容2",
					CommentNum:    3,
					CommentStatus: "已评论",
					Comments: []Comment{
						{
							Content: "标题2的评论1",
						},
						{
							Content: "标题2的评论2",
						},
						{
							Content: "标题2的评论3",
						},
					},
				},
			},
		},
	}
	gorm.G[User](db).CreateInBatches(ctx, &users, 10)
}

func TestAdvanceGorm() {
	db, err := gorm.Open(sqlite.Open("test_gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	initData(db)
	testQuery(db)
	testHook(db)
}
