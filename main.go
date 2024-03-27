package main

import "dbo/assignment-test/configuration"

func main() {
	configuration := configuration.Init()

	configuration.Start()

	defer configuration.Stop()
}
