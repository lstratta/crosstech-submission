# test the application code
# pgadmin does not need to run
# cleanup first to remove any leftover data
test: cleanup
	@echo "starting containers"
	@docker compose -f docker/compose.yaml up -d postgres
	@sleep 2 # prevents panic when not able to connect to db
	@go test -v ./...
	@echo "cleaning up.. "
	@docker stop postgres
	@docker rm postgres
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
.PHONY: start

# starts the servre without Air
run:
	@docker compose -f docker/compose.yaml up -d 
	@sleep 2 # prevents panic when not able to connect to db
	@go run cmd/main.go
.PHONY: run 

# stops and removes all containers 
cleanup: 
	@docker stop pgadmin postgres -i
	@docker rm pgadmin postgres -i
.PHONY: cleanup

# TODO(luke): add docker build and run commands
docker-build: test
	@docker build . -t crosstech/track-sig:latest

run-containers:
	@docker compose -f docker/compose-all.yaml up -d 
.PHONY: run-containers
