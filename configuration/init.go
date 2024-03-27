package configuration

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type configuration struct {
	Router     *gin.Engine
	Server     *http.Server
	portServer string
	DB         *sql.DB
}

func Init() *configuration {
	var configuration configuration

	return configuration.
		initTimezone().
		initPostgreSql().
		migrate().
		initServer().
		initService()
}

func (c *configuration) Start() {
	fmt.Println("starting apps...")
	go func() {
		defer func() {
			if err, ok := recover().(error); ok && err != nil {
				log.Printf("[configuration][Start] recover error: %v\n", err)
			}
		}()

		c.Server = &http.Server{
			Handler: c.Router,
			Addr:    ":" + c.portServer,
		}

		if err := c.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[configuration][Start] shutting down the server: %v", err)
		}
	}()
}

func (c *configuration) Stop() {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("stopping apps...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.Server.Shutdown(ctx); err != nil {
		log.Fatalf("[configuration][Stop] shutting down serverr %v\n", err)
	}

	<-ctx.Done()
	log.Println("[configuration][Stop] timeout of 10 seconds.")

	log.Println("Server exiting")
}
