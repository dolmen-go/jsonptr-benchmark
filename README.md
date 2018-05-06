
# Benchmark for JSON Pointer (RFC 6901) implementations for Go

## Tested implementations

* [dolmen-go/jsonptr](https://github.com/dolmen-go/jsonptr)
* [dustin/go-jsonpointer](https://github.com/dustin/go-jsonpointer)
* [xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer)
* [mickep76/jsonptr](https://github.com/mickep76/jsonptr)
* [lestrrat/go-jspointer](https://github.com/lestrrat/go-jspointer)
* [rnd42/go-jsonpointer](https://github.com/rnd42/go-jsonpointer)
* [twindagger/jsonptr](https://github.com/twindagger/jsonptr)

## Results

### 2018-05-06 (go1.10.2 darwin/amd64)

(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))

#### [Get](benchmark_test.go#L40)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **105 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 154 ns/op | 48 B/op | **1 allocs/op** |
| lestrrat/go-jspointer | 662 ns/op | 368 B/op | 9 allocs/op |
| mickep76/jsonptr | 258 ns/op | 64 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 405 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 363 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 467 ns/op | 144 B/op | 3 allocs/op |

#### [Parse](benchmark_test.go#L95) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **133 ns/op** | **64 B/op** | **2 allocs/op** |
| lestrrat/go-jspointer | 286 ns/op | 240 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 346 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 305 ns/op | 112 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 138 ns/op | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L95) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 204 ns/op | 80 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 283 ns/op | 240 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 576 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 402 ns/op | 144 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **126 ns/op** | **64 B/op** | **2 allocs/op** |

#### [Parse](benchmark_test.go#L95) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 213 ns/op | 72 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 254 ns/op | 224 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 599 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 489 ns/op | 160 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **127 ns/op** | **64 B/op** | **2 allocs/op** |

#### [BackToString](benchmark_test.go#L104) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 118 ns/op | 80 B/op | 3 allocs/op |
| **lestrrat/go-jspointer** | 8.22 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **5.84 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 306 ns/op | 112 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 104 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L104) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 123 ns/op | 80 B/op | 3 allocs/op |
| **lestrrat/go-jspointer** | 8.37 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **5.72 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 403 ns/op | 144 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 104 ns/op | 64 B/op | 2 allocs/op |

#### [BackToString](benchmark_test.go#L104) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 88.0 ns/op | 32 B/op | 2 allocs/op |
| **lestrrat/go-jspointer** | 8.09 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **6.07 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 449 ns/op | 144 B/op | 8 allocs/op |
| xeipuuv/gojsonpointer | 95.6 ns/op | 32 B/op | 2 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L121) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 286 ns/op | 144 B/op | 5 allocs/op |
| lestrrat/go-jspointer | 295 ns/op | 240 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | 362 ns/op | **128 B/op** | **3 allocs/op** |
| twindagger/jsonptr | 655 ns/op | 224 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **236 ns/op** | **128 B/op** | 4 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L121) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 354 ns/op | 160 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 287 ns/op | 240 B/op | **4 allocs/op** |
| rnd42/go-jsonpointer | 588 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 843 ns/op | 288 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **249 ns/op** | **128 B/op** | **4 allocs/op** |

#### [ParseAndBackToString](benchmark_test.go#L121) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 310 ns/op | 104 B/op | 5 allocs/op |
| lestrrat/go-jspointer | 268 ns/op | 224 B/op | **4 allocs/op** |
| rnd42/go-jsonpointer | 604 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 917 ns/op | 304 B/op | 15 allocs/op |
| **xeipuuv/gojsonpointer** | **233 ns/op** | **96 B/op** | **4 allocs/op** |


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
