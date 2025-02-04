FROM golang:1.23.5

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./api ./api
COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/matrix

EXPOSE 8080

CMD ["./matrix"]