test:
	@go test -v ./...

build: 
	@go build -o main cmd/main.go
.PHONY: build

run:
	@docker compose -f docker/compose.yaml up -d 
	@sleep 2
	@go run cmd/main.go
.PHONY: run 

run-dev: 
	@docker compose -f docker/compose.yaml up -d 
	@sleep 2 # prevents panic when not able to connect to db
	@air
.PHONY: run-dev 

cleanup: 
	@docker stop pgadmin postgres
	@docker rm pgadmin postgres
.PHONY: cleanup

