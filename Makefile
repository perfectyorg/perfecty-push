SHELL := /bin/bash

default:
	@echo "Make utility. Supports the following actions: "
	@echo "  up: Start the service containers"
	@echo "  down: Stop the service containers"

up:
	@docker-compose up -d

down:
	@docker-compose down
