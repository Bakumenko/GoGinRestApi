package handler

import (
	"apiserver/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createRole(c *gin.Context) {
	var input model.Role

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		//log.Fatal("error occured while binding role in create: %s", err.Error())
		return
	}

	id, err := h.services.Role.CreateRole(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		//log.Fatal("error occured while creating role: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
