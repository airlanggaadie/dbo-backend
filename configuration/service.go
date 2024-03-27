package configuration

import (
	handler "dbo/assignment-test/service/delivery/server"
	"dbo/assignment-test/service/repository/postgresql"
	"dbo/assignment-test/service/usecase"
	"fmt"
)

func (c *configuration) initService() *configuration {
	fmt.Println("setting up some features...")
	authRepository := postgresql.NewAuthRepository(c.DB)
	userRepository := postgresql.NewUserRepository(c.DB)
	orderRepository := postgresql.NewOrderRepository(c.DB)

	authUsecase := usecase.NewAuthUsecase(authRepository, userRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)
	orderUsecase := usecase.NewOrderUsecase(orderRepository)

	handler.NewHandler(c.Router, c.DB, authUsecase, userUsecase, orderUsecase)

	return c
}
