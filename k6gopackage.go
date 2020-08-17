package k6gopackage

import (
	"context"
	// more imports
	"strings"
)

type Obj struct{}

func New() *Obj {
	return &Obj{}
}

// Method names must begin with a capital letter
func (o *Obj) Method(ctx context.Context, arg1 string) string {
	// do something with arg1 and return a string
	return strings.ToLower(arg1)
}
