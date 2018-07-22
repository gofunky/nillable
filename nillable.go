package nillable

type (
	// Implement this interface for all structs that shall circumvent the reflective nil check.
	Nillable interface {
		// Returns true if the interface's struct object is nil.
		IsNil() bool
	}
)
