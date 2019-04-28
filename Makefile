
GO?=go

TARGET="app"

BUILD_VERSION=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date +"%Y-%m-%d %H:%M:%S")

TARGET_BUILD=$(GO) build -v -o $(TARGET) -ldflags "-w -s -X 'main.BuildVersion=${BUILD_VERSION}' -X 'main.BuildTime=${BUILD_TIME}'" main.go

.PHONY: build clean run build-docker

build: clean
	$(TARGET_BUILD)

run: build
	RUNTIME_CONFIG=./config/ ./app server

clean:
	$(GO) clean

docker-build:
	GOOS=linux GOARCH=amd64 $(TARGET_BUILD)
	docker build -t soooooul/gitbook-api:latest .

docker-push: docker-build
	docker push soooooul/gitbook-api:latest
