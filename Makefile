
GO?=go

TARGET="gitbook-api"

BUILD_VERSION=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date +"%Y-%m-%d %H:%M:%S")

DockerTag?="latest"

TARGET_BUILD=$(GO) build -v -o $(TARGET) -ldflags "-w -s -X 'github.com/SoulSu/gitbook-api/cmd.BuildVersion=${BUILD_VERSION}' -X 'github.com/SoulSu/gitbook-api/cmd.BuildTime=${BUILD_TIME}'" main.go

.PHONY: build clean run build-docker test

build: clean
	$(TARGET_BUILD)

run: build
	RUNTIME_CONFIG=./config/ ./app server

test:
	go test ./...

clean:
	$(GO) clean

docker-build:
#	GOOS=linux GOARCH=amd64 $(TARGET_BUILD)
	docker build -t soooooul/gitbook-api:$(DockerTag) .

docker-push: docker-build
	docker push soooooul/gitbook-api:$(DockerTag)
