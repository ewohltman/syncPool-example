.PHONY: lint test benchOriginal benchPool bench

lint:
	golangci-lint run ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

benchOriginal:
	@echo -n "" > original.txt
	@$(foreach var,$(shell seq 1 10),go test -run=NONE -bench=. -benchmem ./pkg/original | grep 'Benchmark' >> original.txt;)

benchPool:
	@echo -n "" > pool.txt
	@$(foreach var,$(shell seq 1 10),go test -run=NONE -bench=. -benchmem ./pkg/pool | grep 'Benchmark' >> pool.txt;)

bench: benchOriginal benchPool
	benchstat original.txt pool.txt
