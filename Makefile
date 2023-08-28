########################################################################
 # Copyright (c) Intel Corporation 2023
########################################################################

.PHONY: build tidy test clean docker
GO=CGO_ENABLED=1 go

# VERSION file is not needed for local development, In the CI/CD pipeline, a temporary VERSION file is written
# if you need a specific version, just override below
MSVERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)

# This pulls the version of the SDK from the go.mod file. If the SDK is the only required module,
# it must first remove the word 'required' so the offset of $2 is the same if there are multiple required modules
SDKVERSION=$(shell cat ./go.mod | grep 'github.com/edgexfoundry/app-functions-sdk-go/v3 v' | sed 's/require//g' | awk '{print $$2}')

PROJECT=hackathon

MICROSERVICES=
DOCKERS=docker-grafana

GIT_SHA=$(shell git rev-parse HEAD)

define COMPOSE_DOWN
	docker compose -p edgex -f docker-compose-edgex.yml -f docker-compose-apps.yml down $1
	docker compose -p monitor -f docker-compose-monitor.yml down $1
endef

tidy:
	go mod tidy

# NOTE: This is only used for local development. Jenkins CI does not use this make target
docker: ${DOCKERS}

docker-grafana:
	$(MAKE) -C grafana docker-grafana

test:
	$(GO) test -coverprofile=coverage.out `go list ./... | grep -v integration-tests`
	$(GO) vet ./...
	gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")
	[ "`gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")`" = "" ]
	#./bin/test-attribution-txt.sh

run-portainer:
	docker compose -p portainer -f docker-compose-portainer.yml up -d

run:
	docker compose -p edgex \
		-f docker-compose-edgex.yml \
		-f docker-compose-apps.yml \
		up -d


run-monitor: docker-grafana
	docker compose -p monitor \
		-f docker-compose-monitor.yml \
		up -d

down-portainer:
	docker compose -p portainer -f docker-compose-portainer.yml down

down:
	$(COMPOSE_DOWN)

down-clean:
	$(call COMPOSE_DOWN,-v)


clean-volumes:
	docker volume prune -f --filter all=true
