# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY . /app/
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

EXPOSE 8080

CMD ["/main"]
