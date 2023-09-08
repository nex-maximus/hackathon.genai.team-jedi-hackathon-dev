########################################################################
 # Copyright (c) Intel Corporation 2023
########################################################################

.PHONY: build tidy test clean docker

GO=CGO_ENABLED=1 go

PROJECT=hackathon
MICROSERVICE=app-sample-service
GOFLAGS=-ldflags "-X github.com/edgexfoundry/app-functions-sdk-go/internal.SDKVersion=$(SDKVERSION) -X github.com/edgexfoundry/app-functions-sdk-go/internal.ApplicationVersion=$(APPVERSION)"

tidy:
	go mod tidy

build: tidy
	$(GO) build $(GOFLAGS) -o $(MICROSERVICE)

docker: clean tidy
	docker build \
		--build-arg http_proxy \
		--build-arg https_proxy \
		-f Dockerfile \
		-t $(PROJECT)/$(MICROSERVICE):app-$(app_version) \
		-t $(PROJECT)/$(MICROSERVICE):dev \
		.

test:
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) vet ./...
	gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")
	[ "`gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")`" = "" ]

clean:
	rm -f $(MICROSERVICE)