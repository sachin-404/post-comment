package models

import "time"

type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

func NewPost(title, content string, userID int) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}
