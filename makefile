test:
	go test -timeout 20s -v -race ./...
run:
	go run main.go

.DEFAULT_GOAL := run