package benchmark

import (
	"github.com/lestrrat/go-jspointer"
)

type LestrratGoJsPointer struct{}

func (d LestrratGoJsPointer) Get(doc interface{}, pointer string) (interface{}, error) {
	ptr, err := jspointer.New(pointer)
	if err != nil {
		return nil, err
	}
	return ptr.Get(doc)
}

func (d LestrratGoJsPointer) Set(pdoc *interface{}, pointer string, value interface{}) error {
	ptr, err := jspointer.New(pointer)
	if err != nil {
		return err
	}
	if pointer == "" {
		return errSetRoot
	}
	return ptr.Set(*pdoc, value)
}

func (d LestrratGoJsPointer) Parse(pointer string) (Stringer, error) {
	return jspointer.New(pointer)
}
