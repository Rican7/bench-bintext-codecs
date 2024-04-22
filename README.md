# bench-bintext-codecs

_Benchmarks of binary-to-text codecs (encoder/decoders) in Go._


## Why?

Honestly, I've been reading so much about different kinds of k-sortable, unique IDs for entities/resources, and the different encodings to use for different reasons (length, case-sensitivty, etc), that I wondered if any of the well-known implementations (in Go) had any performance benefits, and decided to benchmark them.

Will this be the bottleneck of your app? Nah, most likely not, but it was a fun exploration nonetheless.

(Oh, and the name... I might come back and benchmark some values other than just UUIDs... so I kept it generic enough... whatever.)


## Benchmark Results

Here's the results of the benchmarking on my machine, but I'd suggest to run the benchmarks yourself:

```
$ uname -r
4.4.0-22621-Microsoft
$ go test -bench=. -benchmem ./...
goos: linux
goarch: amd64
pkg: github.com/Rican7/bench-bintext-codecs/internal/uuid
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkDefaultString/Encode-24                          207656              5344 ns/op            4800 B/op        100 allocs/op
BenchmarkDefaultString/Decode-24                          429649              2839 ns/op               0 B/op          0 allocs/op
BenchmarkBase64StdString/Encode-24                        168640              6570 ns/op            4800 B/op        200 allocs/op
BenchmarkBase64StdString/Decode-24                        208429              5532 ns/op            2400 B/op        100 allocs/op
BenchmarkBase32StdString/Encode-24                        149569              7206 ns/op            6400 B/op        200 allocs/op
BenchmarkBase32StdString/Decode-24                        119950              9964 ns/op            3200 B/op        100 allocs/op
BenchmarkShortUUIDV3String/Encode-24                        2240            492481 ns/op          268081 B/op      11242 allocs/op
BenchmarkShortUUIDV3String/Decode-24                        2656            430707 ns/op           41081 B/op       3609 allocs/op
BenchmarkShortUUIDV4String/Encode-24                        2659            437826 ns/op          263377 B/op       9744 allocs/op
BenchmarkShortUUIDV4String/Decode-24                        6204            180970 ns/op           18354 B/op        800 allocs/op
BenchmarkULIDV2CrockfordBase32String/Encode-24            249980              4540 ns/op            3200 B/op        100 allocs/op
BenchmarkULIDV2CrockfordBase32String/Decode-24            746518              1623 ns/op               0 B/op          0 allocs/op
BenchmarkTypeIDCrockfordBase32String/Encode-24            262449              4361 ns/op            3200 B/op        100 allocs/op
BenchmarkTypeIDCrockfordBase32String/Decode-24            261058              4391 ns/op            1600 B/op        100 allocs/op
BenchmarkBTCBase58String/Encode-24                          8098            136950 ns/op           31640 B/op       2565 allocs/op
BenchmarkBTCBase58String/Decode-24                         10000            111665 ns/op           20000 B/op        800 allocs/op
PASS
ok      github.com/Rican7/bench-bintext-codecs/internal/uuid    19.046s
```
