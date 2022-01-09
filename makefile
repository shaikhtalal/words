SHELL := /bin/bash -o pipefail
.PHONY: help

help:
	@echo "Usage: make <TARGET>"
	@echo ""
	@echo "Available targets are:"
	@echo ""
	@echo "    run-server                         Run the backend service"
	@echo ""


.PHONY: run-server
run-server:
	@echo "starting backend service!"
	@go run main.go serve 
