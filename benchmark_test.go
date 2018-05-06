package benchmark

import (
	"testing"
)

type GetImpl interface {
	Get(document interface{}, pointer string) (interface{}, error)
}

type SetImpl interface {
	GetImpl
	Set(document *interface{}, pointer string, value interface{}) error
}

var implementations = map[string]GetImpl{
	"dolmen-go/jsonptr":     DolmenGoJsonPtr{},
	"xeipuuv/gojsonpointer": XeipuuvGoJsonPointer{},
	"mickep76/jsonptr":      Mickep76JsonPtr{},
	"lestrrat/go-jspointer": LestrratGoJsPointer{},
	"dustin/go-jsonpointer": DustinGoJsonPointer{},
	"rnd42/go-jsonpointer":  Rnd42JSONPointer{},
	"twindagger/jsonptr":    TwindaggerJSONPtr{},
	"qri-io/jsonpointer":    QriIoJSONPointer{},
}

type PointerParser interface {
	Parse(pointer string) (Stringer, error)
}

var parsers = make(map[string]PointerParser, len(implementations))

func init() {
	for name, impl := range implementations {
		if impl, ok := impl.(PointerParser); ok {
			parsers[name] = impl
		}
	}
}

func BenchmarkGet(b *testing.B) {
	doc := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": []interface{}{
				true,
				nil,
				false,
			},
		},
	}
	const ptr = "/foo/bar/2"
	for name, impl := range implementations {
		// Check the implementation
		res, err := impl.Get(doc, ptr)
		if err != nil {
			b.FailNow()
		}
		if res != false {
			b.FailNow()
		}

		b.Run("["+name+"]", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, err = impl.Get(doc, ptr)
			}
		})
	}
}

var pointers = []string{
	// "/foo/bar",
	// "/foo/bar/baz/~0~1",
	"/definitions/Location",
	"/path/~1home~1dolmen",
	"/path/~0~1dolmen",
}

func benchmarkParsers(b *testing.B, f func(func(string) (Stringer, error), string) func(*testing.B)) {
	for _, pointer := range pointers {
		b.Run(`"`+pointer+`"`, func(b *testing.B) {
			for name, p := range parsers {
				f2 := f(p.Parse, pointer)
				if f2 == nil {
					continue
				}
				b.Run("["+name+"]", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						f2(b)
					}
				})
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	benchmarkParsers(b, func(Parse func(string) (Stringer, error), pointer string) func(*testing.B) {
		// Returns the body of the benchmark: parsing the given pointer
		return func(*testing.B) {
			_, _ = Parse(pointer)
		}
	})
}

func BenchmarkBackToString(b *testing.B) {
	benchmarkParsers(b, func(Parse func(string) (Stringer, error), pointer string) func(*testing.B) {
		// Parse the pointer, but before starting the benchmark
		ptr, err := Parse(pointer)
		if err != nil {
			b.FailNow()
		}
		// Returns the body of the benchmark: formatting the given pointer
		return func(b *testing.B) {
			out := ptr.String()
			if out != pointer {
				b.FailNow()
			}
		}
	})
}

func BenchmarkParseAndBackToString(b *testing.B) {
	benchmarkParsers(b, func(Parse func(string) (Stringer, error), pointer string) func(*testing.B) {
		// Parse the pointer, but before starting the benchmark
		// Returns the body of the benchmark: formatting the given pointer
		return func(b *testing.B) {
			ptr, err := Parse(pointer)
			if err != nil {
				b.FailNow()
			}
			out := ptr.String()
			if out != pointer {
				b.FailNow()
			}
		}
	})
}
