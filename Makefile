GIT_COMMIT = $(shell git rev-parse --short=7 HEAD)
GIT_VERSION ?= $(shell git describe --tags)
ifdef TAG_NAME
	ENVIRONMENT = "production"
endif
ENVIRONMENT ?= "development"
LAUNCHPAD_VERSION ?= $(or ${TAG_NAME},dev)
LD_FLAGS = -s -w -X github.com/Mirantis/mcc/version.Environment=$(ENVIRONMENT) -X github.com/Mirantis/mcc/version.GitCommit=$(GIT_COMMIT) -X github.com/Mirantis/mcc/version.Version=$(LAUNCHPAD_VERSION)
BUILD_FLAGS = -trimpath -a -tags "netgo static_build" -installsuffix netgo -ldflags "$(LD_FLAGS) -extldflags '-static'"
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
gosrc = $(wildcard *.go */*.go */*/*.go */*/*/*.go)

terraform_install_root = ~/.terraform.d/plugins/terraform.mirantis.com/mirantis/launchpad/$(GIT_VERSION)

clean:
	sudo rm -rf ./bin

builder:
	docker build -t $(BUILDER_IMAGE) -f Dockerfile.builder .

unit-test: builder
	$(GO) go test -v ./...

$(TARGET): $(gosrc)
	docker build -t $(BUILDER_IMAGE) -f Dockerfile.builder .
	GOOS=${GOOS} $(GO) go build $(BUILD_FLAGS) -o $(TARGET) main.go

build: $(TARGET)

build-all: builder
	GOOS=linux GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-linux-x64 ./cli/main.go
	GOOS=linux GOARCH=arm64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-linux-arm64 ./cli/main.go
	GOOS=windows GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-win-x64.exe ./cli/main.go
	GOOS=darwin GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-darwin-x64 ./cli/main.go
	GOOS=darwin GOARCH=arm64 $(GO) go build $(BUILD_FLAGS) -o bin/launchpad-darwin-arm64 ./cli/main.go

terraform-build-all:
	GOOS=linux GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/terraform-linux-amd64 ./terraform/main.go
	GOOS=linux GOARCH=arm64 $(GO) go build $(BUILD_FLAGS) -o bin/terraform-linux-arm64 ./terraform/main.go
	GOOS=windows GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/terraform-win-amd64.exe ./terraform/main.go
	GOOS=darwin GOARCH=amd64 $(GO) go build $(BUILD_FLAGS) -o bin/terraform-darwin-amd64 ./terraform/main.go
	GOOS=darwin GOARCH=arm64 $(GO) go build $(BUILD_FLAGS) -o bin/terraform-darwin-arm64 ./terraform/main.go

terraform-install-all: terraform-build-all
	mkdir -p $(terraform_install_root)/linux_amd64 && cp bin/terraform-linux-amd64 $(terraform_install_root)/linux_amd64/terraform-provider-launchpad
	mkdir -p $(terraform_install_root)/linux_arm64 && cp bin/terraform-linux-arm64 $(terraform_install_root)/linux_arm64/terraform-provider-launchpad
	mkdir -p $(terraform_install_root)/windows_amd64 && cp bin/terraform-win-amd64.exe $(terraform_install_root)/windows_amd64/terraform-provider-launchpad
	mkdir -p $(terraform_install_root)/darwin_amd64 && cp bin/terraform-darwin-amd64 $(terraform_install_root)/darwin_amd64/terraform-provider-launchpad
	mkdir -p $(terraform_install_root)/darwin_arm64 && cp bin/terraform-darwin-arm64 $(terraform_install_root)/darwin_arm64/terraform-provider-launchpad

release: build-all
	./release.sh

lint: builder
	$(GO) go vet ./...
	$(GO) golint -set_exit_status ./...

smoke-register-test: build
	./test/smoke_register.sh

smoke-apply-test: build
	./test/smoke_apply.sh

smoke-apply-upload-test: build
	./test/smoke_upload.sh

smoke-apply-local-repo-test: build
	./test/smoke_localrepo.sh

smoke-reset-local-repo-test: build
	./test/smoke_localrepo_reset.sh

smoke-apply-test-localhost: build
	./test/smoke_apply_local.sh

smoke-apply-bastion-test: build
	./test/smoke_apply_bastion.sh

smoke-apply-forward-test: build
	./test/smoke_apply_forward.sh

smoke-upgrade-test: build
	./test/smoke_upgrade.sh

smoke-prune-test: build
	./test/smoke_prune.sh

smoke-reset-test: build
	./test/smoke_reset.sh

smoke-cleanup:
	./test/smoke_cleanup.sh

smoke-test: smoke-apply-test
