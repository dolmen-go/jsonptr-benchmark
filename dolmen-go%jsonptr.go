package benchmark

import (
	"github.com/dolmen-go/jsonptr"
)

type DolmenGoJsonPtr struct{}

func (d DolmenGoJsonPtr) Get(doc interface{}, pointer string) (interface{}, error) {
	return jsonptr.Get(doc, pointer)
}

func (d DolmenGoJsonPtr) Set(pdoc *interface{}, pointer string, value interface{}) error {
	return jsonptr.Set(pdoc, pointer, value)
}

func (d DolmenGoJsonPtr) Parse(pointer string) (Stringer, error) {
	return jsonptr.Parse(pointer)
}
