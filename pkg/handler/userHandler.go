package handler

import (
	"apiserver/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type getAllUsersResponse struct {
	Data []model.User `json:"Users"`
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		log.Fatal("error occured while getting users: %s", err.Error())
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}

func (h *Handler) createUser(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		log.Fatal("error occured while binding user: %s", err.Error())
	}

	id, err := h.services.User.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatal("error occured while creating user: %s", err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateUser(c *gin.Context) {

}
