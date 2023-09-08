.PHONY: build clean

GO=CGO_ENABLED=1 go

APPVERSION=2.0.0

# This pulls the version of the SDK from the go.mod file. If the SDK is the only required module,
# it must first remove the word 'required' so the offset of $2 is the same if there are multiple required modules
SDKVERSION=$(shell cat ./go.mod | grep 'github.com/edgexfoundry/app-functions-sdk-go v' | sed 's/require//g' | awk '{print $$2}')

PROJECT=hackathon
MICROSERVICE=app-data-export
GOFLAGS=-ldflags "-X github.com/edgexfoundry/app-functions-sdk-go/internal.SDKVersion=$(SDKVERSION) -X github.com/edgexfoundry/app-functions-sdk-go/internal.ApplicationVersion=$(APPVERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

tidy:
	go mod tidy

build: tidy
	$(GO) build $(GOFLAGS) -o $(MICROSERVICE)

clean:
	rm -f $(MICROSERVICE)

# NOTE: This is only used for local development. Jenkins CI does not use this make target
docker: clean tidy
	docker build \
	    --build-arg http_proxy \
	    --build-arg https_proxy \
		-f Dockerfile \
		-t $(PROJECT)/$(MICROSERVICE):app-$(app_version) \
		-t $(PROJECT)/$(MICROSERVICE):dev \
		.

docker-data-organizer:
	docker build \
	    --build-arg http_proxy \
	    --build-arg https_proxy \
		-f ${ORGANIZER}/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t ${PROJECT}/${ORGANIZER}:$(GIT_SHA) \
		-t ${PROJECT}/${ORGANIZER}:${MSVERSION}-dev \
		.