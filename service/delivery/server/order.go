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

func (h Handler) listOrder(c *gin.Context) {
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

	search := c.Query("q")

	response, err := h.orderUsecase.GetOrdersPaginate(c.Request.Context(), search, page, limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) orderDetail(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	orderUUID, err := uuid.Parse(orderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id is not valid",
		})
		return
	}

	response, err := h.orderUsecase.GetOrder(c.Request.Context(), orderUUID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "order not found",
			})
			return
		}

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) orderSearch(c *gin.Context) {
	search := c.Query("q")

	response, err := h.orderUsecase.SearchOrder(c.Request.Context(), search)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) addNewOrder(c *gin.Context) {
	var body model.NewOrderRequest
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid payload",
		})
		return
	}

	response, err := h.orderUsecase.AddNewOrder(c.Request.Context(), body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) updateOrder(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	orderUUID, err := uuid.Parse(orderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id is not valid",
		})
		return
	}

	var body model.UpdateOrderRequest
	err = json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid payload",
		})
		return
	}

	response, err := h.orderUsecase.UpdateOrder(c.Request.Context(), orderUUID, body.BuyerName)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) deleteOrder(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	orderUUID, err := uuid.Parse(orderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id is not valid",
		})
		return
	}

	err = h.orderUsecase.DeleteOrder(c.Request.Context(), orderUUID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "order not found",
			})
			return
		}

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
