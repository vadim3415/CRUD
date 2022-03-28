package handler

import (
	"HTTP31/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createNewUser(c *gin.Context) {
	var inputRequest *model.User

	if err := c.BindJSON(&inputRequest); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid inputRequest body")
		return
	}
	id, err := h.services.CreateUser(inputRequest)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"use created, id:": id,
	})
}

func (h *Handler) makeFriends(c *gin.Context) {
	var inputRequest *model.User

	if err := c.BindJSON(&inputRequest); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid inputRequest body")
		return
	}
	friends, err := h.services.MakeFriends(inputRequest)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, friends)
}

type getAllUsers struct {
	Data []model.UserGet `json:"data"`
}

func (h *Handler) getAllUser(c *gin.Context) {
	user, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllUsers{
		Data: user,
	})
}
func (h *Handler) getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	friends, err := h.services.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, friends)
}

func (h *Handler) updateUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input model.UserUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input.NewAge)

	UpdateId, err := h.services.UpdateUser(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"use update, id": UpdateId,
	})
}

func (h *Handler) deleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	deleteId, err := h.services.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"use delete, id": deleteId,
	})
}

