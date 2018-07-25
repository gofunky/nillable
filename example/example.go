package example

import "github.com/gofunky/nillable"

// TestStruct is a simple example struct
type TestStruct struct {
	something string
}

// IsNil returns true if the TestStruct is nil
func (t *TestStruct) IsNil() bool {
	if t == nil {
		return true
	}
	return false
}

// Make sure that the Nillable interface is implemented
var _ nillable.Nillable = (*TestStruct)(nil)
