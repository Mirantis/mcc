GIT_COMMIT = $(shell git rev-parse --short=7 HEAD)
ifdef TAG_NAME
	ENVIRONMENT = "production"
endif
ENVIRONMENT ?= "development"
LAUNCHPAD_VERSION ?= $(or ${TAG_NAME},dev)
LD_FLAGS = "-w -X github.com/Mirantis/mcc/version.Environment=$(ENVIRONMENT) -X github.com/Mirantis/mcc/version.GitCommit=$(GIT_COMMIT) -X github.com/Mirantis/mcc/version.Version=$(LAUNCHPAD_VERSION)
BUILD_FLAGS = -a -tags "netgo static_build" -installsuffix netgo -ldflags $(LD_FLAGS) -extldflags '-static'"
ifeq ($(OS),Windows_NT)
       uname_s := "windows"
       TARGET ?= "bin\\launchpad.exe"
else
       uname_s := $(shell uname -s | tr '[:upper:]' '[:lower:]')
       TARGET ?= "bin/launchpad"
endif
GOOS ?= ${uname_s}
BUILDER_IMAGE = launchpad-builder
GO = docker run --rm -v "$(CURDIR)":/go/src/github.com/Mirantis/mcc \
	-w "/go/src/github.com/Mirantis/mcc" \
	-e GOPATH\
	-e GOOS \
	-e GOARCH \
	-e GOEXE \
	$(BUILDER_IMAGE)

clean:
	sudo rm -f bin/launchpad

builder:
	docker build -t $(BUILDER_IMAGE) -f Dockerfile.builder .

unit-test: builder
	$(GO) go test -v ./...

build: clean builder
	GOOS=${GOOS} $(GO) go build $(BUILD_FLAGS) -o $(TARGET) main.go

build-all: builder
	GOOS=linux GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-linux-x64 main.go
	GOOS=windows GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-win-x64.exe main.go
	GOOS=darwin GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-darwin-x64 main.go

sign-windows: build-all
	if ! which osslsigncode ; then
		sudo apt-get install -y osslsigncode
	fi
	echo -n "${WIN_PKCS12} | base64 -D > windows.pkcs12
	osslsigncode sign -pkcs12 windows.pkcs12 -pass "${WIN_PKCS12_PASSWD}" -i https://mirantis.com -n "Launchpad" -in ./bin/launchpad-win-x64.exe -out ./bin/launchpad-signed-win-x64.exe

release: build-all
	./release.sh

lint: builder
	$(GO) go vet ./...
	$(GO) golint -set_exit_status ./...

smoke-test: build
	./test/smoke.sh

smoke-upgrade-test: build
	./test/smoke_upgrade.sh

smoke-prune-test: build
	./test/smoke_prune.sh

smoke-reset-test: build
	./test/smoke_reset.sh
