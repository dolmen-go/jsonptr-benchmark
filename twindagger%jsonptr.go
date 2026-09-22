package benchmark

import (
	"github.com/twindagger/jsonptr"
)

type TwindaggerJSONPtr struct{}

func (TwindaggerJSONPtr) Get(doc any, pointer string) (any, error) {
	return jsonptr.Get(doc, pointer)
}

func (TwindaggerJSONPtr) Set(pdoc *any, pointer string, value any) error {
	if pointer == "" {
		return errSetRoot
	}
	return jsonptr.Set(*pdoc, pointer, value)
}

func (TwindaggerJSONPtr) Parse(pointer string) (Stringer, error) {
	return jsonptr.New(pointer)
}
