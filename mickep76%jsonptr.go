package benchmark

import (
	"github.com/mickep76/jsonptr"
)

type Mickep76JsonPtr struct{}

func (Mickep76JsonPtr) Get(doc interface{}, pointer string) (interface{}, error) {
	return jsonptr.Resolve(doc, pointer)
}
