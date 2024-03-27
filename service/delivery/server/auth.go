package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) login(c *gin.Context) {
	// TODO: prepare body and implement to parameters
	response, err := h.authUsecase.Login(c.Request.Context(), "", "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) loginReport(c *gin.Context) {
	// TODO: prepare query param and implement to parameters
	response, err := h.authUsecase.Report(c.Request.Context(), 0, 0)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
