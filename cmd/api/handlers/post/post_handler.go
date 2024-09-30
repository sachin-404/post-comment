package post

import (
	"github.com/labstack/echo/v4"
	"github.com/sachin-404/post-comment/common/auth"
	"github.com/sachin-404/post-comment/internal/service/post_service"
	"net/http"
	"strconv"
)

type PostHandler struct {
	postService *post_service.PostService
}

func NewPostHandler(postService *post_service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unable to get user ID: "+err.Error())
	}
	var req post_service.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := h.postService.CreatePost(&req, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Post created successfully"})
}

func (h *PostHandler) GetPost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	post, err := h.postService.GetPostByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
	}

	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	if err := h.postService.DeletePost(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted successfully"})
}
