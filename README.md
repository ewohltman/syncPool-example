# `syncPool-example`
Example usage of sync.Pool.

## Benchmark Results

### Original implementation without sync.Pool
```
goos: linux
goarch: amd64
pkg: github.com/ewohltman/syncPool-example/pkg/original
BenchmarkUnmarshalObject-8   	  130453	      9251 ns/op	     288 B/op	       7 allocs/op
PASS
ok  	github.com/ewohltman/syncPool-example/pkg/original	3.310s
```

### Pooled implementation with sync.Pool
```
goos: linux
goarch: amd64
pkg: github.com/ewohltman/syncPool-example/pkg/pool
BenchmarkUnmarshalObject-8   	  121831	      9902 ns/op	     258 B/op	       5 allocs/op
PASS
ok  	github.com/ewohltman/syncPool-example/pkg/pool	3.279s
```