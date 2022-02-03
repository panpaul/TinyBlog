PROJECT=server

GO := go

LDFLAGS += -X "main.Version=$(shell cat VERSION)"
LDFLAGS += -X "main.BuildTimeStamp=$(shell date -u '+%Y-%m-%d %I:%M:%S')"

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
