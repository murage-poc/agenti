BUILD_OUTPUT_FILENAME ?= agenti
DOCKER_TAG ?= latest

# platform-specific settings
GOOS := $(shell go env GOOS)

ifeq ($(GOOS),windows)
	ifneq ($(suffix $(BUILD_OUTPUT_FILENAME)),.exe)
		BUILD_OUTPUT_FILENAME := $(addsuffix .exe,$(BUILD_OUTPUT_FILENAME))
	endif
endif

.PHONY: build
build:
	go build -o dist/api/$(BUILD_OUTPUT_FILENAME) $(BUILD_FLAGS) ./cmd/api

.PHONY: dev
dev:
	go run cmd/api/api.go

.PHONY: docker\:build
docker\:build:
	 docker build -f tools/Dockerfile -t agenti:$(DOCKER_TAG) .
