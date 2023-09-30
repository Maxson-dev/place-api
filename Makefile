TARGET:=place-api
LOCAL_BIN:=$(CURDIR)/bin
GO_IMPORTS:=${LOCAL_BIN}/goimports
GO_SWAGGER:=$(LOCAL_BIN)/swag
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOOSE:=$(LOCAL_BIN)/goose

# linter version to use
GOLANGCI_TAG:=1.51.1

.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: install-lint
install-lint: ## install golangci-lint binary
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
endif

.PHONY: download-deps
download-deps: install-lint
	$(info Getting project specific dependencies...)
	mkdir -p ${LOCAL_BIN}
	test -f ${GOOSE} || GOBIN=${LOCAL_BIN} go install github.com/pressly/goose/v3/cmd/goose@latest
	test -f ${GO_IMPORTS} || GOBIN=${LOCAL_BIN} go install golang.org/x/tools/cmd/goimports@latest
	test -f ${GO_SWAGGER} || GOBIN=${LOCAL_BIN} go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: go-generate
go-generate: ### run go-generate
	go generate ./...

.PHONY: swag-v1
swag-v1: ### swag init
	$(GO_SWAGGER) init -g ./internal/controller/service.go -o ./api

.PHONY: generate
generate: download-deps go-generate swag-v1 ## runs all code generations

.PHONY: lint
lint: install-lint ## linter for golang
	$(GOLANGCI_BIN) run ./...

.PHONY: go-fmt
go-fmt:	## runs go fmt
	go fmt $(PWD)/...

.PHONY: new-migration
new-migration: ## create new migration
	$(GOOSE) -dir migration/migrations create migration sql

.PHONY: up
up: ## run app
	docker-compose up -d

.PHONY: down
down: ## down app
	docker-compose down

.PHONY: build-up
build-up: ## build and run app
	docker-compose up --build -d


.PHONY: test
test: ## run tests
	go test -v ./...