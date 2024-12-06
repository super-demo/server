# Docker
docker-up-dev:
	docker compose up -d postgres
	sleep 2
	docker compose up api

docker-down:
	docker compose down

# Migration
migrate-up:
	docker compose exec api go run cmd/db/main.go up

migrate-down:
	docker compose exec api go run cmd/db/main.go down

migrate-status:
	docker compose exec api go run cmd/db/main.go status

migrate-reset:
	docker compose exec api go run cmd/db/main.go reset

# Go
format:
	go fmt ./...

test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=coverage.txt

mockgen:
	go run cmd/mocks/main.go
