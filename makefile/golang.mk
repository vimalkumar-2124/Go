# Set the shell of the Makefile so that it fails if one of the parts of a pipe fails.  This is used to prevent the FIX_OUTPUT from 
SHELL = /bin/bash
.SHELLFLAGS = -o pipefail -c

# Get the module name from the gitlab repo name
GITHUB := github.com
BASE_MODULE_NAME := $(GITHUB)/$(shell echo "$(shell git remote get-url origin)" | sed -r 's/.*://g' | sed -r 's/\.git//g')
MODULE := $(shell echo $(BASE_MODULE_NAME) | sed 's/.*\///g')


# The Registry to pull the build images from and to push the result to
DOCKER_REGISTRY := vimal2124

# Get the current user and group to prevent file permission issue
USER := $(shell id -u "$$USER")
GROUP := $(shell id -g "$$USER")

# Select the appropriate GO version and backing OS
GO_VERSION := 1.19
GO_IMAGE_OS := alpine
GO_IMAGE := golang:$(GO_VERSION)-$(GO_IMAGE_OS)

# SETUP the go Direcoties variables
BASE_DIR := $(CURDIR)
OUTPUT_DIR := $(BASE_DIR)/bin
DIRS := $(OUTPUT_DIR)

# Check for the local cache file
ifneq (,$(wildcard $(HOME)/.cache/go-build))
	CACHE_DIR := $(HOME)/.cache/go-build
else
	CACHE_DIR := $(BASE_DIR)/.cache
	DIRS += $(CACHE_DIR)
endif

# If there is a local GOPATH, we can use the pkg directory from that so that we do not need to download the packages for each service
ifdef GOPATH
	PKG_DIR := $(GOPATH)/pkg
else
	PKG_DIR := $(BASE_DIR)/.pkg
	DIRS += $(PKG_DIR) 
endif

# The complement of the above directories inside the docker container so we can mount the directories for the build process
DOCKER_GOPATH := /go/src
MODULE_PATH := $(DOCKER_GOPATH)/$(BASE_MODULE_NAME)
DOCKER_PKG_DIR := /go/pkg
DOCKER_OUTPUT_DIR := /out
DOCKER_CACHE_DIR := /go/.cache
DOCKER_NETRC := /.netrc

#We need to set the NETRC in order to setup credentials for accessing gitlab.  These credentials should not be checked in
ifndef NETRC
	NETRC := "$$HOME/.netrc"
endif

# This variable allows us to run go commands.  You will want to use it using $(GO_COMMAND) <subcommand>.  For example: $(GO_COMMAND) build . 
define GO_COMMAND
docker run --rm \
	-u $(USER):$(GROUP) \
	-v $(BASE_DIR):$(MODULE_PATH) \
	-v $(PKG_DIR):$(DOCKER_PKG_DIR) \
	-v $(CACHE_DIR):$(DOCKER_CACHE_DIR) \
	-v $(OUTPUT_DIR):$(DOCKER_OUTPUT_DIR) \
	-v $(NETRC):$(DOCKER_NETRC) \
	-w $(MODULE_PATH) \
	--net host \
	-e GOCACHE=$(DOCKER_CACHE_DIR) \
	-e GOOS=linux \
	-e GOINSECURE=$(GITHUB)/* \
	-e GOPRIVATE=$(GITHUB)/* \
	-e NETRC=$(DOCKER_NETRC) \
	$(GO_IMAGE) \
	go 
endef

# This allows us to run upx without needing it installed to compress the binaries
define COMPRESS_COMMAND
	docker run \
	--rm \
	-w $(BASE_DIR) \
	-v $(BASE_DIR):$(BASE_DIR) \
	34.194.193.152:5000/upx:latest \
	--best \
	--lzma \
	-f
endef

# This is used to convert the error messages to use the local path instead of the docker path
FIX_OUTPUT := 2>&1 | sed 's\#\./\#$(BASE_DIR)/\#g'
#FIX_OUTPUT := 2>&1 | sed 's\#$(MODULE_PATH)/\#$(BASE_DIR)\#g'

