BINARY_NAME=simplepatientorder

.PHONY:build run docker-up docker-down mock install-mocks test
build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

mock: install-mocks
	mockgen --source=./internal/repository/patient.go --destination ./internal/repository/patient_mock.go --package repository

install-mocks: # could be replaced by local bin to avoid different version
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen@v1.6.0

test:
	go test ./...