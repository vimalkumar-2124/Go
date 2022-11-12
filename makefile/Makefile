include golang.mk

LD_FLAGS := -s -w
LD_FLAGS := $(LD_FLAGS) -extldflags=-static
BUILD_FLAGS := $(BUILD_FLAGS) -ldflags="$(LD_FLAGS)"

VERSION := 0.1
DOCKER_TAG := $(DOCKER_REGISTRY)/$(MODULE):$(VERSION)

.PHONY: all tidy docker fmt test

all: docker

docker: $(MODULE)
	docker build . -t $(DOCKER_TAG)
	

$(MODULE): test
	$(GO_COMMAND) build $(BUILD_FLAGS) -o .$(MODULE) ./ $(FIX_OUTPUT)
	$(COMPRESS_COMMAND) -o $(MODULE) .$(MODULE)

test: fmt
	$(GO_COMMAND) test $(FIX_OUTPUT)

coverage: fmt
	$(GO_COMMAND) test -coverprofile test_coverage.txt $(BASE_MODULE_NAME) $(FIX_OUTPUT)

fmt: tidy
	$(GO_COMMAND) fmt $(FIX_OUTPUT)

tidy: go.mod
	$(GO_COMMAND) mod tidy $(FIX_OUTPUT)

go.mod: | $(DIRS)
	$(GO_COMMAND) mod init $(BASE_MODULE_NAME) $(FIX_OUTPUT)

$(DIRS):
	mkdir -p $@
