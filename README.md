# `syncPool-example`
Example usage of `sync.Pool`.

## Benchmark Results

Run `make bench` to see the results yourself.

You'll need `benchstat` to see the comparison: `go get -u golang.org/x/perf/cmd/benchstat`

```
name               old time/op    new time/op    delta
UnmarshalObject-8    1.33µs ± 8%    1.25µs ± 2%   -6.27%  (p=0.001 n=10+8)

name               old alloc/op   new alloc/op   delta
UnmarshalObject-8      352B ± 0%      248B ± 0%  -29.55%  (p=0.000 n=10+10)

name               old allocs/op  new allocs/op  delta
UnmarshalObject-8      8.00 ± 0%      5.00 ± 0%  -37.50%  (p=0.000 n=10+10)
```
