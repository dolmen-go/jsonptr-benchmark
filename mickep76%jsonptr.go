package benchmark

import (
	"github.com/mickep76/jsonptr"
)

type Mickep76JsonPtr struct{}

func (d Mickep76JsonPtr) Get(doc interface{}, pointer string) (interface{}, error) {
	return jsonptr.Resolve(doc, pointer)
}
