GOCMD=go
GORUN=$(GOCMD) run
GOFMT=$(GOCMD) fmt
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test


.PHONY: help
help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: fmt
fmt: ## go fmt
	$(GOFMT) $(go list ./... | grep -v /vendor/)

.PHONY: test
test: ## go test
	$(GOTEST) ./...

.PHONY: clean
clean: ## go clean
	$(GOCLEAN)


.PHONY: build
build: ## build in ./bin/gdb
	$(GOBUILD) -o bin/gbd main.go
