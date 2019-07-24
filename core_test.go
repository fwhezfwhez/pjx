package main

import (
	"fmt"
	"testing"
)

func TestPjx_ImportPrefix(t *testing.T) {
	var p = Pjx{}
	fmt.Println(p.ImportPrefix())
}
