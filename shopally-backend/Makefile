# Simple Makefile to run tests using gotestsum

SHELL := /bin/sh
GO ?= go
GOTESTSUM ?= $(shell command -v gotestsum 2>/dev/null)
BIN_DIR := $(PWD)/bin

.PHONY: tools test tidy clean

# Install local tooling (gotestsum) into ./bin if not already available
tools:
	@if [ -z "$(GOTESTSUM)" ]; then \
		echo "Installing gotestsum into $(BIN_DIR)..."; \
		GOBIN=$(BIN_DIR) $(GO) install gotest.tools/gotestsum@latest; \
	else \
		echo "gotestsum present: $(GOTESTSUM)"; \
	fi

# Run test suite with gotestsum. Extra args can be passed via TEST_ARGS, e.g., TEST_ARGS='-run FXHandler'
TEST_FORMAT ?= short-verbose
TEST_ARGS ?=

test: tools
	@echo "Running tests with gotestsum..."
	@PATH=$(BIN_DIR):$$PATH gotestsum --format=$(TEST_FORMAT) -- -race -count=1 $(TEST_ARGS) ./...

# Keep go.mod/go.sum tidy
 tidy:
	$(GO) mod tidy

clean:
	@rm -rf $(BIN_DIR)
