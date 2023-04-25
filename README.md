
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
* [oas3/json-pointer](https://github.com/oas3/json-pointer)
* [yukithm/json2csv/jsonpointer](https://github.com/yukithm/json2csv/tree/master/jsonpointer)

## Results

### 2023-04-25 (go1.20.3 darwin/amd64)

<!--
(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))
-->

#### [Get](benchmark_test.go#L45)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **601.9 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 609.1 ns/op | 48 B/op | **1 allocs/op** |
| go-openapi/jsonpointer | 3497 ns/op | 240 B/op | 8 allocs/op |
| json-validate/json-pointer-go | 1953 ns/op | 80 B/op | 2 allocs/op |
| lestrrat/go-jspointer | 2772 ns/op | 320 B/op | 9 allocs/op |
| mickep76/jsonptr | 1177 ns/op | 64 B/op | **1 allocs/op** |
| oas3/jsonpointer | 1485 ns/op | 144 B/op | 3 allocs/op |
| qri-io/jsonpointer | 1893 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 1398 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 2440 ns/op | 136 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 1865 ns/op | 144 B/op | 3 allocs/op |
| yukithm/json2csv/jsonpointer | 4559 ns/op | 317 B/op | 17 allocs/op |

#### [Parse](benchmark_test.go#L100) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **597.3 ns/op** | **56 B/op** | **2 allocs/op** |
| go-openapi/jsonpointer | 1301 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 1394 ns/op | 72 B/op | **2 allocs/op** |
| lestrrat/go-jspointer | 1452 ns/op | 184 B/op | 4 allocs/op |
| oas3/jsonpointer | 852.8 ns/op | **56 B/op** | **2 allocs/op** |
| qri-io/jsonpointer | 2074 ns/op | **56 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 1586 ns/op | 120 B/op | 3 allocs/op |
| twindagger/jsonptr | 1742 ns/op | 104 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 655.1 ns/op | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 2349 ns/op | 184 B/op | 9 allocs/op |

#### [Parse](benchmark_test.go#L100) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 1113 ns/op | 72 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 1356 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 1631 ns/op | 88 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 1276 ns/op | 184 B/op | 4 allocs/op |
| **oas3/jsonpointer** | 841.8 ns/op | **56 B/op** | **2 allocs/op** |
| qri-io/jsonpointer | 1598 ns/op | 72 B/op | 3 allocs/op |
| rnd42/go-jsonpointer | 2972 ns/op | 136 B/op | 4 allocs/op |
| twindagger/jsonptr | 1892 ns/op | 120 B/op | 4 allocs/op |
| **xeipuuv/gojsonpointer** | **595.5 ns/op** | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 3282 ns/op | 173 B/op | 9 allocs/op |

#### [Parse](benchmark_test.go#L100) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 1243 ns/op | 64 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 1325 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 2009 ns/op | 96 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 1108 ns/op | 176 B/op | 4 allocs/op |
| **oas3/jsonpointer** | **639.6 ns/op** | **56 B/op** | **2 allocs/op** |
| qri-io/jsonpointer | 1705 ns/op | 80 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 3015 ns/op | 137 B/op | 5 allocs/op |
| twindagger/jsonptr | 4139 ns/op | 128 B/op | 5 allocs/op |
| xeipuuv/gojsonpointer | 1073 ns/op | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 3098 ns/op | 168 B/op | 9 allocs/op |

#### [BackToString](benchmark_test.go#L109) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 550.7 ns/op | 72 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 905.3 ns/op | 48 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 2051 ns/op | 96 B/op | 4 allocs/op |
| **lestrrat/go-jspointer** | 66.05 ns/op | **0 B/op** | **0 allocs/op** |
| oas3/jsonpointer | 1244 ns/op | 64 B/op | 3 allocs/op |
| qri-io/jsonpointer | 1049 ns/op | 40 B/op | 2 allocs/op |
| **rnd42/go-jsonpointer** | **33.83 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 2878 ns/op | 96 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 901.4 ns/op | 48 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 2000 ns/op | 112 B/op | 4 allocs/op |

#### [BackToString](benchmark_test.go#L109) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 646.4 ns/op | 64 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 568.7 ns/op | 48 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 1893 ns/op | 112 B/op | 5 allocs/op |
| **lestrrat/go-jspointer** | 38.75 ns/op | **0 B/op** | **0 allocs/op** |
| oas3/jsonpointer | 1197 ns/op | 64 B/op | 3 allocs/op |
| qri-io/jsonpointer | 1645 ns/op | 45 B/op | 3 allocs/op |
| **rnd42/go-jsonpointer** | **27.81 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 1670 ns/op | 112 B/op | 5 allocs/op |
| xeipuuv/gojsonpointer | 625.3 ns/op | 48 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 2057 ns/op | 176 B/op | 8 allocs/op |

#### [BackToString](benchmark_test.go#L109) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 496.1 ns/op | 32 B/op | 2 allocs/op |
| go-openapi/jsonpointer | 499.1 ns/op | 32 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 2283 ns/op | 112 B/op | 6 allocs/op |
| **lestrrat/go-jspointer** | 38.46 ns/op | **0 B/op** | **0 allocs/op** |
| oas3/jsonpointer | 1333 ns/op | 48 B/op | 3 allocs/op |
| qri-io/jsonpointer | 1686 ns/op | 48 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | **35.97 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 1656 ns/op | 112 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 481.8 ns/op | 32 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 2353 ns/op | 160 B/op | 8 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L126) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 1556 ns/op | 128 B/op | 5 allocs/op |
| go-openapi/jsonpointer | 2027 ns/op | 168 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 2254 ns/op | 168 B/op | 6 allocs/op |
| lestrrat/go-jspointer | **1383 ns/op** | 184 B/op | 4 allocs/op |
| oas3/jsonpointer | 1963 ns/op | 120 B/op | 5 allocs/op |
| qri-io/jsonpointer | 2065 ns/op | **96 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 2092 ns/op | 120 B/op | **3 allocs/op** |
| twindagger/jsonptr | 2350 ns/op | 200 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | 1536 ns/op | 104 B/op | 4 allocs/op |
| yukithm/json2csv/jsonpointer | 4762 ns/op | 296 B/op | 13 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L126) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 2407 ns/op | 136 B/op | 6 allocs/op |
| go-openapi/jsonpointer | 2122 ns/op | 168 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 2850 ns/op | 200 B/op | 8 allocs/op |
| **lestrrat/go-jspointer** | **1625 ns/op** | 184 B/op | **4 allocs/op** |
| oas3/jsonpointer | 1931 ns/op | 120 B/op | 5 allocs/op |
| qri-io/jsonpointer | 2794 ns/op | 117 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 3168 ns/op | 137 B/op | **4 allocs/op** |
| twindagger/jsonptr | 4816 ns/op | 232 B/op | 9 allocs/op |
| **xeipuuv/gojsonpointer** | 1651 ns/op | **104 B/op** | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 5676 ns/op | 349 B/op | 17 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L126) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 2038 ns/op | 96 B/op | 5 allocs/op |
| go-openapi/jsonpointer | 1844 ns/op | 152 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 3979 ns/op | 208 B/op | 10 allocs/op |
| lestrrat/go-jspointer | 1326 ns/op | 176 B/op | **4 allocs/op** |
| oas3/jsonpointer | 2062 ns/op | 104 B/op | 5 allocs/op |
| qri-io/jsonpointer | 4250 ns/op | 136 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 3707 ns/op | 136 B/op | 5 allocs/op |
| twindagger/jsonptr | 5139 ns/op | 240 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **1036 ns/op** | **88 B/op** | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 4162 ns/op | 328 B/op | 17 allocs/op |

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
