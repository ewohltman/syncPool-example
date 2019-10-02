.PHONY: lint test debug build

lint:
	golangci-lint run --enable-all --deadline=5m ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

bench:
	go test -v -race -run=X -bench=. -benchmem ./...