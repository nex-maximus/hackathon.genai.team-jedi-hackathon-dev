########################################################################
 # Copyright (c) Intel Corporation 2023
 # SPDX-License-Identifier: BSD-3-Clause
########################################################################

.PHONY: build serve docker-build-bento docker-build run docker-down

# build bento service locally
build: 
	bentoml build

# deploy bento service locally
serve: 
	bentoml serve

# build bento docker image
docker-build:
ifndef BENTO_TAG
	$(error BENTO_TAG is not set, correct usage - make docker-build BENTO_TAG=<bento_image_name>:<bento_image_tag>)
else
	bentoml containerize ${BENTO_TAG}	
endif

# start bento docker service
docker-run:
ifndef BENTO_TAG
	$(error BENTO_TAG is not set, correct usage - make run BENTO_TAG=<bento_image_name>:<bento_image_tag> PROJECT_REPO_PATH=<project_repo_path>)
else ifndef PROJECT_REPO_PATH
	$(error PROJECT_REPO_PATH is not set, correct usage - make run BENTO_TAG=<bento_image_name>:<bento_image_tag> PROJECT_REPO_PATH=<project_repo_path>)
else
	docker run -itd -p 3000:3000 -v ${PROJECT_REPO_PATH}:${PROJECT_REPO_PATH} ${BENTO_TAG} serve
endif

# stop & remove bento docker service
docker-down:
ifndef BENTO_TAG
	$(error BENTO_TAG is not set, correct usage - make run BENTO_TAG=<bento_image_name>:<bento_image_tag>)
else
	docker rm $$(docker stop $$(docker ps | awk -v image=${BENTO_TAG} '$$2 == image {print $$1}'))
endif