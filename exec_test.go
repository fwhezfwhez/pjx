package main

import (
	"fmt"
	"testing"
)

func TestEleminateAnno(t *testing.T) {
	fmt.Println(string(eleminateAnotation([]byte(`
// hello
echo hello
`))))
}
