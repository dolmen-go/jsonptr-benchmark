
# Benchmark for JSON Pointer (RFC 6901) implementations for Go

## Tested implementations

* [dolmen-go/jsonptr](https://github.com/dolmen-go/jsonptr)
* [dustin/go-jsonpointer](https://github.com/dustin/go-jsonpointer)
* [xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer)
* [mickep76/jsonptr](https://github.com/mickep76/jsonptr)
* [lestrrat/go-jspointer](https://github.com/lestrrat/go-jspointer)

## Results

### 2017-03-31

| Benchmark | impl | speed | allocs bytes | allocs count |
| --- | --- | ---: | ---: | ---: |
| BenchmarkGet | dolmen-go/jsonptr | **86.7 ns/op** | **0 B/op** | **0 allocs/op** |
| BenchmarkGet | dustin/go-jsonpointer | 274 ns/op | 48 B/op | 1 allocs/op |
| BenchmarkGet | lestrrat/go-jspointer | 1240 ns/op | 272 B/op | 9 allocs/op |
| BenchmarkGet | mickep76/jsonptr | 572 ns/op | 67 B/op | 4 allocs/op |
| BenchmarkGet | rnd42/go-jsonpointer | 617 ns/op | 128 B/op | 3 allocs/op |
| BenchmarkGet | xeipuuv/gojsonpointer | 385 ns/op | 48 B/op | 1 allocs/op |
| BenchmarkParse/"/foo/bar" | dolmen-go/jsonptr | 243 ns/op | **64 B/op** | **2 allocs/op** |
| BenchmarkParse/"/foo/bar" | lestrrat/go-jspointer | 439 ns/op | 128 B/op | 4 allocs/op |
| BenchmarkParse/"/foo/bar" | rnd42/go-jsonpointer | 459 ns/op | 104 B/op | 3 allocs/op |
| BenchmarkParse/"/foo/bar" | xeipuuv/gojsonpointer | **206 ns/op** | **64 B/op** | **2 allocs/op** |
| BenchmarkParse/"/foo/bar/baz/\~0\~1" | dolmen-go/jsonptr | 398 ns/op | 98 B/op | 3 allocs/op |
| BenchmarkParse/"/foo/bar/baz/\~0\~1" | lestrrat/go-jspointer | 848 ns/op | 234 B/op | 9 allocs/op |
| BenchmarkParse/"/foo/bar/baz/\~0\~1" | rnd42/go-jsonpointer | 1001 ns/op | 170 B/op | 7 allocs/op |
| BenchmarkParse/"/foo/bar/baz/\~0\~1" | xeipuuv/gojsonpointer | **296 ns/op** | **96 B/op** | **2 allocs/op** |
| BenchmarkBackToString/"/foo/bar" | dolmen-go/jsonptr | 111 ns/op | 24 B/op | 2 allocs/op |
| BenchmarkBackToString/"/foo/bar" | lestrrat/go-jspointer | 11.6 ns/op | **0 B/op** | **0 allocs/op** |
| BenchmarkBackToString/"/foo/bar" | rnd42/go-jsonpointer | **7.80 ns/op** | **0 B/op** | **0 allocs/op** |
| BenchmarkBackToString/"/foo/bar" | xeipuuv/gojsonpointer | 127 ns/op | 16 B/op | 2 allocs/op |
| BenchmarkBackToString/"/foo/bar/baz/\~0\~1" | dolmen-go/jsonptr | 159 ns/op | 64 B/op | 2 allocs/op |
| BenchmarkBackToString/"/foo/bar/baz/\~0\~1" | lestrrat/go-jspointer | 11.3 ns/op | **0 B/op** | **0 allocs/op** |
| BenchmarkBackToString/"/foo/bar/baz/\~0\~1" | rnd42/go-jsonpointer | **7.91 ns/op** | **0 B/op** | **0 allocs/op** |
| BenchmarkBackToString/"/foo/bar/baz/\~0\~1" | xeipuuv/gojsonpointer | 208 ns/op | 64 B/op | 3 allocs/op |
| BenchmarkParseAndBackToString/"/foo/bar" | dolmen-go/jsonptr | 373 ns/op | 88 B/op | 4 allocs/op |
| BenchmarkParseAndBackToString/"/foo/bar" | lestrrat/go-jspointer | 443 ns/op | 128 B/op | 4 allocs/op |
| BenchmarkParseAndBackToString/"/foo/bar" | rnd42/go-jsonpointer | 463 ns/op | 104 B/op | **3 allocs/op** |
| BenchmarkParseAndBackToString/"/foo/bar" | xeipuuv/gojsonpointer | **336 ns/op** | **80 B/op** | 4 allocs/op |
| BenchmarkParseAndBackToString/"/foo/bar/baz/\~0\~1" | dolmen-go/jsonptr | 572 ns/op | 162 B/op | **5 allocs/op** |
| BenchmarkParseAndBackToString/"/foo/bar/baz/\~0\~1" | lestrrat/go-jspointer | 873 ns/op | 234 B/op | 9 allocs/op |
| BenchmarkParseAndBackToString/"/foo/bar/baz/\~0\~1" | rnd42/go-jsonpointer | 1000 ns/op | 170 B/op | 7 allocs/op |
| BenchmarkParseAndBackToString/"/foo/bar/baz/\~0\~1" | xeipuuv/gojsonpointer | **519 ns/op** | **160 B/op** | **5 allocs/op** |

## Run

```go
go test -bench . -benchmem
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
