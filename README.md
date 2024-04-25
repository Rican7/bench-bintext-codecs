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
BenchmarkDefaultString/Encode-24                          217360              5271 ns/op            4800 B/op        100 allocs/op
BenchmarkDefaultString/Decode-24                          423742              2899 ns/op               0 B/op          0 allocs/op
BenchmarkBase64Std/Encode-24                              180175              6581 ns/op            4800 B/op        200 allocs/op
BenchmarkBase64Std/Decode-24                              211950              5552 ns/op            2400 B/op        100 allocs/op
BenchmarkBase32Std/Encode-24                              164416              7405 ns/op            6400 B/op        200 allocs/op
BenchmarkBase32Std/Decode-24                              117444              9901 ns/op            3200 B/op        100 allocs/op
BenchmarkShortUUIDV3/Encode-24                              2406            487583 ns/op          268537 B/op      11269 allocs/op
BenchmarkShortUUIDV3/Decode-24                              2546            457102 ns/op           40987 B/op       3603 allocs/op
BenchmarkShortUUIDV4/Encode-24                              2578            442245 ns/op          263753 B/op       9769 allocs/op
BenchmarkShortUUIDV4/Decode-24                              6205            182044 ns/op           18466 B/op        800 allocs/op
BenchmarkULIDV2CrockfordBase32/Encode-24                  235040              4657 ns/op            3200 B/op        100 allocs/op
BenchmarkULIDV2CrockfordBase32/Decode-24                  738193              1620 ns/op               0 B/op          0 allocs/op
BenchmarkTypeIDCrockfordBase32/Encode-24                  263750              4364 ns/op            3200 B/op        100 allocs/op
BenchmarkTypeIDCrockfordBase32/Decode-24                  265795              4389 ns/op            1600 B/op        100 allocs/op
BenchmarkBTCBase58/Encode-24                                7982            138216 ns/op           31776 B/op       2572 allocs/op
BenchmarkBTCBase58/Decode-24                                9816            112937 ns/op           20000 B/op        800 allocs/op
PASS
ok      github.com/Rican7/bench-bintext-codecs/internal/uuid    19.319s
```
