package handler

import (
	"apiserver/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllUsers(c *gin.Context)  {

}

func (h *Handler) createUser(c *gin.Context)  {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.User.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{} {
		"id": id,
	})
}

func (h *Handler) updateUser(c *gin.Context)  {

}