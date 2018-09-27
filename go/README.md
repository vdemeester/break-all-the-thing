# Merge strings benchmark

```
~/s/g/v/break-all-the-thing{master:…}/go λ go test -benchmem -bench=.
goos: linux
goarch: amd64
pkg: github.com/vdemeester/break-all-the-thing/go
BenchmarkSliceContaining/1-8    20000000                99.7 ns/op            96 B/op          1 allocs/op
BenchmarkSliceContaining/2-8    20000000                99.3 ns/op            96 B/op          1 allocs/op
BenchmarkSliceContaining/4-8    20000000                98.9 ns/op            96 B/op          1 allocs/op
BenchmarkSliceContaining/8-8    20000000                98.2 ns/op            96 B/op          1 allocs/op
BenchmarkSliceContaining/16-8           20000000                98.3 ns/op            96 B/op          1 allocs/op
BenchmarkWithMap/1-8                    10000000               225 ns/op              96 B/op          1 allocs/op
BenchmarkWithMap/2-8                    10000000               226 ns/op              96 B/op          1 allocs/op
BenchmarkWithMap/4-8                    10000000               226 ns/op              96 B/op          1 allocs/op
BenchmarkWithMap/8-8                    10000000               224 ns/op              96 B/op          1 allocs/op
BenchmarkWithMap/16-8                   10000000               225 ns/op              96 B/op          1 allocs/op
PASS
ok      github.com/vdemeester/break-all-the-thing/go    22.844s
```
