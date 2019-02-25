package benchmark

import (
	jsonpointer "github.com/json-validate/json-pointer-go"
)

type JsonValidateJsonPointerGo struct{}

func (JsonValidateJsonPointerGo) Get(doc interface{}, pointer string) (interface{}, error) {
	ptr, err := jsonpointer.New(pointer)
	if err != nil {
		return nil, err
	}
	element, err := ptr.Eval(doc)
	if err != nil {
		return nil, err
	}
	return *element, nil
}

func (JsonValidateJsonPointerGo) Parse(pointer string) (Stringer, error) {
	return jsonpointer.New(pointer)
}
