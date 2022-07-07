.PHONY: build clean tool lint help dev update-pkg-cache
.PHONY: docker-build docker-compose-up docker-compose-build docker-compose-down

help:
	@echo "make tool: run specified go tool"
	@echo "make clean: remove object files and cached files"
	@echo "make build: compile packages and dependencies in local environment"
	@echo "make docker-build: build docker image in aws elastic container registry"

### Local ###
dev:
	@go run ./...

build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

clean:
	rm -rf cafe24app-backend
	go clean -i .

update-pkg-cache:
    GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)


### Docker ###
docker-compose-up:
	@docker-compose up -d
docker-compose-build:
	@docker-compose build
docker-compose-down:
	@docker-compose down --volume --remove-orphan
docker-build-aws:
	@docker build ./Dockerfile --tag=$(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_ACCOUNT_REGION).amazonaws.com/metrics-api:$(TAG)
docker-build-gcp:
	@docker build ./Dockerfile --tag=gcr.io/$(GCP_PROJECT_ID)/metrics-api:$(TAG)