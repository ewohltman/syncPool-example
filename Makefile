.PHONY: lint test benchOriginal benchPool bench

lint:
	golangci-lint run --enable-all --deadline=5m ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

benchOriginal:
	go test -run=NONE -bench=. -benchmem ./pkg/original > original.txt

benchPool:
	go test -run=NONE -bench=. -benchmem ./pkg/pool > pool.txt

bench: benchOriginal benchPool
	benchcmp original.txt pool.txt