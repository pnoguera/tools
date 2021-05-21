.PHONY: download install-tools

default: install-tools
all: install-tools

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | while read line; do \
		REPO=$$(echo $$line | awk -F'"' '{print $$2}' ); \
		BINARY=$$(echo $$line | grep _ | awk -F'"' '{print $$2}' | sed 's/.*\///;s/@.*//'); \
		NEW_BINARY=$$(echo $$line | sed -n 's/.*#name:\(.\+\)#.*/\1/p'); \
		echo "Installing $${REPO}"; \
		go install -v $${REPO}; \
		[ -n "$${NEW_BINARY}" ] && mv "$${GOBIN}/$${BINARY}" "$${GOBIN}/$${NEW_BINARY}" || true ; \
	done
