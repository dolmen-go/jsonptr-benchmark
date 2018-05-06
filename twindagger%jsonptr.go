package benchmark

import (
	"github.com/twindagger/jsonptr"
)

type TwindaggerJSONPtr struct{}

func (TwindaggerJSONPtr) Get(doc interface{}, pointer string) (interface{}, error) {
	return jsonptr.Get(doc, pointer)
}

func (TwindaggerJSONPtr) Set(pdoc *interface{}, pointer string, value interface{}) error {
	if pointer == "" {
		return errSetRoot
	}
	return jsonptr.Set(*pdoc, pointer, value)
}

func (TwindaggerJSONPtr) Parse(pointer string) (Stringer, error) {
	return jsonptr.New(pointer)
}
