# Context

Application in Golang responsible for exposing endpoints that receive an input file in csv, which its content contains a matrix of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row.

# How to run

### Locally

- Run web server
```
go run ./cmd/matrix
```
or via Makefile
```
make start-matrix-app-local
```

- Send request
```
curl -F 'file=@./tests/data/matrix.csv' "localhost:8080/echo"
```

### Container

- Run web server
```
COMPOSE_PROJECT_NAME=mcezario_backend_challenge docker-compose -f ./deployments/docker/docker-compose-matrix.yaml up -d
```
or via Makefile
```
start-matrix-app-container
```

- Send request
```
curl -F 'file=@./tests/data/matrix.csv' "localhost:8080/echo"
```

### Unit Tests

```
go test $(shell go list ./... | grep -v /cmd/)
```

or via Makefile
```
make unit-tests
```

### Enpoints
```
curl -F 'file=@./tests/data/matrix.csv' "localhost:8080/echo"
```
```
curl -F 'file=@./tests/data/matrix.csv' "localhost:8080/invert"
```
```
curl -F 'file=@./tests/data/matrix.csv' "localhost:8080/flatten"
```
```
curl -F 'file=@./tests/data/matrix.csv' "localhost:8080/sum"
```
```
curl -F 'file=@./tests/data/matrix.csv' "localhost:8080/multiply"
```
