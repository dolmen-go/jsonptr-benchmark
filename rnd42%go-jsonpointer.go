package benchmark

import (
	"github.com/rnd42/go-jsonpointer"
)

type Rnd42JSONPointer struct{}

func (d Rnd42JSONPointer) Get(doc interface{}, pointer string) (interface{}, error) {
	ptr, err := jsonpointer.NewJSONPointerFromString(pointer)
	if err != nil {
		return nil, err
	}
	return ptr.Get(doc, -1)
}

func (d Rnd42JSONPointer) Set(pdoc *interface{}, pointer string, value interface{}) error {
	ptr, err := jsonpointer.NewJSONPointerFromString(pointer)
	if err != nil {
		return err
	}
	*pdoc, err = ptr.Set(*pdoc, value, -1)
	return err
}

func (d Rnd42JSONPointer) Parse(pointer string) (Stringer, error) {
	return jsonpointer.NewJSONPointerFromString(pointer)
}
