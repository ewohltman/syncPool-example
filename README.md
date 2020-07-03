# `syncPool-example`
Example usage of `sync.Pool`.

## Benchmark Results

Run `make bench` to see the results yourself.

You'll need `benchcmp` to see the comparison: `go get -u golang.org/x/tools/cmd/benchcmp`

```
benchmark                      old ns/op     new ns/op     delta
BenchmarkUnmarshalObject-8     1286          1211          -5.83%

benchmark                      old allocs     new allocs     delta
BenchmarkUnmarshalObject-8     8              5              -37.50%

benchmark                      old bytes     new bytes     delta
BenchmarkUnmarshalObject-8     352           248           -29.55%
```
