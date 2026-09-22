package benchmark

import (
	"encoding/json/jsontext"
	"errors"
)

type StdJsonTextPtr struct{}

func (StdJsonTextPtr) Get(doc any, pointer string) (any, error) {
	panic("not implemented!")
}

type jsontextPtr jsontext.Pointer

func (p jsontextPtr) String() string {
	return string(p)
}

func (StdJsonTextPtr) Parse(pointer string) (Stringer, error) {
	ok := jsontext.Pointer(pointer).IsValid()
	if !ok {
		return nil, errors.New("invalid")
	}
	return jsontextPtr(pointer), nil
}
