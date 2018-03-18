package model

import "github.com/jinzhu/gorm"

type Post struct {
	Guid string
	Title string
	Content string
	gorm.Model
}

func NewPost(guid, title, content string) {
	db.Create(&Post{Guid:guid, Title:title, Content:content})
}

func GetPost(guid string) *Post{
	var post Post
	db.First(&post, "guid = ?", guid)
	return &post
}

func UpdatePost(post *Post, title, content string) {
	if len(title) != 0 {
		db.Model(post).Update("title", title)
	}
	if len(content) != 0 {
		db.Model(post).Update("content", content)
	}
}

func DeletePost(post *Post) {
	db.Delete(post)
}

