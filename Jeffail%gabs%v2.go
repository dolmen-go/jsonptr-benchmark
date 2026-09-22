package benchmark

import (
	"github.com/Jeffail/gabs/v2"
)

type JeffailGabs struct{}

func (JeffailGabs) Get(doc any, pointer string) (any, error) {
	c, err := gabs.Wrap(doc).JSONPointer(pointer)
	if err != nil {
		return nil, err
	}
	return c.Data(), nil
}

func (JeffailGabs) GetRawBytes(doc []byte, pointer string) (any, error) {
	c, err := gabs.ParseJSON(doc)
	if err != nil {
		return nil, err
	}
	c, err = c.JSONPointer(pointer)
	if err != nil {
		return nil, err
	}
	return c.Data(), nil
}

func (JeffailGabs) Set(pdoc *any, pointer string, value any) error {
	if pointer == "" {
		return errSetRoot
	}
	_, err := gabs.Wrap(*pdoc).SetJSONPointer(value, pointer)
	return err
}
