
# Benchmark for JSON Pointer (RFC 6901) implementations for Go

## Tested implementations

* [dolmen-go/jsonptr](https://github.com/dolmen-go/jsonptr)
* [dustin/go-jsonpointer](https://github.com/dustin/go-jsonpointer)
* [xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer)
* [mickep76/jsonptr](https://github.com/mickep76/jsonptr)
* [lestrrat-go/jspointer](https://github.com/lestrrat-go/jspointer)
* [rnd42/go-jsonpointer](https://github.com/rnd42/go-jsonpointer)
* [twindagger/jsonptr](https://github.com/twindagger/jsonptr)
* [qri-io/jsonpointer](https://github.com/qri-io/jsonpointer)
* [json-validate/json-pointer-go](https://github.com/json-validate/json-pointer-go)
* [go-openapi/jsonpointer](https://github.com/go-openapi/jsonpointer)
* [oas3/json-pointer](https://github.com/oas3/json-pointer)
* [yukithm/json2csv/jsonpointer](https://github.com/yukithm/json2csv/tree/master/jsonpointer)

## Results

### 2026-09-21 (go1.27.0 linux/amd64)

Versions tested are in [`go.mod`](go.mod).

<!--
(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))
-->

#### [Get](benchmark_test.go#L45)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **41.46 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 58.92 ns/op | 48 B/op | **1 allocs/op** |
| go-openapi/jsonpointer | 278.4 ns/op | 240 B/op | 8 allocs/op |
| json-validate/json-pointer-go | 123.8 ns/op | 80 B/op | 2 allocs/op |
| lestrrat-go/jspointer | 218.9 ns/op | 320 B/op | 9 allocs/op |
| mickep76/jsonptr | 117.4 ns/op | 64 B/op | **1 allocs/op** |
| oas3/jsonpointer | 149.2 ns/op | 144 B/op | 3 allocs/op |
| qri-io/jsonpointer | 115.9 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 184.8 ns/op | 129 B/op | 3 allocs/op |
| twindagger/jsonptr | 135.5 ns/op | 136 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 144.8 ns/op | 144 B/op | 3 allocs/op |
| yukithm/json2csv/jsonpointer | 355.3 ns/op | 317 B/op | 17 allocs/op |

#### [Parse](benchmark_test.go#L100) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 45.72 ns/op | **56 B/op** | **2 allocs/op** |
| go-openapi/jsonpointer | 84.56 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 103.0 ns/op | 72 B/op | **2 allocs/op** |
| lestrrat-go/jspointer | 125.2 ns/op | 184 B/op | 4 allocs/op |
| **oas3/jsonpointer** | 38.50 ns/op | **56 B/op** | **2 allocs/op** |
| **qri-io/jsonpointer** | 90.85 ns/op | **56 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 160.5 ns/op | 121 B/op | 3 allocs/op |
| twindagger/jsonptr | 109.3 ns/op | 104 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **37.65 ns/op** | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 170.9 ns/op | 184 B/op | 9 allocs/op |

#### [Parse](benchmark_test.go#L100) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 81.27 ns/op | 72 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 84.97 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 146.6 ns/op | 88 B/op | 3 allocs/op |
| lestrrat-go/jspointer | 118.5 ns/op | 184 B/op | 4 allocs/op |
| **oas3/jsonpointer** | 38.50 ns/op | **56 B/op** | **2 allocs/op** |
| qri-io/jsonpointer | 141.5 ns/op | 72 B/op | 3 allocs/op |
| rnd42/go-jsonpointer | 266.3 ns/op | 137 B/op | 4 allocs/op |
| twindagger/jsonptr | 162.3 ns/op | 120 B/op | 4 allocs/op |
| **xeipuuv/gojsonpointer** | **36.91 ns/op** | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 199.1 ns/op | 173 B/op | 9 allocs/op |

#### [Parse](benchmark_test.go#L100) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 71.56 ns/op | 64 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 84.16 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 170.3 ns/op | 96 B/op | 4 allocs/op |
| lestrrat-go/jspointer | 107.6 ns/op | 176 B/op | 4 allocs/op |
| **oas3/jsonpointer** | 37.74 ns/op | **56 B/op** | **2 allocs/op** |
| qri-io/jsonpointer | 163.3 ns/op | 80 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 277.7 ns/op | 137 B/op | 5 allocs/op |
| twindagger/jsonptr | 176.4 ns/op | 128 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **37.28 ns/op** | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 187.1 ns/op | 168 B/op | 9 allocs/op |

#### [BackToString](benchmark_test.go#L109) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 49.19 ns/op | 56 B/op | 2 allocs/op |
| go-openapi/jsonpointer | 55.79 ns/op | 48 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 122.5 ns/op | 64 B/op | 3 allocs/op |
| **lestrrat-go/jspointer** | 3.314 ns/op | **0 B/op** | **0 allocs/op** |
| oas3/jsonpointer | 91.11 ns/op | 64 B/op | 3 allocs/op |
| qri-io/jsonpointer | 82.10 ns/op | 40 B/op | 2 allocs/op |
| **rnd42/go-jsonpointer** | **3.282 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 121.2 ns/op | 64 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 53.55 ns/op | 48 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 122.0 ns/op | 112 B/op | 4 allocs/op |

#### [BackToString](benchmark_test.go#L109) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 62.64 ns/op | 48 B/op | 2 allocs/op |
| go-openapi/jsonpointer | 54.71 ns/op | 48 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 159.7 ns/op | 80 B/op | 4 allocs/op |
| **lestrrat-go/jspointer** | 3.454 ns/op | **0 B/op** | **0 allocs/op** |
| oas3/jsonpointer | 87.67 ns/op | 64 B/op | 3 allocs/op |
| qri-io/jsonpointer | 118.4 ns/op | 45 B/op | 3 allocs/op |
| **rnd42/go-jsonpointer** | **3.098 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 163.7 ns/op | 80 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 54.83 ns/op | 48 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 164.6 ns/op | 144 B/op | 6 allocs/op |

#### [BackToString](benchmark_test.go#L109) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 39.89 ns/op | 16 B/op | 1 allocs/op |
| go-openapi/jsonpointer | 56.15 ns/op | 32 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 172.1 ns/op | 80 B/op | 5 allocs/op |
| **lestrrat-go/jspointer** | **3.395 ns/op** | **0 B/op** | **0 allocs/op** |
| oas3/jsonpointer | 91.11 ns/op | 48 B/op | 3 allocs/op |
| qri-io/jsonpointer | 130.9 ns/op | 53 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 3.466 ns/op | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 170.1 ns/op | 80 B/op | 5 allocs/op |
| xeipuuv/gojsonpointer | 50.54 ns/op | 32 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 156.7 ns/op | 128 B/op | 6 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L126) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | **95.65 ns/op** | 112 B/op | 4 allocs/op |
| go-openapi/jsonpointer | 153.8 ns/op | 168 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 236.5 ns/op | 136 B/op | 5 allocs/op |
| lestrrat-go/jspointer | 129.8 ns/op | 184 B/op | 4 allocs/op |
| oas3/jsonpointer | 142.9 ns/op | 120 B/op | 5 allocs/op |
| qri-io/jsonpointer | 189.4 ns/op | **96 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 168.4 ns/op | 121 B/op | **3 allocs/op** |
| twindagger/jsonptr | 258.8 ns/op | 168 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 113.7 ns/op | 104 B/op | 4 allocs/op |
| yukithm/json2csv/jsonpointer | 325.6 ns/op | 296 B/op | 13 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L126) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 157.9 ns/op | 120 B/op | 5 allocs/op |
| go-openapi/jsonpointer | 155.0 ns/op | 168 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 336.3 ns/op | 168 B/op | 7 allocs/op |
| lestrrat-go/jspointer | 123.8 ns/op | 184 B/op | **4 allocs/op** |
| oas3/jsonpointer | 141.9 ns/op | 120 B/op | 5 allocs/op |
| qri-io/jsonpointer | 276.8 ns/op | 117 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 266.2 ns/op | 137 B/op | **4 allocs/op** |
| twindagger/jsonptr | 347.1 ns/op | 200 B/op | 8 allocs/op |
| **xeipuuv/gojsonpointer** | **95.99 ns/op** | **104 B/op** | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 414.2 ns/op | 317 B/op | 15 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L126) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 129.7 ns/op | **80 B/op** | **4 allocs/op** |
| go-openapi/jsonpointer | 145.9 ns/op | 152 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 371.0 ns/op | 176 B/op | 9 allocs/op |
| lestrrat-go/jspointer | 108.3 ns/op | 176 B/op | **4 allocs/op** |
| oas3/jsonpointer | 152.2 ns/op | 104 B/op | 5 allocs/op |
| qri-io/jsonpointer | 296.6 ns/op | 136 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 283.9 ns/op | 137 B/op | 5 allocs/op |
| twindagger/jsonptr | 395.7 ns/op | 208 B/op | 10 allocs/op |
| **xeipuuv/gojsonpointer** | **91.43 ns/op** | 88 B/op | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 381.4 ns/op | 296 B/op | 15 allocs/op |

## Run

```sh
go test -bench . -benchmem | tee bench.txt
perl ./bench-to-table.pl < bench.txt > bench.md
```

## License

Copyright 2016-2021 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
