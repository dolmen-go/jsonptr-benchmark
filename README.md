
# Benchmark for JSON Pointer (RFC 6901) implementations for Go

## Results

| package                                                           | speed       | allocs       | alloc size |
|-------------------------------------------------------------------|------------:|-------------:|-----------:|
| [dolmen-go/jsonptr](https://github.com/dolmen-go/jsonptr)         |   108 ns/op |  0 allocs/op |     0 B/op |
| [dustin/go-jsonpointer](https://github.com/dustin/go-jsonpointer) |   289 ns/op |  1 allocs/op |    48 B/op |
| [xeipuuv/gojsonpointer](https://github.com/xeipuuv/gojsonpointer) |   402 ns/op |  1 allocs/op |    48 B/op |
| [mickep76/jsonptr](https://github.com/mickep76/jsonptr)           |   580 ns/op |  4 allocs/op |    67 B/op |
| [lestrrat/go-jspointer](https://github.com/lestrrat/go-jspointer) |  1271 ns/op |  9 allocs/op |   272 B/op |

## Run

```go
go test -bench . -benchmem
```

## License

Copyright 2016 Olivier Mengué

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
