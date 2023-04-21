BINARY_NAME=k8s-image-admission-controller
IMAGE_NAME=$(BINARY_NAME)
PORT=8080

DEBUG ?=

SHELL = /bin/bash -o pipefail

DIR = $(shell pwd)

CONFIG_HOME = $(or ${XDG_CONFIG_HOME},${XDG_CONFIG_HOME},${HOME}/.config)

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
INFO_COLOR=\033[36m
WHITE_COLOR=\033[1m
OK=[✅]
KO=[❌]
WARN=[⚠️]

.DEFAULT_GOAL := help

COMMIT?=${BUILDCOMMIT}
VERSION?=${BUILDTAG}

# enable cgo because it's required by OSX keychain library
CGO_ENABLED=1
# enable go modules
GO111MODULE=on


export CGO_ENABLED GO111MODULE

check-%:
	@if $$(hash $* 2> /dev/null); then \
		echo -e "$(OK_COLOR)$(OK)$(NO_COLOR) $*"; \
	else \
		echo -e "$(ERROR_COLOR)$(KO)$(NO_COLOR) $*"; \
	fi

define check_output
	if [ $(1) -eq 0 ]; then \
		echo -e "$(OK) $(OK_COLOR) SUCCESS $(NO_COLOR)"; \
	else \
		echo -e "$(KO) $(ERROR_COLOR) FAIL $(NO_COLOR)"; \
	fi
endef


##@ Utils

.PHONY: help
help: ## Help
	@echo -e "$(OK_COLOR)                  $(BANNER)$(NO_COLOR)"
	@echo "------------------------------------------------------------------"
	@echo ""
	@echo -e "${ERROR_COLOR}Usage${NO_COLOR}: make ${INFO_COLOR}<target>${NO_COLOR}"
	@awk 'BEGIN {FS = ":.*##"; } /^[a-zA-Z_-]+:.*?##/ { printf "  ${INFO_COLOR}%-25s${NO_COLOR} %s\n", $$1, $$2 } /^##@/ { printf "\n${WHITE_COLOR}%s${NO_COLOR}\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	@echo ""

.PHONY: check
check: check-go check-golangci-lint check-goreleaser check-docker ## Check utils requirements are installed

##@ Docker

.PHONY: Docker
docker: ## Docker
	@docker build . -t ghcr.io/thibaultserti/$(IMAGE_NAME):latest

.PHONY: Docker debug
docker-debug: Docker debug
	@docker build . -f Dockerfile.debug  -t $(IMAGE_NAME)-debug:latest

.PHONY: Docker run
docker-run: ## Docker run
	@docker run -it --rm -p $(PORT):$(PORT) ghcr.io/thibaultserti/$(IMAGE_NAME):latest

.PHONY: Docker debug run
docker-debug-run: ## Docker debug run
	@docker run -it --rm -p $(PORT):$(PORT) $(IMAGE_NAME)-debug:latest

.PHONY: Docker push
docker-push: ## Docker debug run
	@docker push ghcr.io/thibaultserti/$(IMAGE_NAME):latest

##@ Go

.PHONY: dep
dep: ## Download deps in vendor/
	@go mod vendor

.PHONY: tests
tests: ## Tests
	@echo "Running tests..."
	@go test ./test; \
	$(call check_output,$$?)

.PHONY: coverage
coverage: ## Coverage
	@echo "Running coverage..."
	@go test -v -race -vet off -covermode atomic -coverprofile tmp/coverage.gocov ./test/... -coverpkg=./...; \
	$(call check_output,$$?)

.PHONY: format
format: ## Format
	@echo "Running format..."
	@go fmt ./...; \
	$(call check_output,$$?)

.PHONY: lint
lint: ## Lint
	@echo "Running vet"
	@go vet ./...; \
	$(call check_output,$$?)
	@echo "Running lint..."
	@golangci-lint run; \
	$(call check_output,$$?)

.PHONY: quality
quality: format lint  ## Run all quality

.PHONY: build
build: ## Cross platform build
	@goreleaser build --single-target --clean --snapshot

.PHONY: run
run: # Run
	@go run cmd/${BINARY_NAME}/main.go

.PHONY: clean
clean: ## Clean binaries
	@go clean
	@rm -rf dist/; \
	$(call check_output,$$?)
