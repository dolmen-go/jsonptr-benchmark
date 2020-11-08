package benchmark

import (
	"github.com/yukithm/json2csv/jsonpointer"
)

type YukithmJSON2CSVJsonPointer struct{}

func (YukithmJSON2CSVJsonPointer) Get(doc interface{}, pointer string) (interface{}, error) {
	return jsonpointer.Get(doc, pointer)
}

func (YukithmJSON2CSVJsonPointer) Parse(pointer string) (Stringer, error) {
	return jsonpointer.New(pointer)
}
