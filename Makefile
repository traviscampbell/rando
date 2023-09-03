
# --- adapted from: https://github.com/filebrowser/filebrowser/blob/master/Makefile --------------
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

help: ## Show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "  ${YELLOW}%-15s${GREEN}%s${RESET}\n", $$1, $$2} \
	}' $(MAKEFILE_LIST)
# ------------------------------------------------------------------------------------------------

GOOS ?= $(shell go env GOOS)
ifeq ($(GOOS),windows)
	BIN_EXT := .exe
	RM := rmdir /s /q
else
	RM := rm -rf
endif

APP_NAME=rando

VERSION=`git describe --always --long --dirty|tr '\n' '-';date +%Y.%m.%d`
LDFLAGS=-ldflags "-w -s -X main.version=$(VERSION)" -trimpath
GO_BUILD=GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) go build $(LDFLAGS)

.SILENT: ;

build: ## tidy up, build the binary, and put it in ./bin
	echo "Building binary to ./bin/$(APP_NAME)"
	go mod tidy
	mkdir -p ./bin
	$(GO_BUILD) -o ./bin/$(APP_NAME)$(BIN_EXT) ./cmd/$(APP_NAME)/main.go

clean: ## run go clean and delete ./bin if it exists
	go clean
	$(RM) ./bin

# --- stolen from: https://github.com/traefikturkey/onramp/blob/master/Makefile ☜(꒡⌓꒡) ----------
excuse: ## fucking epic!
	curl -s programmingexcuses.com | grep -Eo "<a[^<>]+>[^<>]+</a>" | grep -Eo "[^<>]+" | sed -n 2p
# -------------------------------------------------------------------------------------------------

# bunch of fuckin' phonies
.PHONY: build docker-build docker-run clean excuse help
