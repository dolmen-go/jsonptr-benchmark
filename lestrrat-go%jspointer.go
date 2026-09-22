package benchmark

import (
	"github.com/lestrrat-go/jspointer"
)

type LestrratGoJsPointer struct{}

func (LestrratGoJsPointer) Get(doc any, pointer string) (any, error) {
	ptr, err := jspointer.New(pointer)
	if err != nil {
		return nil, err
	}
	return ptr.Get(doc)
}

func (LestrratGoJsPointer) Set(pdoc *any, pointer string, value any) error {
	ptr, err := jspointer.New(pointer)
	if err != nil {
		return err
	}
	if pointer == "" {
		return errSetRoot
	}
	return ptr.Set(*pdoc, value)
}

func (LestrratGoJsPointer) Parse(pointer string) (Stringer, error) {
	return jspointer.New(pointer)
}
