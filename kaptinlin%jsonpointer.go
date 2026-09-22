package benchmark

import (
	"github.com/kaptinlin/jsonpointer"
)

type KaptinlinJSONPointer struct{}

func (KaptinlinJSONPointer) Get(doc any, pointer string) (any, error) {
	return jsonpointer.Value(doc, pointer)
}

func (KaptinlinJSONPointer) Parse(pointer string) (Stringer, error) {
	p, err := jsonpointer.Parse(pointer)
	if err != nil {
		return nil, err
	}
	return p, nil
}
