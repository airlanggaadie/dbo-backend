package handler

import (
	"database/sql"
	"dbo/assignment-test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine

	authUsecase  service.AuthUsecase
	userUsecase  service.UserUsecase
	orderUsecase service.OrderUsecase
}

func NewHandler(router *gin.Engine, db *sql.DB, authUsecase service.AuthUsecase, userUsecase service.UserUsecase, orderUseCase service.OrderUsecase) {
	var handler = Handler{
		router:       router,
		authUsecase:  authUsecase,
		userUsecase:  userUsecase,
		orderUsecase: orderUseCase,
	}

	handler.routes()
}

func (h Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
