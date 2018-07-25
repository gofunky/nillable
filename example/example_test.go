package example

import (
	"testing"
)

func TestTestStruct_IsNil(t *testing.T) {
	tests := []struct {
		name       string
		testStruct *TestStruct
		want       bool
	}{
		{
			name: "Nil Struct",
			want: true,
		},
		{
			name:       "Non-Nil Struct",
			testStruct: &TestStruct{},
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.testStruct.IsNil(); got != tt.want {
				t.Errorf("TestStruct.IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
