# Luke Stratta - Crosstech Backend Developer Challenge

## My Process

I started off by reading over the task. I broke it down into sections and created mini-tasks from this. This allowed me to see the project with the steps I needed to take.

I then examined the `data.json` file. I noted a few things from this. Firstly, the `NaN` in the `mileage` field occurences would need to be dealt with, as mentioned in the task.

There were also a few situations I found where `signal_name` field had a JSON `null` value that would 

## What this application does

## Project Structure and Extras

The `main.go` file lives in the `cmd/` directory. This is the main application directory and the entrypoint for the application.

Nearly everything else lives in the `internal/` directory. This is to separate out access to the parts of the code other imports don't need access to.

There is a `docker/` directory that hosts the Docker Compose yaml file for both Postgres and PGAdmin as a supporting piece of software. You can still `docker exec` into a container and access Postgres that way, if you so wish (I also like using the command-line for that too).

A Makefile is present to add some convenience aliases. See the Makefile for all commands.

Air is used as a hot-reload support tool for development. It helps when quickly making changes in the code and automatically watches for changes, builds a binary, and then runs it.

A Dockerfile to build a container image is available.

## Getting Started 

### Prerequisites

You must have the following installed:

- Go@v1.23 (minimum)
- Docker (or alternative, using the `docker` alias so the Makefile can be used)

This project uses the following required dependencies:

- [Echo](https://echo.labstack.com/)

```bash
go get github.com/labstack/echo/v4
```

- [go-pg](https://github.com/go-pg/pg) 
    - N.B. pg-go is in maintenance mode and will not be receiving new features

```bash
go get github.com/go-pg/pg/v10
```

I recommend the following optional dependencies:

- [Air](https://github.com/air-verse/air) for fast-reload development

```bash
go install github.com/air-verse/air@latest
```

### Configure the application

Firstly, copy the `docker/.env.example` file to `docker/.env`

```bash
cp docker/.env.example docker/.env
```

### Run the application

```bash
# starts the project in development mode using Air and Docker
make start

# alternatively, you can bypass Air
make docker
make run
```

### Clean up the application

```bash
# stops and removes all Docker containers
make cleanup
```

### Testing the application

Tests are associated by file.

```bash
# To run the tests
make test
```

## Accessing the application data

Once the application is running, to access the backend API, I recommend using something like [Insomnia]() or [Postman](). 

You can then make API requests to the server.

To see what requests you can make, please go to: 

```

```