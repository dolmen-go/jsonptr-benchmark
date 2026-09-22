package benchmark

import (
	"github.com/chanced/jsonpointer"
)

type ChancedJsonPointer struct{}

func (ChancedJsonPointer) Get(doc any, pointer string) (out any, err error) {
	err = jsonpointer.Resolve(doc, jsonpointer.Pointer(pointer), &out)
	return
}

func (ChancedJsonPointer) Set(pdoc *any, pointer string, value any) error {
	return jsonpointer.Assign(pdoc, jsonpointer.Pointer(pointer), value)
}

func (ChancedJsonPointer) Parse(pointer string) (Stringer, error) {
	return jsonpointer.Parse(pointer)
}
