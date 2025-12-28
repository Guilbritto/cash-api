.PHONY: migrate-up migrate-down migrate-steps migrate-version migrate-force

ifneq (,$(wildcard .env))
	include .env
	export
endif

docs:
	swag init -g cmd/api/main.go -o internal/docs

up:
	docker compose up -d

run:
	go run ./cmd/api

watch: 
	air

# ====== CONFIG ======
MIGRATIONS_PATH=migrations
CLI=go run ./cmd/cli

# ====== MIGRATIONS ======

migrate-up:
	$(CLI) -path=$(MIGRATIONS_PATH) -cmd=up

migrate-down:
	$(CLI) -path=$(MIGRATIONS_PATH) -cmd=down

migrate-steps:
	@if [ -z "$(STEPS)" ]; then \
		echo "Use: make migrate-steps STEPS=N"; \
		exit 1; \
	fi
	$(CLI) -path=$(MIGRATIONS_PATH) -cmd=steps -steps=$(STEPS)

migrate-version:
	$(CLI) -path=$(MIGRATIONS_PATH) -cmd=version

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Use: make migrate-force VERSION=N"; \
		exit 1; \
	fi
	$(CLI) -path=$(MIGRATIONS_PATH) -cmd=force -force=$(VERSION)

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Use: make migrate-create NAME=nome_da_migration"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)
