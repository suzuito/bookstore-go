package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter(root *gin.Engine, app Application) {
	root.GET("/status", GetStatus(app))
	root.GET("/books", GetBooks(app))
	root.GET("/books/:id", GetBooksByID(app))
}
