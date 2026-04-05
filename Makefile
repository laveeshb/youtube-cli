BINARY_DIR := bin
GO := go

.PHONY: all build clean

all: build

build:
	$(GO) build -o $(BINARY_DIR)/yt ./cmd/yt
	$(GO) build -o $(BINARY_DIR)/youtube ./cmd/youtube

clean:
	rm -rf $(BINARY_DIR)
