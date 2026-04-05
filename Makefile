BINARY_DIR := bin
GO := go

.PHONY: all build install clean test

all: build

build:
	$(GO) build -o $(BINARY_DIR)/yt ./cmd/yt
	$(GO) build -o $(BINARY_DIR)/youtube ./cmd/youtube

install:
	$(GO) install ./cmd/yt
	$(GO) install ./cmd/youtube

test:
	$(GO) test ./...

clean:
	rm -rf $(BINARY_DIR)
