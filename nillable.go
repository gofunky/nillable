package nillable

type (
	// Nillable is to be implemented by all structs that shall circumvent the reflective nil check.
	Nillable interface {
		// Returns true if the interface's struct object is nil.
		IsNil() bool
	}
)
