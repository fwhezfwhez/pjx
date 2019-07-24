package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFormatPath(t *testing.T) {

	dr, _ := os.Getwd()

	fmt.Println(dr)
	fmt.Println(FormatPath(dr))
	fmt.Println(FormatPath("/xx/h\\yu\\jk"))
}
