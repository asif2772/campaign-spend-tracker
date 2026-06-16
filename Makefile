deps fmt vet build test run

deps:
	docker compose up -d

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build ./...

test:
	go test -race ./...

run:
	go run ./cmd/server