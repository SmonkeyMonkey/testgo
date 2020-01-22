test:
	go test -timeout 20s -v -race -bench=. -benchmem ./...
run:
	go run main.go


.DEFAULT_GOAL := run