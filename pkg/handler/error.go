package handler

import "github.com/gin-gonic/gin"

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, error{message})
}


