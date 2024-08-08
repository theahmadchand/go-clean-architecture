package config

import (
	"github.com/gin-gonic/gin"
	"github.com/theahmadchand/go-clean-architecture/internal/adapters/http"
)

func SetupRoutes(engine *gin.Engine, postHandler *http.PostHandler) {
	engine.POST("/posts", postHandler.CreatePost)
	engine.GET("/posts", postHandler.GetPosts)
	engine.GET("/posts/:id", postHandler.GetPost)
	engine.DELETE("/posts/:id", postHandler.DeletePost)
}