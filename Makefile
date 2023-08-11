APP_NAME :=go-schedule-example
DOCKER_REPO=localhost:5000
# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# DOCKER TASKS
# Build the container
build-docker: ## Build the container
	docker build -t $(APP_NAME) .
clear-none:
	docker rmi -f `docker images -a |grep 'none'|awk '{print \$$3}'`

clear:
ifeq ($(HELM),)
	echo 'not exist broker'
else
	--helm delete $(APP_NAME) --purge
endif
	--docker rmi -f `docker images -a |grep '$(APP_NAME)'|awk '{print\$$3}'`
	--docker rmi -f `docker images -f "dangling=true" -q `

clearall:
	docker rmi -f `docker images -a |grep '$(APP_NAME)\|none'|awk '{print\$$3}'`

clean: stop clear

build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .

run: ## Run container
	docker run \
	--name $(APP_NAME) -p 1323:1323  -d $(APP_NAME)

stop: ## Stop and remove a running container
	docker stop $(APP_NAME); docker rm -f $(APP_NAME)

release-local: build-nc publish-local clear

# Docker publish
publish: login publish-latest publish-version ## Publish the `{version}` ans `latest` tagged containers to ECR

login: ## Login dockerhub
	--aws ecr get-login-password --region ap-southeast-1 --profile=onsky | docker login --username AWS --password-stdin $(DOCKER_REPO)	
publish-latest: tag-latest ## Publish the `latest` taged container to ECR
	--aws ecr create-repository --repository-name $(APP_NAME) --region ap-southeast-1 --profile=onsky
	--aws ecr batch-delete-image --repository-name $(APP_NAME) --image-ids imageTag=latest --region ap-southeast-1 --profile=onsky
	@echo 'publish latest to $(DOCKER_REPO)'
	AWS_DEFAULT_REGION=ap-southeast-1 AWS_PROFILE=onsky docker push $(DOCKER_REPO)/$(APP_NAME):latest

publish-version: tag-version ## Publish the `{version}` taged container to ECR
	@echo 'publish $(VERSION) to $(DOCKER_REPO)'
	AWS_DEFAULT_REGION=ap-southeast-1 AWS_PROFILE=onsky docker push $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

# Docker publish local
publish-local: tag-latest publish-latest-local

publish-latest-local:
	docker push $(DOCKER_REPO)/$(APP_NAME):latest

# Docker tagging
tag: tag-latest tag-version ## Generate container tags for the `{version}` ans `latest` tags

tag-latest: ## Generate container `{version}` tag
	@echo 'create tag latest'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):latest

tag-version: ## Generate container `latest` tag
	@echo 'create tag $(VERSION)'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

run-go: ## Run service
	go run cmd/api/main.go

build-go: ## Build go
	GOOS=linux CGO_ENABLED=0  GOARCH=amd64 go build -o $(APP_NAME) cmd/api/main.go

deploy-dev:
	helm install $(APP_NAME) deployment --set ENV=app \
	--set image.repository=$(DOCKER_REPO)/$(APP_NAME) \
	--set fullnameOverride=$(APP_NAME)
undeloy-dev:
	helm uninstall $(APP_NAME)