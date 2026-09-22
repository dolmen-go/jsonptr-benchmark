package benchmark

import (
	"encoding/json/jsontext"

	"github.com/crhntr/jsonschema/jsonptr"
)

type CrhntrJsonSchemaJsonPtr struct{}

func (CrhntrJsonSchemaJsonPtr) Get(doc any, pointer string) (any, error) {
	_, value, err := jsonptr.FindValue(jsonptr.Pointer(pointer), doc)
	return value, err
}

func (CrhntrJsonSchemaJsonPtr) GetRawBytes(doc []byte, pointer string) (any, error) {
	_, value, err := jsonptr.FindValue(jsonptr.Pointer(pointer), jsontext.Value(doc))
	return value, err
}

type crhntrPtr jsonptr.Pointer

func (p crhntrPtr) String() string {
	return string(p)
}

func (CrhntrJsonSchemaJsonPtr) Parse(pointer string) (Stringer, error) {
	p := jsonptr.Pointer(pointer)
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return crhntrPtr(p), nil
}
