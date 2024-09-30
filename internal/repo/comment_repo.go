package repo

import (
	"github.com/sachin-404/post-comment/common/database"
	"github.com/sachin-404/post-comment/internal/models"
)

type CommentRepo struct{}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{}
}

func (r *CommentRepo) Create(comment *models.Comment) error {
	return database.DB.Create(comment).Error
}

func (r *CommentRepo) GetByID(id int) (*models.Comment, error) {
	var comment models.Comment
	err := database.DB.
		Preload("User").
		First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepo) Delete(id int) error {
	return database.DB.Delete(&models.Comment{}, id).Error
}
