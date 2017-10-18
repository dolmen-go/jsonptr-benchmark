
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

### 2017-10-18

#### BenchmarkGet

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **137 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 174 ns/op | 48 B/op | **1 allocs/op** |
| lestrrat/go-jspointer | 979 ns/op | 272 B/op | 9 allocs/op |
| mickep76/jsonptr | 302 ns/op | 64 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 482 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 406 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 301 ns/op | 48 B/op | **1 allocs/op** |

#### BenchmarkParse/"/foo/bar"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 155 ns/op | **64 B/op** | **2 allocs/op** |
| lestrrat/go-jspointer | 335 ns/op | 128 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 397 ns/op | 104 B/op | 3 allocs/op |
| twindagger/jsonptr | 343 ns/op | 112 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **146 ns/op** | **64 B/op** | **2 allocs/op** |

#### BenchmarkParse/"/foo/bar/baz/\~0\~1"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 226 ns/op | 98 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 633 ns/op | 234 B/op | 9 allocs/op |
| rnd42/go-jsonpointer | 833 ns/op | 170 B/op | 7 allocs/op |
| twindagger/jsonptr | 666 ns/op | 186 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **165 ns/op** | **96 B/op** | **2 allocs/op** |

#### BenchmarkBackToString/"/foo/bar"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 74.0 ns/op | 24 B/op | 2 allocs/op |
| **lestrrat/go-jspointer** | 9.58 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **6.94 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 340 ns/op | 64 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 119 ns/op | 16 B/op | 2 allocs/op |

#### BenchmarkBackToString/"/foo/bar/baz/\~0\~1"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 108 ns/op | 64 B/op | 2 allocs/op |
| **lestrrat/go-jspointer** | 9.23 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **7.23 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 532 ns/op | 160 B/op | 9 allocs/op |
| xeipuuv/gojsonpointer | 146 ns/op | 64 B/op | 3 allocs/op |

#### BenchmarkParseAndBackToString/"/foo/bar"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | **239 ns/op** | 88 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 348 ns/op | 128 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 426 ns/op | 104 B/op | **3 allocs/op** |
| twindagger/jsonptr | 839 ns/op | 176 B/op | 7 allocs/op |
| xeipuuv/gojsonpointer | 255 ns/op | **80 B/op** | 4 allocs/op |

#### BenchmarkParseAndBackToString/"/foo/bar/baz/\~0\~1"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 459 ns/op | 162 B/op | **5 allocs/op** |
| lestrrat/go-jspointer | 698 ns/op | 234 B/op | 9 allocs/op |
| rnd42/go-jsonpointer | 842 ns/op | 170 B/op | 7 allocs/op |
| twindagger/jsonptr | 1178 ns/op | 344 B/op | 16 allocs/op |
| **xeipuuv/gojsonpointer** | **312 ns/op** | **160 B/op** | **5 allocs/op** |


## Run

```sh
go test -bench . -benchmem | tee bench.txt
perl ./bench-to-table.pl < bench.txt > bench.md
```

## License

Copyright 2017 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
