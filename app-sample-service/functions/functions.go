/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 **********************************************************************/

package functions

import (
	"fmt"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
)

const ResourceName = "PipelineParameters"

type Pipeline struct {
	stuff string
	lc    logger.LoggingClient
}

func New(sample string) Pipeline {
	return Pipeline{
		stuff: sample,
	}
}

// ProcessEvent is the entry point App Pipeline Function that receives and processes the EdgeX Event/Reading that
// contains the Pipeline Parameters
func (p *Pipeline) ProcessEvent(ctx interfaces.AppFunctionContext, data interface{}) (bool, interface{}) {
	p.lc = ctx.LoggingClient()
	p.lc.Debugf("Running ProcessEvent...")

	if data == nil {
		err := fmt.Errorf("no data received")
		p.lc.Errorf("ProcessEvent failed: %s", err.Error())
		return false, err // Terminate since can not send back status w/o knowing the URL that is passed in the data
	}

	event, ok := data.(dtos.Event)
	if !ok {
		err := fmt.Errorf("type received is not an Event")
		p.lc.Errorf("ProcessEvent failed: %s", err.Error())
		return false, err // Terminate since can not send back status w/o knowing the URL that is passed in the data
	}
	p.lc.Debugf("current Event: %v", event)

	return true, data // All is good, this indicates success for the next function
}
