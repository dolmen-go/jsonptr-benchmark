package benchmark

import "errors"

var errSetRoot = errors.New("This implementation doesn't handle replacing the document root")

// same as fmt.Stringer
type Stringer interface {
	String() string
}
