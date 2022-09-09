FROM golang:1.19.0-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /nina-api

CMD ["/nina-api"]
