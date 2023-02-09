APP_VERSION = `git describe --tag --abbrev=0`
BUILD_DATE 	= `date -u +%Y%m%d.%H%M%S`
DOCKER_BUILD_TAG := dev

test: ## Run unit-tests
	@echo "\n${GREEN}Running unit-tests${NC}"
	go test -race -v -covermode=atomic -coverprofile=coverage.out $$(go list ./... | grep -v cmd)
	go tool cover -func coverage.out | grep total | awk '{print "" $$3}'


fmt: ## Auto formatting Golang code
	@echo "\n${GREEN}Auto formatting golang code with golangci-lint${NC}"
	golangci-lint run --fix
	@echo "\n${GREEN}Auto formatting golang code with gofmt${NC}"
	gofmt -w -l $$(go list -f "{{ .Dir }}" ./...); if [ "$${errors}" != "" ]; then echo "$${errors}"; fi

linting: golangci-lint ## Linting Golang code

golangci-lint: ## Linting Golang code with golangci
	@echo "\n${GREEN}Linting Golang code with golangci${NC}"
	golangci-lint --version
	golangci-lint cache clean
	golangci-lint run ./... -v --timeout 240s

build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.buildDate=$(date -u +%Y%m%d.%H%M%S) -X main.buildCommit=${APP_VERSION}" \
