BINARY_NAME=simplepatientorder

.PHONY:build run 
build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}