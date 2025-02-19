# Luke Stratta - Crosstech Backend Developer Challenge

## My Process

I started off by reading over the task. I broke it down into sections and created mini-tasks from this. This allowed me to see the project with the steps I needed to take.

I then examined the `data.json` file. I noted a few things from this. Firstly, the `NaN` in the `mileage` field occurences would need to be dealt with, as mentioned in the task.

There are also a few situations where `signal_id` is not unique and has different `elr` and `mileage` values for the same `signal_id`. This got me thinking that every entry would need to have an unique primary key. This may not be needed in the future, but it will help with data management going forward.

I don't know much about train line signals, so I had to research what some of the basic fields meant with a bit of internet searching, as well as making some assumptions for the project.

I will explain what these are shortly, and what I would to if I had stakeholders and subject matter experts around me.

Next it was time to start setting up the project. 

I wasn't familiar with the Echo web framework or pg-go, so I had to read up on the documentation, and use those as stepping stones to build out the application.

I have added comments throught out the code where I have not followed the exact way either of the modules have suggested it being used.

I then got into the meat and potatoes of the application and built out all the routes and database interactions.

I added tests for all the routes, but I know that these can be vastly improved. I will get into compromises I have made at the end of the README.

## What this application does

This application is a simple RESTful API with CRUD functionality to manage track and signal data.

To use it, you can use the command line or API testing software like [Insomnia](https://insomnia.rest/) or [Postman](https://www.postman.com/). 

You can also make raw SQL queries using the accompanying PGAdmin instance. 

See [Getting Started](#getting-started) to try it out.

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
make run
```

### Clean up the application

```bash
# stops and removes all Docker containers
# NOTE: this deletes all changes made to data
make cleanup
```

### Testing the application

Tests are associated by file.

```bash
# To run the tests
make test
```

## Accessing the application data

Once the application is running, to access the backend API, I recommend using something like [Insomnia](https://insomnia.rest/) or [Postman](https://www.postman.com/). 

You can then make API requests to the server.

Here is a list of endpoints available, what requests you can make, and examples of their responses.

```bash
GET tracks 

http://localhost:7777/tracks

RESPONSE: 

200 OK

{
	"tracks": [
		{
			"track_id": 1,
			"source": "London Liverpool Street",
			"target": "Bethnal Green"
		},
		{
			"track_id": 39,
			"source": "Bethnal Green",
			"target": "London Liverpool Street"
		},
		{
			"track_id": 50,
			"source": "Hackney Downs",
			"target": "Bethnal Green"
		}
    ]
}

---

GET /tracks/:id 

http://localhost:7777/tracks/3349

RESPONSE: 

200 OK

{
    "track_id": 3349,
    "source": "Waterloo East",
    "target": "London Bridge",
    "signal_ids": [
        {
            "signal_id": 13998,
            "signal_name": "SIG:TL4420(PL) BOROUGH MARKET",
            "elr": "YUE",
            "mileage": 4.0256
        },
        {
            "signal_id": 14000,
            "signal_name": "SIG:TL4424(CO) LONDON BRIDGE SE STN",
            "elr": "YUE",
            "mileage": 4.0558
        }
    ]
}

---

DELETE /tracks/:id
 
http://localhost:7777/tracks/3349

RESPONSE: 

200 OK

{
	"tracks": null,
	"message": "delete successful"
}

---

POST /tracks

body
{
	"track_id": 92774,
	"source": "Test Station 3",
	"target": "Test Station 4",
	"signal_ids": [
		{
			"signal_id": 13393,
			"signal_name": "SIG:WM791(CO)WEMBLEY CENTRAL STN",
			"elr": "MFD1",
			"mileage": 8.3815
		},
		{
			"signal_id": 13399,
			"signal_name": "SIG:WM1252(PL)WEMBLEY CENTRAL STN",
			"elr": "XGF1",
			"mileage": 2.9309
		}
	]		
}

RESPONSE: 

201 Created

{
	"tracks": [
		{
			"track_id": 92774,
			"source": "Test Station 3",
			"target": "Test Station 4",
			"signal_ids": [
				{
					"signal_id": 13393,
					"signal_name": "SIG:WM791(CO)WEMBLEY CENTRAL STN",
					"elr": "MFD1",
					"mileage": 8.3815
				},
				{
					"signal_id": 13399,
					"signal_name": "SIG:WM1252(PL)WEMBLEY CENTRAL STN",
					"elr": "XGF1",
					"mileage": 2.9309
				}
			]
		}
	],
	"message": "successfully created"
}

---

PUT /tracks/

```