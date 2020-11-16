
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
* [go-openapi/jsonpointer](https://github.com/go-openapi/jsonpointer)
* [yukithm/json2csv/jsonpointer](https://github.com/yukithm/json2csv/tree/master/jsonpointer)

## Results

### 2020-11-01 (go1.15.3 darwin/amd64)

(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))

#### [Get](benchmark_test.go#L44)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **202 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 266 ns/op | 48 B/op | **1 allocs/op** |
| go-openapi/jsonpointer | 1039 ns/op | 240 B/op | 8 allocs/op |
| json-validate/json-pointer-go | 307 ns/op | 80 B/op | 2 allocs/op |
| lestrrat/go-jspointer | 1092 ns/op | 320 B/op | 9 allocs/op |
| mickep76/jsonptr | 408 ns/op | 64 B/op | **1 allocs/op** |
| qri-io/jsonpointer | 378 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 466 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 551 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 600 ns/op | 144 B/op | 3 allocs/op |
| yukithm/json2csv/jsonpointer | 5242 ns/op | 2717 B/op | 41 allocs/op |

#### [Parse](benchmark_test.go#L99) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 215 ns/op | **64 B/op** | **2 allocs/op** |
| go-openapi/jsonpointer | 305 ns/op | 128 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 248 ns/op | 80 B/op | **2 allocs/op** |
| lestrrat/go-jspointer | 327 ns/op | 192 B/op | 4 allocs/op |
| **qri-io/jsonpointer** | 236 ns/op | **64 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 379 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 513 ns/op | 112 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **182 ns/op** | **64 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 3306 ns/op | 1776 B/op | 25 allocs/op |

#### [Parse](benchmark_test.go#L99) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 229 ns/op | 80 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 343 ns/op | 128 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 486 ns/op | 112 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 287 ns/op | 192 B/op | 4 allocs/op |
| qri-io/jsonpointer | 366 ns/op | 96 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 665 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 522 ns/op | 144 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **144 ns/op** | **64 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 4021 ns/op | 1765 B/op | 25 allocs/op |

#### [Parse](benchmark_test.go#L99) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 374 ns/op | 72 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 489 ns/op | 128 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 516 ns/op | 128 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 316 ns/op | 176 B/op | 4 allocs/op |
| qri-io/jsonpointer | 465 ns/op | 112 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 684 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 618 ns/op | 160 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **197 ns/op** | **64 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 5340 ns/op | 1760 B/op | 25 allocs/op |

#### [BackToString](benchmark_test.go#L108) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 226 ns/op | 80 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 221 ns/op | 64 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 536 ns/op | 112 B/op | 4 allocs/op |
| **lestrrat/go-jspointer** | **11.4 ns/op** | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 290 ns/op | 48 B/op | 2 allocs/op |
| rnd42/go-jsonpointer | 12.0 ns/op | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 713 ns/op | 112 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 182 ns/op | 64 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 16185 ns/op | 26944 B/op | 28 allocs/op |

#### [BackToString](benchmark_test.go#L108) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 210 ns/op | 80 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 130 ns/op | 64 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 433 ns/op | 144 B/op | 6 allocs/op |
| **lestrrat/go-jspointer** | 9.27 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 274 ns/op | 69 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | **6.39 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 475 ns/op | 144 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 114 ns/op | 64 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 14325 ns/op | 27008 B/op | 32 allocs/op |

#### [BackToString](benchmark_test.go#L108) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 99.6 ns/op | 32 B/op | 2 allocs/op |
| go-openapi/jsonpointer | 104 ns/op | 32 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 482 ns/op | 144 B/op | 8 allocs/op |
| **lestrrat/go-jspointer** | 15.1 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 318 ns/op | 80 B/op | 6 allocs/op |
| **rnd42/go-jsonpointer** | **9.03 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 478 ns/op | 144 B/op | 8 allocs/op |
| xeipuuv/gojsonpointer | 112 ns/op | 32 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 10345 ns/op | 26976 B/op | 32 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L125) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 306 ns/op | 144 B/op | 5 allocs/op |
| go-openapi/jsonpointer | 436 ns/op | 192 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 637 ns/op | 192 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 461 ns/op | 192 B/op | 4 allocs/op |
| qri-io/jsonpointer | 421 ns/op | **112 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 414 ns/op | 128 B/op | **3 allocs/op** |
| twindagger/jsonptr | 738 ns/op | 224 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | **280 ns/op** | 128 B/op | 4 allocs/op |
| yukithm/json2csv/jsonpointer | 11536 ns/op | 28720 B/op | 53 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L125) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 378 ns/op | 160 B/op | 6 allocs/op |
| go-openapi/jsonpointer | 442 ns/op | 192 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 833 ns/op | 256 B/op | 10 allocs/op |
| lestrrat/go-jspointer | 490 ns/op | 192 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 759 ns/op | 165 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 609 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 918 ns/op | 288 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **272 ns/op** | **128 B/op** | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 11703 ns/op | 28776 B/op | 57 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L125) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 621 ns/op | 104 B/op | 5 allocs/op |
| go-openapi/jsonpointer | 655 ns/op | 160 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 943 ns/op | 272 B/op | 14 allocs/op |
| lestrrat/go-jspointer | 277 ns/op | 176 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 725 ns/op | 192 B/op | 12 allocs/op |
| rnd42/go-jsonpointer | 674 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 1016 ns/op | 304 B/op | 15 allocs/op |
| **xeipuuv/gojsonpointer** | **263 ns/op** | **96 B/op** | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 12719 ns/op | 28736 B/op | 57 allocs/op |


## Run

```sh
go test -bench . -benchmem | tee bench.txt
perl ./bench-to-table.pl < bench.txt > bench.md
```

## License

Copyright 2016-2020 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
