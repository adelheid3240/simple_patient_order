BINARY_NAME=simplepatientorder

.PHONY:build run docker-up docker-down
build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down