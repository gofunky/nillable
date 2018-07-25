# nillable
[![Build Status](https://travis-ci.org/gofunky/nillable.svg?branch=master)](https://travis-ci.org/gofunky/nillable)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofunky/nillable)](https://goreportcard.com/report/github.com/gofunky/nillable)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/442f9e83f9594093823ac39d341bfca2)](https://www.codacy.com/app/gofunky/nillable?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gofunky/nillable&amp;utm_campaign=Badge_Grade)

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
