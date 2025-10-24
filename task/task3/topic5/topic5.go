package topic5

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID     int
	Name   string
	Age    int
	Gender string
	Posts  []Post
	Count  int
}

type Post struct {
	ID            int
	Title         string
	Content       string
	UserID        int
	CommentStatus string
	Comments      []Comment
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("count", gorm.Expr("count + ?", 1))
	return
}

type Comment struct {
	ID      int
	Content string
	PostID  int
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
	}
	return
}

func Run(db *gorm.DB) {
	// 创建表结构
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	db.Create(&User{
		Name:   "Jan",
		Age:    18,
		Gender: "Male",
		Posts: []Post{
			{
				Title:   "My First Post",
				Content: "This is my first post.",
				Comments: []Comment{
					{
						Content: "Nice post!",
					},
					{
						Content: "Thanks for sharing!",
					},
				},
			},
			{
				Title:   "My Second Post",
				Content: "This is my second post.",
				Comments: []Comment{
					{
						Content: "Great post!",
					},
					{
						Content: "I like it!",
					},
					{
						Content: "I'm interested in this topic.",
					},
				},
			},
		},
	})

	// 查询某个用户发布的所有文章及其对应的评论信息
	var user User
	db.Debug().Model(&User{}).Preload("Posts").Preload("Posts.Comments").Where("name = ?", "Jan").Find(&user)
	fmt.Println(user)

	// 查询评论数量最多的文章信息
	var post Post
	db.Debug().Model(&Post{}).Preload("Comments").Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Select("posts.*, COUNT(comments.id) as comment_count").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).Find(&post)
	fmt.Println(post)
}
