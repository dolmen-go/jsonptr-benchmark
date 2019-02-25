
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
* [json-validate/json-pointer-go](https://json-validate/json-pointer-go)

## Results

### 2019-02-25 (go1.11.5 darwin/amd64)

(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))

#### [Get](benchmark_test.go#L42)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **111 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 166 ns/op | 48 B/op | **1 allocs/op** |
| json-validate/json-pointer-go | 290 ns/op | 80 B/op | 2 allocs/op |
| lestrrat/go-jspointer | 737 ns/op | 368 B/op | 9 allocs/op |
| mickep76/jsonptr | 275 ns/op | 64 B/op | **1 allocs/op** |
| qri-io/jsonpointer | 413 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 449 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 390 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 397 ns/op | 144 B/op | 3 allocs/op |

#### [Parse](benchmark_test.go#L97) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 152 ns/op | **64 B/op** | **2 allocs/op** |
| json-validate/json-pointer-go | 259 ns/op | 80 B/op | **2 allocs/op** |
| lestrrat/go-jspointer | 320 ns/op | 240 B/op | 4 allocs/op |
| **qri-io/jsonpointer** | 238 ns/op | **64 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 354 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 318 ns/op | 112 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **136 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L97) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 216 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 384 ns/op | 112 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 304 ns/op | 240 B/op | 4 allocs/op |
| qri-io/jsonpointer | 366 ns/op | 96 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 589 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 637 ns/op | 144 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **136 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L97) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 220 ns/op | 72 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 1092 ns/op | 128 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 275 ns/op | 224 B/op | 4 allocs/op |
| qri-io/jsonpointer | 385 ns/op | 112 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 631 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 469 ns/op | 160 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **135 ns/op** | **64 B/op** | **2 allocs/op** |

#### [BackToString](benchmark_test.go#L106) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 129 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 329 ns/op | 112 B/op | 4 allocs/op |
| **lestrrat/go-jspointer** | 9.14 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 158 ns/op | 48 B/op | 2 allocs/op |
| **rnd42/go-jsonpointer** | **6.43 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 320 ns/op | 112 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 110 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L106) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 132 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 450 ns/op | 144 B/op | 6 allocs/op |
| **lestrrat/go-jspointer** | 9.15 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 396 ns/op | 69 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | **5.97 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 467 ns/op | 144 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 108 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L106) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 92.5 ns/op | 32 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 476 ns/op | 144 B/op | 8 allocs/op |
| **lestrrat/go-jspointer** | 8.92 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 291 ns/op | 80 B/op | 6 allocs/op |
| **rnd42/go-jsonpointer** | **6.30 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 781 ns/op | 144 B/op | 8 allocs/op |
| xeipuuv/gojsonpointer | 103 ns/op | 32 B/op | 2 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 288 ns/op | 144 B/op | 5 allocs/op |
| json-validate/json-pointer-go | 683 ns/op | 192 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 329 ns/op | 240 B/op | 4 allocs/op |
| qri-io/jsonpointer | 860 ns/op | **112 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 377 ns/op | 128 B/op | **3 allocs/op** |
| twindagger/jsonptr | 1014 ns/op | 224 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | **252 ns/op** | 128 B/op | 4 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 366 ns/op | 160 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 815 ns/op | 256 B/op | 10 allocs/op |
| lestrrat/go-jspointer | 314 ns/op | 240 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 725 ns/op | 165 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 595 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 955 ns/op | 288 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **254 ns/op** | **128 B/op** | **4 allocs/op** |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 527 ns/op | 104 B/op | 5 allocs/op |
| json-validate/json-pointer-go | 931 ns/op | 272 B/op | 14 allocs/op |
| lestrrat/go-jspointer | 286 ns/op | 224 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 706 ns/op | 192 B/op | 12 allocs/op |
| rnd42/go-jsonpointer | 633 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 1079 ns/op | 304 B/op | 15 allocs/op |
| **xeipuuv/gojsonpointer** | **249 ns/op** | **96 B/op** | **4 allocs/op** |


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
