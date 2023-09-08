/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 **********************************************************************/

package config

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
)

type Configuration struct {
	Sample string
}

func New(service interfaces.ApplicationService) (*Configuration, error) {
	config := Configuration{Sample: "hello world"}

	return &config, nil
}
