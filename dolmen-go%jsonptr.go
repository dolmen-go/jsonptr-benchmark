package benchmark

import (
	"encoding/json"

	"github.com/dolmen-go/jsonptr"
)

type DolmenGoJsonPtr struct{}

func (DolmenGoJsonPtr) Get(doc any, pointer string) (any, error) {
	return jsonptr.Get(doc, pointer)
}

func (DolmenGoJsonPtr) GetRawBytes(doc []byte, pointer string) (any, error) {
	return jsonptr.Get(json.RawMessage(doc), pointer)
}

func (DolmenGoJsonPtr) Set(pdoc *any, pointer string, value any) error {
	return jsonptr.Set(pdoc, pointer, value)
}

func (DolmenGoJsonPtr) Parse(pointer string) (Stringer, error) {
	return jsonptr.Parse(pointer)
}
