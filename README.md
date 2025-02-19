# Luke Stratta - Crosstech Backend Developer Challenge

Thank you for taking the time to look at my application. I have included a table of contents for easy navigating around this README.

If you want to get started trying out the application, head to [Getting started](#getting-started).

## Contents
1. [My process](#my-process-and-the-application)
1. [What this application does](#what-this-application-does)
1. [Project structure and extras](#project-structure-and-extras)
1. [Assumptions, considerations, and improvements](#assumptions-considerations-and-improvements)
1. [Getting started](#getting-started)
1. [Accessing the applicatioin](#using-the-application)

### API Documentation

#### Track API endpoints

1. [GET /tracks](#get-tracks)
1. [GET /tracks/:id](#get-tracksid)
1. [GET /tracks?signal-id=id](#get-trackssignal-idid)
1. [DELETE /tracks/:id](#delete-tracksid)
1. [POST /tracks/:id](#post-tracks)


#### Signal API Endpoints

1. [GET /signals](#get-signals)
1. [POST /signals](#post-signals)
1. [PUT /signals](#put-signals)
1. [GET /signals/:id](#get-signalsid)
1. [DELETE /signals/:id](#delete-signalsid)

## My process and the application

I started off by reading over the task. I broke it down into sections and created mini-tasks from this. This allowed me to see the project with the steps I needed to take.

I then examined the `data.json` file. I noted a few things from this. Firstly, the `NaN` in the `mileage` field occurences would need to be dealt with, as mentioned in the task.

There are also a few situations where `signal_id` is not unique and has different `elr` and `mileage` values for the same `signal_id`. This got me thinking that every entry would need to have an unique primary key. This may not be needed in the future, but it will help with data management going forward.

I don't know much about train line signals, so I had to research what some of the basic fields meant with a bit of internet searching, as well as making some assumptions for the project.

I will explain what these are shortly, and what I would to if I had stakeholders and subject matter experts around me.

Next it was time to start setting up the project. 

I wasn't familiar with the Echo web framework or pg-go, so I had to read up on the documentation, and use those as stepping stones to build out the application.

I have added comments throught out the code where I have not followed the exact way either of the modules have suggested it being used and why I did that.

I then got into the meat and potatoes of the application and built out all the routes and database interactions.

I added tests for all the routes, but I know that these can be vastly improved. I will get into conisderations for developing out this project and compromises I have made later on in this README.

### What this application does

This application is a simple RESTful API with CRUD functionality to manage track and signal data.

To use it, you can use the command line or API testing software like [Insomnia](https://insomnia.rest/) or [Postman](https://www.postman.com/). 

You can also make raw SQL queries using the accompanying PGAdmin instance. 

See [Getting Started](#getting-started) to try it out.

### Project Structure and Extras

The `main.go` file lives in the `cmd/` directory. This is the main application directory and the entrypoint for the application.

Nearly everything else lives in the `internal/` directory. This is to separate out access to the parts of the code other imports don't need access to.

There is a `docker/` directory that hosts the Docker Compose yaml file for both Postgres and PGAdmin as a supporting piece of software. You can still `docker exec` into a container and access Postgres that way, if you so wish (I also like using the command-line for that too).

A Makefile is present to add some convenience aliases. See the Makefile for all commands.

Air is used as a hot-reload support tool for development. It helps when quickly making changes in the code and automatically watches for changes, builds a binary, and then runs it.

A Dockerfile to build a container image is available.

### Assumptions, considerations, and improvements

#### Assumptions made

I had to make a few assumptions for this project. Most of them revolved around what the data was actually showing. I had to assume that there were never going to be two tracks or signals that are the same.

If this was a real world project where I was able to interact with subject matter experts and other stakeholders, I would make sure that I have the contextual knowledge needed to be able to produce an application to spec while being able to provide consistent updates and implement feedback.

#### Considerations

To get this project done within the allotted timeframe, I had to reconsider a few things.

Data could have been handled better within the application, especially around the signals. There were lots of signals that had the same `signal_id` but the rest of the data with in the object would vary. This has caused some issues, so to resolve this, a primary key system would have been used throughout the application, rather than using the `signal_id`.

The tests only cover "happy path" situations. They don't cover malformed data or incorrect inputs, but there is some validation there to cover these instances, just not recorded in the tests. I would certainly spend time building these out to cover more of the code, and start to take note of any edge cases that I, and other people, experience.

On the validation side, I ran out of time here. Apart from a couple of `BadRequest` status codes, there is no input validation on the code. It is something I would take time to implement.

Another thing I ran out of time on was deploying it to GCP. I have a Dockerfile which is fully functioning and the whole application can be run as containers locally.

#### Improvements I would like to make

As noted above, I would like to move to a unique primary key system throughout the application, to guarantee data uniqueness. This may have led to duplicate data, but if I had more time, I would also implement data normalisatioin to see if any data could be combined into one object.

The handlers currently inteface directly through the database, but because there wasn't much business logic for this project, I felt it was an sensible omission to save time and prevent complexity. I would implement a service layer that would handle these interactions.

Tests would be greatly improved make sure endpoints, business logic, and data access functions are adhering strict standards.

Validation would be implemented to prevent any unwanted effects happening on the server with unvalidated code.

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

You will also find the application works with environment variables. You can export these if you would like to change the values. **They are configured even if you don't touch them**.

For more information, see `config/config.go` on how I have handled envirionment variables.

They are as follows with their default values:

```bash
export HOST="localhost"
export PORT="7777"
export DATABASE_URI="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
export ALLOWED_ORIGINS="http://localhost:7777,https://localhost:7777"
export POSTGRES_USERNAME="postgres"
export POSTGRES_PASSWORD="postgres"
export POSTGRES_DATABASE="postgres"
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

### Run the application as a Docker container

```bash
make docker-build # also runs the tests again

make docker-run
```

The application will then be live at:

```bash
http://localhost:7777
```

## Using the application

Once the application is running, to access the backend API, I recommend using something like [Insomnia](https://insomnia.rest/) or [Postman](https://www.postman.com/). 

You can then make API requests to the server.

Here is a list of endpoints available, what requests you can make, and examples of their responses.

### GET /tracks 

```bash
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

```

[Back to contents](#contents)

### GET /tracks/:id 

```bash
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

```

[Back to contents](#contents)


### DELETE /tracks/:id

```bash
http://localhost:7777/tracks/3349


RESPONSE: 

200 OK

{
	"message": "delete successful"
}

```

[Back to contents](#contents)

### POST /tracks

```bash
# example request
http://localhost:7777/tracks

# example request body
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

```

[Back to contents](#contents)

### PUT /tracks/

```bash
# example request
http://localhost:7777/tracks

# example request body
{
	"track_id": 92774,
	"source": "Test Station NEW",
	"target": "Test Station 4"
}

RESPONSE

200 OK

{
	"tracks": [
		{
			"track_pk": 0,
			"track_id": 92774,
			"source": "Test Station NEW",
			"target": "Test Station 4"
		}
	],
	"message": "update successful"
}
```
[Back to contents](#contents)

### GET /tracks?signal-id=id

```bash
# example
http://localhost:7777/tracks?signal-id=453

RESPONSE

200 OK

{
	"tracks": [
		{
			"track_id": 55,
			"source": "Acton Central",
			"target": "Willesden Junction"
		},
		{
			"track_id": 4084,
			"source": "Wembley Central",
			"target": "Acton Central"
		},
		{
			"track_id": 4522,
			"source": "Brent Cross West",
			"target": "Acton Central"
		}
	],
	"message": "request successful"
}

```
[Back to contents](#contents)

### GET /signals

```bash
# example
http://localhost:7777/signals

RESPONSE

200 OK

{
	"signals": [
		{
			"signal_pk": 2,
			"signal_id": 2848,
			"signal_name": "SIG:SN169(CO) IECC PDRF14 LOC R3/107",
			"elr": "ONM1",
			"mileage": 4.2126
		},
		{
			"signal_pk": 3,
			"signal_id": 2849,
			"signal_name": "SIG:SN173(CO) IECC PDMN02 LOC M3/144",
			"elr": "MNO1",
			"mileage": 5.6889
		}
	]
}
```

[Back to contents](#contents)

### POST /signals

```bash
# example 
http://localhost:7777/signals

# example request
{
	"signal_id": 1234,
	"signal_name": "Test signal",
	"elr": "TEN4",
	"mileage": 8.33333
}

RESPONSE

201 Created

{
	"signals": [
		{
			"signal_id": 1234,
			"signal_name": "Test signal",
			"elr": "TEN4",
			"mileage": 8.33333
		}
	],
	"message": "successfully created"
}
```

[Back to contents](#contents)

### PUT /signals

```bash 
# example 
http://localhost:7777/signals

# example request
{
	"signal_id": 453,
	"signal_name": "SIG:AW148(CO) ACTON WELLS JCN -- UPDATED --",
	"elr": "LPC5",
	"mileage": 3.1745
}

RESPONSE

200 OK

{
	"signals": [
		{
			"signal_id": 453,
			"signal_name": "SIG:AW148(CO) ACTON WELLS JCN -- UPDATED --",
			"elr": "LPC5",
			"mileage": 3.1745
		}
	],
	"message": "update successful"
}
```

[Back to contents](#contents)

### GET /signals/:id

```bash
# example 
http://localhost:7777/signals/19745

RESPONSE

200 OK 

{
	"signals": [
		{
			"signal_id": 19745,
			"signal_name": "",
			"elr": "NMO1",
			"mileage": 0
		},
		{
			"signal_id": 19745,
			"signal_name": "",
			"elr": "MNO1",
			"mileage": 0
		},
		{
			"signal_id": 19745,
			"signal_name": "",
			"elr": "OMN1",
			"mileage": 0
		},
		{
			"signal_id": 19745,
			"signal_name": "",
			"elr": "NOM1",
			"mileage": 0
		}
	],
	"message": "request successful"
}
```
[Back to contents](#contents)

### DELETE /signals/:id

```bash
# example 
http://localhost:7777/signals/453

RESPONSE

200 OK

{
	"message": "delete successful"
}

```