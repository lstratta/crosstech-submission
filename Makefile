# test the application code
test:
	@go test -v ./...
.PHONY: test

# build the binary after testing the code
build: test
	@go build -o main cmd/main.go
.PHONY: build

# simply starts the database containers and the development server 
start: 
	@docker compose -f docker/compose.yaml up -d 
	@sleep 2 # prevents panic when not able to connect to db
	@air
.PHONY: run-dev 

# starts the servre without Air
run:
	@docker compose -f docker/compose.yaml up -d 
	@sleep 2 # prevents panic when not able to connect to db
	@go run cmd/main.go
.PHONY: run 

# stops and removes all containers 
cleanup: 
	@docker stop pgadmin postgres
	@docker rm pgadmin postgres
.PHONY: cleanup

# TODO(luke): add docker build and run commands
build-containers: test
