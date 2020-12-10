package handler

import (
	"github.com/gin-gonic/gin"
)

// New function initiates a gin router and sets the routes
func New() *gin.Engine {
	r := gin.Default()

	setSignatureRoutes(r)

	return r
}
