
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

### 2019-02-26 (go1.11.5 darwin/amd64)

(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))

#### [Get](benchmark_test.go#L42)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **105 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 151 ns/op | 48 B/op | **1 allocs/op** |
| json-validate/json-pointer-go | 298 ns/op | 80 B/op | 2 allocs/op |
| lestrrat/go-jspointer | 617 ns/op | 320 B/op | 9 allocs/op |
| mickep76/jsonptr | 252 ns/op | 64 B/op | **1 allocs/op** |
| qri-io/jsonpointer | 244 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 414 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 345 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 390 ns/op | 144 B/op | 3 allocs/op |

#### [Parse](benchmark_test.go#L97) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 131 ns/op | **64 B/op** | **2 allocs/op** |
| json-validate/json-pointer-go | 222 ns/op | 80 B/op | **2 allocs/op** |
| lestrrat/go-jspointer | 265 ns/op | 192 B/op | 4 allocs/op |
| **qri-io/jsonpointer** | 209 ns/op | **64 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 328 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 281 ns/op | 112 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **128 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L97) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 199 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 331 ns/op | 112 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 261 ns/op | 192 B/op | 4 allocs/op |
| qri-io/jsonpointer | 318 ns/op | 96 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 555 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 438 ns/op | 144 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **131 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L97) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 191 ns/op | 72 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 364 ns/op | 128 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 235 ns/op | 176 B/op | 4 allocs/op |
| qri-io/jsonpointer | 351 ns/op | 112 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 580 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 450 ns/op | 160 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **128 ns/op** | **64 B/op** | **2 allocs/op** |

#### [BackToString](benchmark_test.go#L106) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 117 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 305 ns/op | 112 B/op | 4 allocs/op |
| **lestrrat/go-jspointer** | 7.99 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 156 ns/op | 48 B/op | 2 allocs/op |
| **rnd42/go-jsonpointer** | **6.45 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 299 ns/op | 112 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 102 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L106) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 125 ns/op | 80 B/op | 3 allocs/op |
| json-validate/json-pointer-go | 391 ns/op | 144 B/op | 6 allocs/op |
| **lestrrat/go-jspointer** | 7.79 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 251 ns/op | 69 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | **6.43 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 395 ns/op | 144 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 106 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L106) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 87.1 ns/op | 32 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 435 ns/op | 144 B/op | 8 allocs/op |
| **lestrrat/go-jspointer** | 8.45 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 289 ns/op | 80 B/op | 6 allocs/op |
| **rnd42/go-jsonpointer** | **6.44 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 445 ns/op | 144 B/op | 8 allocs/op |
| xeipuuv/gojsonpointer | 90.8 ns/op | 32 B/op | 2 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 274 ns/op | 144 B/op | 5 allocs/op |
| json-validate/json-pointer-go | 578 ns/op | 192 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 283 ns/op | 192 B/op | 4 allocs/op |
| qri-io/jsonpointer | 384 ns/op | **112 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 335 ns/op | 128 B/op | **3 allocs/op** |
| twindagger/jsonptr | 650 ns/op | 224 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | **243 ns/op** | 128 B/op | 4 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 339 ns/op | 160 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 733 ns/op | 256 B/op | 10 allocs/op |
| lestrrat/go-jspointer | 269 ns/op | 192 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 587 ns/op | 165 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 558 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 828 ns/op | 288 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **235 ns/op** | **128 B/op** | **4 allocs/op** |

#### [ParseAndBackToString](benchmark_test.go#L123) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 298 ns/op | 104 B/op | 5 allocs/op |
| json-validate/json-pointer-go | 865 ns/op | 272 B/op | 14 allocs/op |
| lestrrat/go-jspointer | 257 ns/op | 176 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 677 ns/op | 192 B/op | 12 allocs/op |
| rnd42/go-jsonpointer | 641 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 927 ns/op | 304 B/op | 15 allocs/op |
| **xeipuuv/gojsonpointer** | **233 ns/op** | **96 B/op** | **4 allocs/op** |


## Run

```sh
dep ensure
go test -bench . -benchmem | tee bench.txt
perl ./bench-to-table.pl < bench.txt > bench.md
```

## License

Copyright 2016-2019 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
