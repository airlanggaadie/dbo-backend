package configuration

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func (c *configuration) initServer() *configuration {
	fmt.Println("setting up the server...")
	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		env = "development"
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	c.portServer = port

	// setup server
	c.Router = gin.Default()
	if env == "development" {
		c.Router.Use(gin.Logger())
	}

	c.Router.Use(gin.Recovery())

	return c
}
