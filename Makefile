
run:
	go run ./cmd/server

test:
	go test -race ./...

lint:
	go vet ./...

