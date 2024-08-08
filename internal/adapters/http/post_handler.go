package http

import (
	"github.com/gin-gonic/gin"
	"github.com/theahmadchand/go-clean-architecture/internal/entities"
	"github.com/theahmadchand/go-clean-architecture/internal/usecases"
	"net/http"
)

type PostHandler struct {
	useCase *usecases.PostUseCase
}

func NewPostHandler(useCase *usecases.PostUseCase) *PostHandler {
	return &PostHandler{useCase: useCase}
}

func (handler *PostHandler) CreatePost(context *gin.Context) {
	var post entities.Post
	if err := context.ShouldBindJSON(&post); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
	}
	if err := handler.useCase.CreatePost(&post); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create post",
			"details": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, post)
}

func (handler *PostHandler) GetPosts(context *gin.Context) {
	post, err := handler.useCase.GetPosts()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error":   "Post not found",
			"details": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, post)
}

func (handler *PostHandler) GetPost(context *gin.Context) {
	id := context.Param("id")
	post, err := handler.useCase.GetPost(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error":   "Post not found",
			"details": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, post)
}

func (handler *PostHandler) DeletePost(context *gin.Context) {
	id := context.Param("id")
	err := handler.useCase.DeletePost(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{})
}