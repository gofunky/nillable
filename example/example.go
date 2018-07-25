package example

import "github.com/gofunky/nillable"

type TestStruct struct {
	something string
}

func (t *TestStruct) IsNil() bool {
	if t == nil {
		return true
	}
	return false
}

var _ nillable.Nillable = (*TestStruct)(nil)
