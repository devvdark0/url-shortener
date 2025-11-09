.PHONY: run-local
run-local:
	go run cmd/url-shortener/main.go --config-path=./config/local.yaml
.PHONY: run-prod
run-prod:
	go run cmd/url-shortener/main.go --config-path=./config/prod.yaml
.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations -database "sqlite3://./storage/storage.db" up
.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations -database "sqlite3://./storage/storage.db" down