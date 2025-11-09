.PHONY: run-local
run-local:
	go run cmd/url-shortener/main.go --config-path=./config/local.yaml
.PHONY: run-prod
run-prod:
	go run cmd/url-shortener/main.go --config-path=./config/prod.yaml