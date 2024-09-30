package repo

import (
	"github.com/sachin-404/post-comment/common/database"
	"github.com/sachin-404/post-comment/internal/models"
)

type PostRepo struct{}

func NewPostRepo() *PostRepo {
	return &PostRepo{}
}

func (r *PostRepo) Create(post *models.Post) error {
	return database.DB.Create(post).Error
}

func (r *PostRepo) GetByID(id int) (*models.Post, error) {
	var post models.Post
	err := database.DB.
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepo) Delete(id int) error {
	return database.DB.Delete(&models.Post{}, id).Error
}
