/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 **********************************************************************/

package main

import (
	"app-sample-service/config"
	"app-sample-service/controller"
	"app-sample-service/functions"
	"os"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
)

// This application service simulates an ML pipeline platform and is used by Geti and BentoML pipelines.
// There are no unit tests since it is temporary.
func main() {

	service, ok := pkg.NewAppService("app-sample-service")
	if !ok {
		os.Exit(-1)
	}
	// Leverage the built-in logging service in EdgeX
	lc := service.LoggingClient()

	configuration, configErr := config.New(service)
	if configErr != nil {
		lc.Errorf("failed to retrieve read app settings from configuration: %s", configErr.Error())
		os.Exit(-1)
	}

	var err error

	justFilePipeline := functions.New(configuration.Sample)
	err = service.AddFunctionsPipelineForTopics("helloevent", []string{"functionpipeline"},
		justFilePipeline.ProcessEvent,
	)

	if err != nil {
		lc.Error(err.Error())
		os.Exit(-1)
	}

	lc.Info("Functions Pipeline set...")

	appController := controller.New(lc, configuration.Sample)
	if err := appController.RegisterRoutes(service); err != nil {
		lc.Errorf("RegisterRoutes returned error: %s", err.Error())
		os.Exit(-1)
	}

	err = service.Run()
	if err != nil {
		lc.Errorf("Run returned error: %s", err.Error())
		os.Exit(-1)
	}

	// Do any required cleanup here
	os.Exit(0)
}
