package handler

import (
	"apiserver/pkg/model"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, errorResponse{message})
}

type getAllUsersResponse struct {
	Users []model.UserOutput `json:"Users"`
}

type deleteResponse struct {
	Count int64 `json:"count"`
}
