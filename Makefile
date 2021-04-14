run:
	go run main.go
build:
	go build -o bin/application main.go
test:
	go test ./...