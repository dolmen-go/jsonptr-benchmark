
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

### 2021-03-11 (go1.16 darwin/amd64)

(See also [latest results on Travis-CI](https://travis-ci.org/dolmen-go/jsonptr-benchmark))

#### [Get](benchmark_test.go#L44)

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **262.9 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 360.2 ns/op | 48 B/op | **1 allocs/op** |
| go-openapi/jsonpointer | 2590 ns/op | 240 B/op | 8 allocs/op |
| json-validate/json-pointer-go | 644.1 ns/op | 80 B/op | 2 allocs/op |
| lestrrat/go-jspointer | 1604 ns/op | 320 B/op | 9 allocs/op |
| mickep76/jsonptr | 741.4 ns/op | 64 B/op | **1 allocs/op** |
| qri-io/jsonpointer | 1115 ns/op | 48 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 882.2 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 1193 ns/op | 136 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 2438 ns/op | 144 B/op | 3 allocs/op |
| yukithm/json2csv/jsonpointer | 2787 ns/op | 317 B/op | 17 allocs/op |

#### [Parse](benchmark_test.go#L99) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **251.2 ns/op** | **56 B/op** | **2 allocs/op** |
| go-openapi/jsonpointer | 502.9 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 417.2 ns/op | 72 B/op | **2 allocs/op** |
| lestrrat/go-jspointer | 548.2 ns/op | 184 B/op | 4 allocs/op |
| qri-io/jsonpointer | 322.8 ns/op | **56 B/op** | **2 allocs/op** |
| rnd42/go-jsonpointer | 694.3 ns/op | 120 B/op | 3 allocs/op |
| twindagger/jsonptr | 452.5 ns/op | 104 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 261.0 ns/op | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 1143 ns/op | 184 B/op | 9 allocs/op |

#### [Parse](benchmark_test.go#L99) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 424.0 ns/op | 72 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 578.4 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 1291 ns/op | 88 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 458.0 ns/op | 184 B/op | 4 allocs/op |
| qri-io/jsonpointer | 427.0 ns/op | 72 B/op | 3 allocs/op |
| rnd42/go-jsonpointer | 903.9 ns/op | 136 B/op | 4 allocs/op |
| twindagger/jsonptr | 607.9 ns/op | 120 B/op | 4 allocs/op |
| **xeipuuv/gojsonpointer** | **213.4 ns/op** | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 923.9 ns/op | 173 B/op | 9 allocs/op |

#### [Parse](benchmark_test.go#L99) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 732.0 ns/op | 64 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 656.1 ns/op | 120 B/op | 4 allocs/op |
| json-validate/json-pointer-go | 1670 ns/op | 96 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 574.8 ns/op | 176 B/op | 4 allocs/op |
| qri-io/jsonpointer | 590.3 ns/op | 80 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 826.7 ns/op | 136 B/op | 5 allocs/op |
| twindagger/jsonptr | 738.9 ns/op | 128 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **343.5 ns/op** | **56 B/op** | **2 allocs/op** |
| yukithm/json2csv/jsonpointer | 2073 ns/op | 168 B/op | 9 allocs/op |

#### [BackToString](benchmark_test.go#L108) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 211.7 ns/op | 72 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 385.1 ns/op | 48 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 503.2 ns/op | 96 B/op | 4 allocs/op |
| **lestrrat/go-jspointer** | 11.69 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 228.8 ns/op | 40 B/op | 2 allocs/op |
| **rnd42/go-jsonpointer** | **7.864 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 519.3 ns/op | 96 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 370.7 ns/op | 48 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 749.0 ns/op | 112 B/op | 4 allocs/op |

#### [BackToString](benchmark_test.go#L108) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 188.2 ns/op | 64 B/op | 3 allocs/op |
| go-openapi/jsonpointer | 165.8 ns/op | 48 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 533.9 ns/op | 112 B/op | 5 allocs/op |
| **lestrrat/go-jspointer** | 14.99 ns/op | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 351.7 ns/op | 45 B/op | 3 allocs/op |
| **rnd42/go-jsonpointer** | **9.435 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 545.7 ns/op | 112 B/op | 5 allocs/op |
| xeipuuv/gojsonpointer | 233.7 ns/op | 48 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 976.9 ns/op | 176 B/op | 8 allocs/op |

#### [BackToString](benchmark_test.go#L108) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 181.4 ns/op | 32 B/op | 2 allocs/op |
| go-openapi/jsonpointer | 151.1 ns/op | 32 B/op | 2 allocs/op |
| json-validate/json-pointer-go | 602.2 ns/op | 112 B/op | 6 allocs/op |
| **lestrrat/go-jspointer** | **11.11 ns/op** | **0 B/op** | **0 allocs/op** |
| qri-io/jsonpointer | 747.4 ns/op | 48 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 13.39 ns/op | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 684.3 ns/op | 112 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 139.0 ns/op | 32 B/op | 2 allocs/op |
| yukithm/json2csv/jsonpointer | 660.4 ns/op | 160 B/op | 8 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L125) `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 778.5 ns/op | 128 B/op | 5 allocs/op |
| go-openapi/jsonpointer | 1406 ns/op | 168 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 1383 ns/op | 168 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 844.9 ns/op | 184 B/op | 4 allocs/op |
| **qri-io/jsonpointer** | **572.1 ns/op** | **96 B/op** | 4 allocs/op |
| rnd42/go-jsonpointer | 943.2 ns/op | 120 B/op | **3 allocs/op** |
| twindagger/jsonptr | 942.6 ns/op | 200 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | 711.2 ns/op | 104 B/op | 4 allocs/op |
| yukithm/json2csv/jsonpointer | 2395 ns/op | 296 B/op | 13 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L125) `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 701.9 ns/op | 136 B/op | 6 allocs/op |
| go-openapi/jsonpointer | 588.0 ns/op | 168 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 1191 ns/op | 200 B/op | 8 allocs/op |
| lestrrat/go-jspointer | 446.8 ns/op | 184 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 878.5 ns/op | 117 B/op | 6 allocs/op |
| rnd42/go-jsonpointer | 1099 ns/op | 136 B/op | **4 allocs/op** |
| twindagger/jsonptr | 1138 ns/op | 232 B/op | 9 allocs/op |
| **xeipuuv/gojsonpointer** | **416.4 ns/op** | **104 B/op** | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 1388 ns/op | 349 B/op | 17 allocs/op |

#### [ParseAndBackToString](benchmark_test.go#L125) `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 625.2 ns/op | 96 B/op | 5 allocs/op |
| go-openapi/jsonpointer | 589.5 ns/op | 152 B/op | 6 allocs/op |
| json-validate/json-pointer-go | 1502 ns/op | 208 B/op | 10 allocs/op |
| **lestrrat/go-jspointer** | **358.4 ns/op** | 176 B/op | **4 allocs/op** |
| qri-io/jsonpointer | 1416 ns/op | 136 B/op | 8 allocs/op |
| rnd42/go-jsonpointer | 1258 ns/op | 136 B/op | 5 allocs/op |
| twindagger/jsonptr | 1191 ns/op | 240 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | 627.2 ns/op | **88 B/op** | **4 allocs/op** |
| yukithm/json2csv/jsonpointer | 1696 ns/op | 328 B/op | 17 allocs/op |

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
