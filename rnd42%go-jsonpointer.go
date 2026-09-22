package benchmark

import (
	"github.com/rnd42/go-jsonpointer"
)

type Rnd42JSONPointer struct{}

func (Rnd42JSONPointer) Get(doc any, pointer string) (any, error) {
	ptr, err := jsonpointer.NewJSONPointerFromString(pointer)
	if err != nil {
		return nil, err
	}
	return ptr.Get(doc, -1)
}

func (Rnd42JSONPointer) Set(pdoc *any, pointer string, value any) error {
	ptr, err := jsonpointer.NewJSONPointerFromString(pointer)
	if err != nil {
		return err
	}
	*pdoc, err = ptr.Set(*pdoc, value, -1)
	return err
}

func (Rnd42JSONPointer) Parse(pointer string) (Stringer, error) {
	return jsonpointer.NewJSONPointerFromString(pointer)
}
