build: 
	@go build -o main cmd/main.go
.PHONY: build

run:
	@docker compose -f docker/compose.yaml up -d 
	@go run cmd/main.go
.PHONY: run 

run-dev: 
	@docker compose -f docker/compose.yaml up -d 
	@air
.PHONY: run-dev 

cleanup: 
	@docker stop pgadmin postgres
	@docker rm pgadmin postgres
.PHONY: cleanup