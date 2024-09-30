package comment_service

type CreateCommentRequest struct {
	PostID  int    `json:"post_id"`
	Comment string `json:"comment"`
}
