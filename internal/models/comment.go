package models

import "time"

type Comment struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	PostID          int       `json:"post_id"`
	UserID          int       `json:"user_id"`
	ParentCommentID *int      `json:"parent_comment_id"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
	User            User      `json:"user" gorm:"foreignKey:UserID"`
	Post            Post      `json:"post" gorm:"foreignKey:PostID"`
	ParentComment   *Comment  `json:"parent_comment" gorm:"foreignKey:ParentCommentID"`
}
