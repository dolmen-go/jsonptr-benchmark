
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
* [json-validate/json-pointer-go](https://github.com/json-validate/json-pointer-go)

## Results

### 2019-02-26 (go1.11.5 darwin/amd64)

(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))

#### [Get](benchmark_test.go#L42)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **113 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 165 ns/op | 48 B/op | **1 allocs/op** |
| json-validate/json-pointer-go | 290 ns/op | 80 B/op | 2 allocs/op |
| lestrrat/go-jspointer | 681 ns/op | 368 B/op | 9 allocs/op |
| mickep76/jsonptr | 268 ns/op | 64 B/op | **1 allocs/op** |
| qri-io/jsonpointer | 383 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 442 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 383 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 411 ns/op | 144 B/op | 3 allocs/op |

#### [Parse](benchmark_test.go#L97) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 153 ns/op | **64 B/op** | **2 allocs/op** |
| json-validate/json-pointer-go | 263 ns/op | 80 B/op | **2 allocs/op** |
| lestrrat/go-jspointer | 413 ns/op | 240 B/op | 4 allocs/op |
| **qri-io/jsonpointer** | 259 ns/op | **64 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 462 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 1137 ns/op | 112 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **149 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L97) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 216 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 388 ns/op | 112 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 303 ns/op | 240 B/op | 4 allocs/op |
| qri-io/jsonpointer | 435 ns/op | 96 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 629 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 434 ns/op | 144 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **134 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L97) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 209 ns/op | 72 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 408 ns/op | 128 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 450 ns/op | 224 B/op | 4 allocs/op |
| qri-io/jsonpointer | 381 ns/op | 112 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 741 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 498 ns/op | 160 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **131 ns/op** | **64 B/op** | **2 allocs/op** |

#### [BackToString](benchmark_test.go#L106) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 131 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 315 ns/op | 112 B/op | 4 allocs/op |
| **lestrrat/go-jspointer** | 9.28 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 168 ns/op | 48 B/op | 2 allocs/op |
| **rnd42/go-jsonpointer** | **5.98 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 323 ns/op | 112 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 110 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L106) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 242 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 448 ns/op | 144 B/op | 6 allocs/op |
| **lestrrat/go-jspointer** | 10.1 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 258 ns/op | 69 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | **9.47 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 460 ns/op | 144 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 110 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L106) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 94.0 ns/op | 32 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 468 ns/op | 144 B/op | 8 allocs/op |
| **lestrrat/go-jspointer** | 10.2 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 287 ns/op | 80 B/op | 6 allocs/op |
| **rnd42/go-jsonpointer** | **5.92 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 492 ns/op | 144 B/op | 8 allocs/op |
| xeipuuv/gojsonpointer | 104 ns/op | 32 B/op | 2 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 466 ns/op | 144 B/op | 5 allocs/op |
| json-validate/json-pointer-go | 609 ns/op | 192 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 384 ns/op | 240 B/op | 4 allocs/op |
| qri-io/jsonpointer | 419 ns/op | **112 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 362 ns/op | 128 B/op | **3 allocs/op** |
| twindagger/jsonptr | 1068 ns/op | 224 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | **266 ns/op** | 128 B/op | 4 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 369 ns/op | 160 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 819 ns/op | 256 B/op | 10 allocs/op |
| lestrrat/go-jspointer | 314 ns/op | 240 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 683 ns/op | 165 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 597 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 888 ns/op | 288 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **282 ns/op** | **128 B/op** | **4 allocs/op** |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 471 ns/op | 104 B/op | 5 allocs/op |
| json-validate/json-pointer-go | 1173 ns/op | 272 B/op | 14 allocs/op |
| lestrrat/go-jspointer | 368 ns/op | 224 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 703 ns/op | 192 B/op | 12 allocs/op |
| rnd42/go-jsonpointer | 674 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 1028 ns/op | 304 B/op | 15 allocs/op |
| **xeipuuv/gojsonpointer** | **251 ns/op** | **96 B/op** | **4 allocs/op** |


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
