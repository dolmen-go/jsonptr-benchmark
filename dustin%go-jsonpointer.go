package benchmark

import (
	"github.com/dustin/go-jsonpointer"
)

type DustinGoJsonPointer struct{}

func (DustinGoJsonPointer) Get(doc interface{}, pointer string) (interface{}, error) {
	switch doc := doc.(type) {
	case map[string]interface{}:
		return jsonpointer.Get(doc, pointer), nil
	default:
		// Wrap the document in a fake object to workaround the flawed interface
		return jsonpointer.Get(map[string]interface{}{
			"": doc,
		}, "/"+pointer), nil
	}
}
