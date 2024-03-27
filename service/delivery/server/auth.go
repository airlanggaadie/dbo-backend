package handler

import (
	"dbo/assignment-test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) login(c *gin.Context) {
	var body model.LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.authUsecase.Login(c.Request.Context(), body.Username, body.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) loginReport(c *gin.Context) {
	var (
		page  = 0
		limit = 0
		err   error
	)

	queryPage := c.Query("page")
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

	response, err := h.authUsecase.Report(c.Request.Context(), page, limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
