package benchmark

import (
	jsonpointer "github.com/oas3/json-pointer"
)

type OAS3JSONPointer struct{}

func (OAS3JSONPointer) Get(doc interface{}, pointer string) (interface{}, error) {
	ptr, err := jsonpointer.New(pointer)
	if err != nil {
		return nil, err
	}
	return ptr.Get(doc)
}

func (OAS3JSONPointer) Parse(pointer string) (Stringer, error) {
	ptr, err := jsonpointer.New(pointer)
	if err != nil {
		return nil, err
	}
	return &ptr, err
}

func (OAS3JSONPointer) Set(document *interface{}, pointer string, value interface{}) error {
	ptr, err := jsonpointer.New(pointer)
	if err != nil {
		return err
	}
	newDoc, _, err := ptr.Set(*document, value)
	if err != nil {
		return err
	}
	*document = newDoc
	return nil
}
