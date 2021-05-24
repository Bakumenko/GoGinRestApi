package handler

import (
	"apiserver/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
		log.Fatal("error occured while binding user in create: %s", err.Error())
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

func (h *Handler) getOneUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.services.User.GetOneUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatal("error occured while getting users: %s", err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	id := c.Param("id")

	var input model.UpdateUserInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		log.Fatal("error occured while binding user in update: %s", err.Error())
	}

	user, err := h.services.User.UpdateUser(input, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatal("error occured while updating user: %s", err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) deleteUser(c *gin.Context) {
	id := c.Param("id")

	count, err := h.services.User.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatal("error occured while deleting: %s", err.Error())
	}

	c.JSON(http.StatusOK, deleteResponse{
		Count: count,
	})
}
