/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 **********************************************************************/

package controller

import (
	"fmt"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"net/http"
)

type SampleController struct {
	lc     logger.LoggingClient
	sample string
}

func New(lc logger.LoggingClient, sample string) *SampleController {
	return &SampleController{
		lc:     lc,
		sample: sample,
	}
}

func (p *SampleController) RegisterRoutes(service interfaces.ApplicationService) error {
	if err := service.AddRoute("/api/v3/hello", p.getPipelinesHandler, http.MethodGet); err != nil {
		return fmt.Errorf("could not register routes: %s", err.Error())
	}

	p.lc.Info("Routes added...")
	return nil
}

func (p *SampleController) getPipelinesHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("hello"))
}
