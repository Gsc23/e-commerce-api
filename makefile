# Inspired by https://github.com/bxcodec/go-clean-arch/blob/master/Makefile

bin:
	@ mkdir -p bin

BIN_DIR := bin
AIR_BIN := $(BIN_DIR)/air
GOLANGCI_BIN := $(BIN_DIR)/golangci-lint
DOCKER_DEV := docker-compose.dev.yml
DOCKER_PRD := docker-compose.yml

_air: $(AIR_BIN)

$(AIR_BIN): bin
	@ printf "Install air... "

	GOBIN=$(PWD)/$(BIN_DIR) go install github.com/air-verse/air@v1.64.0

	@ echo "done."

_golangci-lint: $(GOLANGCI_BIN)

$(GOLANGCI_BIN): bin
	@ printf "Install golangci-linter... "
	
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(PWD)/$(BIN_DIR) v2.8.0
	
	@ echo "done."

install-deps: _air _golangci-lint

up: 
	docker compose -f $(DOCKER_DEV) up

up-build:
	docker compose -f $(DOCKER_DEV) up --build

stop: 
	docker compose -f $(DOCKER_DEV) down

destroy: 
	docker compose -f $(DOCKER_DEV) down --remove-orphans -v 

run:
	@echo "Running the application..."

	@ $(AIR_BIN)

build:
	@echo "Building the application..."

	@ go build \
			-trimpath  \
			-buildvcs=false \
			-o tmp/main \
			./cmd/api

	@echo "Build completed."