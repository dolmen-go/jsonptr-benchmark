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
	for name, impl := range implementations {
		b.Run("["+name+"]", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, err := impl.Get(doc, "/foo/bar/2")
				if err != nil {
					b.FailNow()
				}
				if res != false {
					b.FailNow()
				}
			}
		})
	}
}
