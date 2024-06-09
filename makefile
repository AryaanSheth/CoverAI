all: main.go
	go run main.go 

build: main.go
	go build -o main main.go
