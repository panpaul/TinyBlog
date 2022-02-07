PROJECT=server

GO := go

LDFLAGS += -X "main.Version=$(shell cat VERSION)"
LDFLAGS += -X "main.BuildTime=$(shell date -u '+%F %I:%M:%S %Z')"

GOBUILD   := $(GO) build -o $(PROJECT) -ldflags '$(LDFLAGS)'

.PHONY: all build clean run dep swagger

default: all

all: clean dep swagger build

build:
	$(GOBUILD)

clean:
	rm -rf $(PROJECT)

run: clean dep swagger build
	./$(PROJECT) run

dep:
	go get
	go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	swag init -g router/v1.go -o router/docs
