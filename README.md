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
BenchmarkDefaultString/Encode-24                          222081              5281 ns/op            4800 B/op        100 allocs/op
BenchmarkDefaultString/Decode-24                          422239              2825 ns/op               0 B/op          0 allocs/op
BenchmarkBase64Std/Encode-24                              176221              6757 ns/op            4800 B/op        200 allocs/op
BenchmarkBase64Std/Decode-24                              205936              5609 ns/op            2400 B/op        100 allocs/op
BenchmarkBase64RawURLPrePadded/Encode-24                  115084              9987 ns/op            7200 B/op        300 allocs/op
BenchmarkBase64RawURLPrePadded/Decode-24                  189194              6037 ns/op            1600 B/op        100 allocs/op
BenchmarkBase32Std/Encode-24                              157074              7597 ns/op            6400 B/op        200 allocs/op
BenchmarkBase32Std/Decode-24                              117090             10083 ns/op            3200 B/op        100 allocs/op
BenchmarkBase32HexPrePadded/Encode-24                     110527             10156 ns/op            8800 B/op        300 allocs/op
BenchmarkBase32HexPrePadded/Decode-24                     120002              9780 ns/op            3200 B/op        100 allocs/op
BenchmarkStdLibCrockfordBase32/Encode-24                  111487             10254 ns/op            8800 B/op        300 allocs/op
BenchmarkStdLibCrockfordBase32/Decode-24                  123644              9691 ns/op            3200 B/op        100 allocs/op
BenchmarkULIDV2CrockfordBase32/Encode-24                  258001              4548 ns/op            3200 B/op        100 allocs/op
BenchmarkULIDV2CrockfordBase32/Decode-24                  713817              1611 ns/op               0 B/op          0 allocs/op
BenchmarkTypeIDCrockfordBase32/Encode-24                  268118              4326 ns/op            3200 B/op        100 allocs/op
BenchmarkTypeIDCrockfordBase32/Decode-24                  258442              4367 ns/op            1600 B/op        100 allocs/op
BenchmarkShortUUIDV3/Encode-24                              2404            488037 ns/op          268217 B/op      11255 allocs/op
BenchmarkShortUUIDV3/Decode-24                              2490            451558 ns/op           40962 B/op       3604 allocs/op
BenchmarkShortUUIDV4/Encode-24                              2616            441826 ns/op          263489 B/op       9758 allocs/op
BenchmarkShortUUIDV4/Decode-24                              6142            184262 ns/op           18386 B/op        800 allocs/op
BenchmarkBTCBase58/Encode-24                                7912            139247 ns/op           31640 B/op       2560 allocs/op
BenchmarkBTCBase58/Decode-24                               10000            115170 ns/op           20000 B/op        800 allocs/op
PASS
ok      github.com/Rican7/bench-bintext-codecs/internal/uuid    26.839s
```
