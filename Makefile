docs:
	swag init -g cmd/api/main.go -o internal/docs

up:
	docker compose up -d

run:
	go run ./cmd/api

dev: containers-up run

watch: 
	air