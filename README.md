
# Benchmark for JSON Pointer (RFC 6901) implementations for Go

## Tested implementations

* [dolmen-go/jsonptr](https://github.com/dolmen-go/jsonptr)
* [dustin/go-jsonpointer](https://github.com/dustin/go-jsonpointer)
* [xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer)
* [mickep76/jsonptr](https://github.com/mickep76/jsonptr)
* [lestrrat/go-jspointer](https://github.com/lestrrat/go-jspointer)

## Results

### 2017-04-05

#### BenchmarkGet

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **82.7 ns/op** | **0 B/op** | **0 allocs/op** |
| dustin/go-jsonpointer | 272 ns/op | 48 B/op | 1 allocs/op |
| lestrrat/go-jspointer | 1226 ns/op | 272 B/op | 9 allocs/op |
| mickep76/jsonptr | 567 ns/op | 67 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 609 ns/op | 128 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 386 ns/op | 48 B/op | 1 allocs/op |

#### BenchmarkParse/"/foo/bar"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 234 ns/op | **64 B/op** | **2 allocs/op** |
| lestrrat/go-jspointer | 437 ns/op | 128 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 461 ns/op | 104 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **209 ns/op** | **64 B/op** | **2 allocs/op** |

#### BenchmarkParse/"/foo/bar/baz/\~0\~1"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 397 ns/op | 98 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 850 ns/op | 234 B/op | 9 allocs/op |
| rnd42/go-jsonpointer | 1001 ns/op | 170 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **288 ns/op** | **96 B/op** | **2 allocs/op** |

#### BenchmarkBackToString/"/foo/bar"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 107 ns/op | 24 B/op | 2 allocs/op |
| **lestrrat/go-jspointer** | 11.4 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **7.83 ns/op** | **0 B/op** | **0 allocs/op** |
| xeipuuv/gojsonpointer | 126 ns/op | 16 B/op | 2 allocs/op |

#### BenchmarkBackToString/"/foo/bar/baz/\~0\~1"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 158 ns/op | 64 B/op | 2 allocs/op |
| **lestrrat/go-jspointer** | 11.7 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **7.83 ns/op** | **0 B/op** | **0 allocs/op** |
| xeipuuv/gojsonpointer | 211 ns/op | 64 B/op | 3 allocs/op |

#### BenchmarkParseAndBackToString/"/foo/bar"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 349 ns/op | 88 B/op | 4 allocs/op |
| lestrrat/go-jspointer | 452 ns/op | 128 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 459 ns/op | 104 B/op | **3 allocs/op** |
| **xeipuuv/gojsonpointer** | **337 ns/op** | **80 B/op** | 4 allocs/op |

#### BenchmarkParseAndBackToString/"/foo/bar/baz/\~0\~1"

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 571 ns/op | 162 B/op | **5 allocs/op** |
| lestrrat/go-jspointer | 864 ns/op | 234 B/op | 9 allocs/op |
| rnd42/go-jsonpointer | 1017 ns/op | 170 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **513 ns/op** | **160 B/op** | **5 allocs/op** |


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
