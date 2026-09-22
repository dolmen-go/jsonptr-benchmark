package benchmark

import (
	"github.com/dustin/go-jsonpointer"
)

type DustinGoJsonPointer struct{}

func (DustinGoJsonPointer) Get(doc any, pointer string) (any, error) {
	switch doc := doc.(type) {
	case map[string]any:
		return jsonpointer.Get(doc, pointer), nil
	default:
		// Wrap the document in a fake object to workaround the flawed interface
		return jsonpointer.Get(map[string]any{
			"": doc,
		}, "/"+pointer), nil
	}
}
