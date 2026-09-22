package benchmark

import (
	"github.com/qri-io/jsonpointer"
)

type QriIoJSONPointer struct{}

func (QriIoJSONPointer) Get(doc any, pointer string) (any, error) {
	ptr, err := jsonpointer.Parse(pointer)
	if err != nil {
		return nil, err
	}
	return ptr.Eval(doc)
}

func (QriIoJSONPointer) Parse(pointer string) (Stringer, error) {
	return jsonpointer.Parse(pointer)
}
