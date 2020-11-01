package benchmark

import (
	jsonpointer "github.com/go-openapi/jsonpointer"
)

type GoOpenAPIJSONPointer struct{}

func (GoOpenAPIJSONPointer) Get(doc interface{}, pointer string) (interface{}, error) {
	ptr, err := jsonpointer.New(pointer)
	if err != nil {
		return nil, err
	}
	element, _, err := ptr.Get(doc)
	if err != nil {
		return nil, err
	}
	return element, nil
}

func (GoOpenAPIJSONPointer) Parse(pointer string) (Stringer, error) {
	ptr, err := jsonpointer.New(pointer)
	if err != nil {
		return nil, err
	}
	return &ptr, err
}

func (GoOpenAPIJSONPointer) Set(document *interface{}, pointer string, value interface{}) error {
	ptr, err := jsonpointer.New(pointer)
	if err != nil {
		return err
	}
	newDoc, err := ptr.Set(*document, value)
	if err != nil {
		return err
	}
	*document = newDoc
	return nil
}
