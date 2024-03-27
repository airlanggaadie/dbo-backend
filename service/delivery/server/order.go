package handler

import (
	"dbo/assignment-test/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) listOrder(c *gin.Context) {
	// TODO: prepare query param and implement to parameters
	response, err := h.orderUsecase.GetOrdersPaginate(c.Request.Context(), "", 0, 0)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) orderDetail(c *gin.Context) {
	// TODO: prepare url param
	response, err := h.orderUsecase.GetOrder(c.Request.Context(), uuid.New())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) orderSearch(c *gin.Context) {
	// TODO: prepare query param and implement to parameters
	response, err := h.orderUsecase.SearchOrder(c.Request.Context(), "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) addNewOrder(c *gin.Context) {
	// TODO: prepare body and implement to parameters
	response, err := h.orderUsecase.AddNewOrder(c.Request.Context(), model.NewOrderRequest{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) updateOrder(c *gin.Context) {
	// TODO: prepare url param and body then implement to parameters
	response, err := h.orderUsecase.UpdateOrder(c.Request.Context(), uuid.New(), model.NewOrderRequest{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) deleteOrder(c *gin.Context) {
	// TODO: prepare url param then implement to parameters
	err := h.orderUsecase.DeleteOrder(c.Request.Context(), uuid.New())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
