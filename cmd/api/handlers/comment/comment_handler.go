package comment

import (
	"github.com/labstack/echo/v4"
	"github.com/sachin-404/post-comment/common/auth"
	"github.com/sachin-404/post-comment/internal/service/comment_service"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	commentService *comment_service.CommentService
}

func NewCommentHandler(commentService *comment_service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unable to get user ID: "+err.Error())
	}

	var req comment_service.CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.commentService.CreateComment(&req, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Comment created successfully"})
}

func (h *CommentHandler) GetComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
	}

	comment, err := h.commentService.GetComment(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unable to get user ID: "+err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid comment ID"})
	}

	if err := h.commentService.DeleteComment(id, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted successfully"})
}
