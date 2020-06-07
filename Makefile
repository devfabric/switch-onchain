export WORKSPACE = $(shell pwd)
export WORKDIR=$(WORKSPACE)
export GO111MODULE=on
#version import dir
VERSION_DIR     := code.uni-ledger.com/switch/resource-access-control/internal/public

BUILD_VERSION   = $(shell git describe --abbrev=0 --tags)
# BUILD_VERSION  := v1.1.1 

BUILD_BRANCH    := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMITID  := $(shell git log --pretty=format:"%h" -1 )
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME      := resource-access-control


all: build

clean:
	rm -rf ./resource-access-control
	@echo "Done clean"

build:
	CGO_ENABLED=1 go build -ldflags \
    " \
    -X '${VERSION_DIR}.buildVersion=${BUILD_VERSION}' \
    -X '${VERSION_DIR}.buildName=${BUILD_NAME}' \
    -X '${VERSION_DIR}.buildBranch=${BUILD_BRANCH}' \
    -X '${VERSION_DIR}.buildCommitID=${BUILD_COMMITID}' \
    -X '${VERSION_DIR}.buildTime=${BUILD_TIME}' \
    " \
    -o $(WORKSPACE)/${BUILD_NAME} main.go
	@echo "Done build"
    #-linkmode external -extldflags -static-all

prod:
	CGO_ENABLED=1 go build -ldflags \
    " \
    -X '${VERSION_DIR}.buildVersion=${BUILD_VERSION}' \
    -X '${VERSION_DIR}.buildName=${BUILD_NAME}' \
    -X '${VERSION_DIR}.buildBranch=${BUILD_BRANCH}' \
    -X '${VERSION_DIR}.buildCommitID=${BUILD_COMMITID}' \
    -X '${VERSION_DIR}.buildTime=${BUILD_TIME}' \
    " \
	-mod=vendor \
    -o $(WORKSPACE)/${BUILD_NAME} main.go
	@echo "Done prod build"

run:
	$(WORKSPACE)/${BUILD_NAME}

.PHONY: clean build
