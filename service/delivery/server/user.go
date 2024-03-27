package handler

import (
	"dbo/assignment-test/model"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) listUser(c *gin.Context) {
	queryPage := c.Query("page")

	var (
		page  = 0
		limit = 0
		err   error
	)
	if queryPage != "" {
		page, err = strconv.Atoi(queryPage)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "page is not a number",
			})
			return
		}
	}

	queryLimit := c.Query("limit")
	if queryLimit != "" {
		limit, err = strconv.Atoi(queryLimit)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "query limit is not a number",
			})
			return
		}
	}

	search := c.Query("q")

	response, err := h.userUsecase.GetUsersPaginate(c.Request.Context(), search, page, limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) userDetail(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id is not valid",
		})
		return
	}

	response, err := h.userUsecase.GetUser(c.Request.Context(), userUUID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) searchUser(c *gin.Context) {
	search := c.Query("q")

	response, err := h.userUsecase.SearchUser(c.Request.Context(), search)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) addNewUser(c *gin.Context) {
	var body model.NewUserRequest
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid payload",
		})
		return
	}

	response, err := h.userUsecase.AddNewUser(c.Request.Context(), body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) updateUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id is not valid",
		})
		return
	}

	var body model.NewUserRequest
	err = json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid payload",
		})
		return
	}

	response, err := h.userUsecase.UpdateUser(c.Request.Context(), userUUID, body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) deleteUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id is not valid",
		})
		return
	}

	err = h.userUsecase.DeleteUser(c.Request.Context(), userUUID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
