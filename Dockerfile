FROM docker.io/golang:1.23-alpine AS build

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/main .

FROM docker.io/alpine
WORKDIR /app

COPY --from=build /go/bin/main .
COPY data/data.json data/data.json

ENV HOST="0.0.0.0"

EXPOSE 7777

CMD ["./main"]