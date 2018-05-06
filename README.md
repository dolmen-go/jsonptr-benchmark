
# Benchmark for JSON Pointer (RFC 6901) implementations for Go

## Tested implementations

* [dolmen-go/jsonptr](https://github.com/dolmen-go/jsonptr)
* [dustin/go-jsonpointer](https://github.com/dustin/go-jsonpointer)
* [xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer)
* [mickep76/jsonptr](https://github.com/mickep76/jsonptr)
* [lestrrat/go-jspointer](https://github.com/lestrrat/go-jspointer)
* [rnd42/go-jsonpointer](https://github.com/rnd42/go-jsonpointer)
* [twindagger/jsonptr](https://github.com/twindagger/jsonptr)
* [qri-io/jsonpointer](https://github.com/qri-io/jsonpointer)

## Results

### 2018-05-06 (go1.10.2 darwin/amd64)

(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))

#### [Get](benchmark_test.go#L41)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **101 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 155 ns/op | 48 B/op | **1 allocs/op** |
| lestrrat/go-jspointer | 641 ns/op | 368 B/op | 9 allocs/op |
| mickep76/jsonptr | 255 ns/op | 64 B/op | **1 allocs/op** |
| qri-io/jsonpointer | 244 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 397 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 357 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 373 ns/op | 144 B/op | 3 allocs/op |

#### [Parse](benchmark_test.go#L96) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 131 ns/op | **64 B/op** | **2 allocs/op** |
| lestrrat/go-jspointer | 280 ns/op | 240 B/op | 4 allocs/op |
| **qri-io/jsonpointer** | 214 ns/op | **64 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 343 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 291 ns/op | 112 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **123 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L96) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 199 ns/op | 80 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 275 ns/op | 240 B/op | 4 allocs/op |
| qri-io/jsonpointer | 318 ns/op | 96 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 554 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 398 ns/op | 144 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **121 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L96) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 193 ns/op | 72 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 252 ns/op | 224 B/op | 4 allocs/op |
| qri-io/jsonpointer | 349 ns/op | 112 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 588 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 430 ns/op | 160 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **120 ns/op** | **64 B/op** | **2 allocs/op** |

#### [BackToString](benchmark_test.go#L105) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 125 ns/op | 80 B/op | 3 allocs/op |
| **lestrrat/go-jspointer** | 8.03 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 143 ns/op | 48 B/op | 2 allocs/op |
| **rnd42/go-jsonpointer** | **6.37 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 301 ns/op | 112 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 101 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L105) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 125 ns/op | 80 B/op | 3 allocs/op |
| **lestrrat/go-jspointer** | 8.72 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 235 ns/op | 69 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | **6.41 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 431 ns/op | 144 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 121 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L105) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 88.0 ns/op | 32 B/op | 2 allocs/op |
| **lestrrat/go-jspointer** | 8.61 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 299 ns/op | 80 B/op | 6 allocs/op |
| **rnd42/go-jsonpointer** | **6.64 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 548 ns/op | 144 B/op | 8 allocs/op |
| xeipuuv/gojsonpointer | 94.0 ns/op | 32 B/op | 2 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L122) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 267 ns/op | 144 B/op | 5 allocs/op |
| lestrrat/go-jspointer | 292 ns/op | 240 B/op | 4 allocs/op |
| qri-io/jsonpointer | 385 ns/op | **112 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 346 ns/op | 128 B/op | **3 allocs/op** |
| twindagger/jsonptr | 626 ns/op | 224 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | **233 ns/op** | 128 B/op | 4 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L122) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 345 ns/op | 160 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 287 ns/op | 240 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 583 ns/op | 165 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 564 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 820 ns/op | 288 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **233 ns/op** | **128 B/op** | **4 allocs/op** |

#### [ParseAndBackToString](benchmark_test.go#L122) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 304 ns/op | 104 B/op | 5 allocs/op |
| lestrrat/go-jspointer | 262 ns/op | 224 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 646 ns/op | 192 B/op | 12 allocs/op |
| rnd42/go-jsonpointer | 595 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 898 ns/op | 304 B/op | 15 allocs/op |
| **xeipuuv/gojsonpointer** | **229 ns/op** | **96 B/op** | **4 allocs/op** |


## Run

```sh
dep ensure
go test -bench . -benchmem | tee bench.txt
perl ./bench-to-table.pl < bench.txt > bench.md
```

## License

Copyright 2016-2018 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
