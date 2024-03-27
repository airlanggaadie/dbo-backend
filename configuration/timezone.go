package configuration

import (
	"log"
	"os"
	"time"
)

func (c *configuration) initTimezone() *configuration {
	// setup the timezone
	timezone, ok := os.LookupEnv("APP_TZ")
	if !ok {
		timezone = "Asia/Jakarta"
	}

	timeLocation, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("[configuration][initTimezone] could not load timezone %s: %v", timezone, err)
	}

	time.Local = timeLocation

	return c
}
