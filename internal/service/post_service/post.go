package post_service

import (
	"github.com/sachin-404/post-comment/internal/models"
	"github.com/sachin-404/post-comment/internal/repo"
)

type PostService struct {
	postRepo *repo.PostRepo
}

func NewPostService(postRepo *repo.PostRepo) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) CreatePost(req *CreatePostRequest, userID int) error {
	post := models.NewPost(req.Title, req.Content, userID)
	return s.postRepo.Create(post)
}

func (s *PostService) GetPostByID(id int) (*models.Post, error) {
	return s.postRepo.GetByID(id)
}

func (s *PostService) DeletePost(id int) error {
	return s.postRepo.Delete(id)
}
