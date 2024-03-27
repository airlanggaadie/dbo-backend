package handler

import (
	"dbo/assignment-test/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) listUser(c *gin.Context) {
	// TODO: prepare query param and implement to parameters
	response, err := h.userUsecase.GetUsersPaginate(c.Request.Context(), "", 0, 0)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) userDetail(c *gin.Context) {
	// TODO: prepare url param
	response, err := h.userUsecase.GetUser(c.Request.Context(), uuid.New())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) searchUser(c *gin.Context) {
	// TODO: prepare query param and implement to parameters
	response, err := h.userUsecase.SearchUser(c.Request.Context(), "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) addNewUser(c *gin.Context) {
	// TODO: prepare body and implement to parameters
	response, err := h.userUsecase.AddNewUser(c.Request.Context(), model.NewUserRequest{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) updateUser(c *gin.Context) {
	// TODO: prepare url param and body then implement to parameters
	response, err := h.userUsecase.UpdateUser(c.Request.Context(), uuid.New(), model.NewUserRequest{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) deleteUser(c *gin.Context) {
	// TODO: prepare url param then implement to parameters
	err := h.userUsecase.DeleteUser(c.Request.Context(), uuid.New())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
