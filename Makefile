.PHONY: build
.DEFAULT_GOAL: build

unit-tests:
	go test $(shell go list ./... | grep -v /cmd/)

start-matrix-app-local:
	go run ./cmd/matrix

start-matrix-app-container:
	COMPOSE_PROJECT_NAME=mcezario_backend_challenge docker-compose -f ./deployments/docker/docker-compose-matrix.yaml up -d

