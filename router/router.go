package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter(root *gin.Engine, app Application) {
	root.GET("/books/:id", GetBooksByID(app))
}
