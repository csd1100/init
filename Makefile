clean:
	rm -r build
format:
	gofmt -w .
build:
	go build -o build/ .
run:
	go run .
test:
	go test ./...