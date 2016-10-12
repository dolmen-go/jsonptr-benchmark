package benchmark

import (
	"github.com/xeipuuv/gojsonpointer"
)

type XeipuuvGoJsonPointer struct{}

func (d XeipuuvGoJsonPointer) Get(doc interface{}, pointer string) (interface{}, error) {
	ptr, err := gojsonpointer.NewJsonPointer(pointer)
	if err != nil {
		return nil, err
	}
	res, _, err := ptr.Get(doc)
	return res, err
}

func (d XeipuuvGoJsonPointer) Set(pdoc *interface{}, pointer string, value interface{}) error {
	ptr, err := gojsonpointer.NewJsonPointer(pointer)
	if err != nil {
		return err
	}
	if pointer == "" {
		return errSetRoot
	}
	_, err = ptr.Set(*pdoc, value)
	return err
}

func (d XeipuuvGoJsonPointer) Parse(pointer string) (Stringer, error) {
	ptr, err := gojsonpointer.NewJsonPointer(pointer)
	return &ptr, err
}
