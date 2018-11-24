package main

import (
	"github.com/jackson6/gcic-service/service-client/app"
	"github.com/jackson6/gcic-service/service-client/client"
	"github.com/jackson6/gcic-service/service-client/config"
)

func main() {
	appConfig := config.GetConfig()

	dcApp := &app.App{}
	dcClient := &client.Client{}
	dcApp.Initialize(appConfig, dcClient)
	dcApp.Run(":3000")
}