FROM golang:1.19.0-alpine
COPY ./nina-api /

CMD ["./nina-api"]
