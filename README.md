
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

### 2018-03-01 (go1.10 linux/amd64)

#### Get

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | **156 ns/op** | **16 B/op** | **1 allocs/op** |
| dustin/go-jsonpointer | 221 ns/op | 48 B/op | **1 allocs/op** |
| lestrrat/go-jspointer | 910 ns/op | 368 B/op | 9 allocs/op |
| mickep76/jsonptr | 360 ns/op | 64 B/op | **1 allocs/op** |
| rnd42/go-jsonpointer | 559 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 519 ns/op | 144 B/op | 3 allocs/op |
| xeipuuv/gojsonpointer | 547 ns/op | 144 B/op | 3 allocs/op |

#### Parse `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| **dolmen-go/jsonptr** | 196 ns/op | **64 B/op** | **2 allocs/op** |
| lestrrat/go-jspointer | 403 ns/op | 240 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 482 ns/op | 128 B/op | 3 allocs/op |
| twindagger/jsonptr | 423 ns/op | 112 B/op | 3 allocs/op |
| **xeipuuv/gojsonpointer** | **184 ns/op** | **64 B/op** | **2 allocs/op** |

#### Parse `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 296 ns/op | 80 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 394 ns/op | 240 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 762 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 565 ns/op | 144 B/op | 5 allocs/op |
| **xeipuuv/gojsonpointer** | **186 ns/op** | **64 B/op** | **2 allocs/op** |

#### Parse `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 285 ns/op | 72 B/op | 3 allocs/op |
| lestrrat/go-jspointer | 361 ns/op | 224 B/op | 4 allocs/op |
| rnd42/go-jsonpointer | 816 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 627 ns/op | 160 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **183 ns/op** | **64 B/op** | **2 allocs/op** |

#### BackToString `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 169 ns/op | 80 B/op | 3 allocs/op |
| **lestrrat/go-jspointer** | 11.2 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **8.23 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 423 ns/op | 112 B/op | 4 allocs/op |
| xeipuuv/gojsonpointer | 148 ns/op | 64 B/op | 2 allocs/op |

#### BackToString `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 190 ns/op | 80 B/op | 3 allocs/op |
| **lestrrat/go-jspointer** | 11.2 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **8.44 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 548 ns/op | 144 B/op | 6 allocs/op |
| xeipuuv/gojsonpointer | 149 ns/op | 64 B/op | 2 allocs/op |

#### BackToString `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 132 ns/op | 32 B/op | 2 allocs/op |
| **lestrrat/go-jspointer** | 11.2 ns/op | **0 B/op** | **0 allocs/op** |
| **rnd42/go-jsonpointer** | **8.36 ns/op** | **0 B/op** | **0 allocs/op** |
| twindagger/jsonptr | 619 ns/op | 144 B/op | 8 allocs/op |
| xeipuuv/gojsonpointer | 142 ns/op | 32 B/op | 2 allocs/op |

#### ParseAndBackToString `"/definitions/Location"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 379 ns/op | 144 B/op | 5 allocs/op |
| lestrrat/go-jspointer | 457 ns/op | 240 B/op | 4 allocs/op |
| **rnd42/go-jsonpointer** | 488 ns/op | **128 B/op** | **3 allocs/op** |
| twindagger/jsonptr | 862 ns/op | 224 B/op | 7 allocs/op |
| **xeipuuv/gojsonpointer** | **340 ns/op** | **128 B/op** | 4 allocs/op |

#### ParseAndBackToString `"/path/\~1home\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 492 ns/op | 160 B/op | 6 allocs/op |
| lestrrat/go-jspointer | 412 ns/op | 240 B/op | **4 allocs/op** |
| rnd42/go-jsonpointer | 780 ns/op | 160 B/op | 5 allocs/op |
| twindagger/jsonptr | 1151 ns/op | 288 B/op | 11 allocs/op |
| **xeipuuv/gojsonpointer** | **344 ns/op** | **128 B/op** | **4 allocs/op** |

#### ParseAndBackToString `"/path/\~0\~1dolmen"`

| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
| dolmen-go/jsonptr | 424 ns/op | 104 B/op | 5 allocs/op |
| lestrrat/go-jspointer | 374 ns/op | 224 B/op | **4 allocs/op** |
| rnd42/go-jsonpointer | 823 ns/op | 160 B/op | 7 allocs/op |
| twindagger/jsonptr | 1257 ns/op | 304 B/op | 15 allocs/op |
| **xeipuuv/gojsonpointer** | **322 ns/op** | **96 B/op** | **4 allocs/op** |


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
