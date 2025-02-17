# Luke Stratta - Crosstech Backend Developer Challenge

## What this application does

## Development on this application

## Getting Started 

### Prerequisites

You must have the following installed:

- Go@v1.23 (minimum)
- Docker (or alternative, using the `docker` alias)

This project uses the following required dependencies:

- [Echo](https://echo.labstack.com/)
- [go-pg](https://github.com/go-pg/pg) 
    - N.B. pg-go is in maintenance mode and will not be receiving new features

It also uses the following optional dependencies:

- [Air](https://github.com/air-verse/air) for fast-reload development

To install all dependencies:

```bash
go mod tidy
```

### Run the application

```bash
# starts the project in development mode using Air and Docker
make run-dev

# alternatively, you can bypass Air
make docker
make run
```

### Clean up the application

```bash
# removes all Docker containers
make cleanup
```