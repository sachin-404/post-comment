package comment_service

import (
	"errors"
	"github.com/sachin-404/post-comment/internal/models"
	"github.com/sachin-404/post-comment/internal/repo"
)

type CommentService struct {
	commentRepo *repo.CommentRepo
	postRepo    *repo.PostRepo
}

func NewCommentService(commentRepo *repo.CommentRepo, postRepo *repo.PostRepo) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func (s *CommentService) CreateComment(req *CreateCommentRequest, userID int) error {
	existingPost, err := s.postRepo.GetByID(req.PostID)
	if err != nil {
		return err
	}
	if existingPost == nil {
		return errors.New("post does not exist")
	}
	comment := models.NewComment(req.PostID, userID, req.Comment)
	return s.commentRepo.Create(comment)
}

func (s *CommentService) GetComment(id int) (*models.Comment, error) {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, errors.New("comment not found")
	}
	return comment, nil
}

func (s *CommentService) DeleteComment(id, userId int) error {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return err
	}

	if comment.UserID != userId {
		return errors.New("user is not authorized to delete this comment")
	}

	return s.commentRepo.Delete(id)
}
