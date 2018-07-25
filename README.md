# nillable
Common Nillable Interface for Golang

```go
func (t *YourStruct) IsNil() bool {
	if (t == nil) {
		return true;
	}
	return false;
}

func DoSomething(n Nillable) {
    if (!n.IsNil()) {
        // Awesome!
    }
}
```
