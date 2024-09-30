package models

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}

func NewComment(postID, userID int, content string) *Comment {
	return &Comment{
		PostID:    postID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
