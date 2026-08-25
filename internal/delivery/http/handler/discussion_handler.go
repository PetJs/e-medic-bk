// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/discussion"
)

// DiscussionHandler handles discussion board (posts + comments) requests.
type DiscussionHandler struct {
	createPostUC    *discussion.CreatePostUseCase
	listPostsUC     *discussion.ListPostsUseCase
	getPostUC       *discussion.GetPostUseCase
	deletePostUC    *discussion.DeletePostUseCase
	pinPostUC       *discussion.PinPostUseCase
	unpinPostUC     *discussion.UnpinPostUseCase
	createCommentUC *discussion.CreateCommentUseCase
	listCommentsUC  *discussion.ListCommentsUseCase
	deleteCommentUC *discussion.DeleteCommentUseCase
}

// NewDiscussionHandler creates a new DiscussionHandler.
func NewDiscussionHandler(
	createPostUC *discussion.CreatePostUseCase,
	listPostsUC *discussion.ListPostsUseCase,
	getPostUC *discussion.GetPostUseCase,
	deletePostUC *discussion.DeletePostUseCase,
	pinPostUC *discussion.PinPostUseCase,
	unpinPostUC *discussion.UnpinPostUseCase,
	createCommentUC *discussion.CreateCommentUseCase,
	listCommentsUC *discussion.ListCommentsUseCase,
	deleteCommentUC *discussion.DeleteCommentUseCase,
) *DiscussionHandler {
	return &DiscussionHandler{
		createPostUC:    createPostUC,
		listPostsUC:     listPostsUC,
		getPostUC:       getPostUC,
		deletePostUC:    deletePostUC,
		pinPostUC:       pinPostUC,
		unpinPostUC:     unpinPostUC,
		createCommentUC: createCommentUC,
		listCommentsUC:  listCommentsUC,
		deleteCommentUC: deleteCommentUC,
	}
}

// CreatePost creates a discussion post under a module.
func (h *DiscussionHandler) CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userID := c.GetString("user_id")
	response, err := h.createPostUC.Execute(c.Request.Context(), c.Param("id"), userID, &req)
	if err != nil {
		if errors.Is(err, discussion.ErrModuleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		}
		internalError(c, "Failed to create post", err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// ListPosts lists a module's discussion posts (?page=&limit=).
func (h *DiscussionHandler) ListPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	response, err := h.listPostsUC.Execute(c.Request.Context(), c.Param("id"), &dto.ListPostsRequest{Page: page, Limit: limit})
	if err != nil {
		if errors.Is(err, discussion.ErrModuleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		}
		internalError(c, "Failed to list posts", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetPost returns a single discussion post.
func (h *DiscussionHandler) GetPost(c *gin.Context) {
	response, err := h.getPostUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, discussion.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		internalError(c, "Failed to get post", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeletePost deletes a post (its own author, or an admin moderating).
func (h *DiscussionHandler) DeletePost(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")
	if err := h.deletePostUC.Execute(c.Request.Context(), userID, role, c.Param("id")); err != nil {
		if errors.Is(err, discussion.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		internalError(c, "Failed to delete post", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

// PinPost pins a post to the top of its module's board (admin).
func (h *DiscussionHandler) PinPost(c *gin.Context) {
	response, err := h.pinPostUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, discussion.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		internalError(c, "Failed to pin post", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// UnpinPost unpins a post (admin).
func (h *DiscussionHandler) UnpinPost(c *gin.Context) {
	response, err := h.unpinPostUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, discussion.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		internalError(c, "Failed to unpin post", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateComment creates a comment (or threaded reply) on a post.
func (h *DiscussionHandler) CreateComment(c *gin.Context) {
	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userID := c.GetString("user_id")
	response, err := h.createCommentUC.Execute(c.Request.Context(), c.Param("id"), userID, &req)
	if err != nil {
		if errors.Is(err, discussion.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		if errors.Is(err, discussion.ErrParentCommentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Parent comment not found"})
			return
		}
		internalError(c, "Failed to create comment", err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// ListComments lists every comment on a post.
func (h *DiscussionHandler) ListComments(c *gin.Context) {
	response, err := h.listCommentsUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, discussion.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		internalError(c, "Failed to list comments", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteComment deletes a comment (its own author, or an admin moderating).
func (h *DiscussionHandler) DeleteComment(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")
	if err := h.deleteCommentUC.Execute(c.Request.Context(), userID, role, c.Param("id")); err != nil {
		if errors.Is(err, discussion.ErrCommentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}
		internalError(c, "Failed to delete comment", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
