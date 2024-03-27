package configuration

import (
	handler "dbo/assignment-test/service/delivery/server"
	"dbo/assignment-test/service/repository/jwt"
	"dbo/assignment-test/service/repository/postgresql"
	"dbo/assignment-test/service/usecase"
	"fmt"
	"os"
)

func (c *configuration) initService() *configuration {
	fmt.Println("setting up some features...")

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		jwtSecret = "supersecret"
	}

	jwtIssuer, ok := os.LookupEnv("JWT_ISSUER")
	if !ok {
		jwtIssuer = "dbo"
	}

	userRepository := postgresql.NewUserRepository(c.DB)
	orderRepository := postgresql.NewOrderRepository(c.DB)
	jwtRepository := jwt.NewJwtRepository(jwtSecret, jwtIssuer)

	authUsecase := usecase.NewAuthUsecase(userRepository, jwtRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)
	orderUsecase := usecase.NewOrderUsecase(orderRepository)

	handler.NewHandler(c.Router, c.DB, authUsecase, userUsecase, orderUsecase, jwtRepository)

	return c
}
