package main

import (
	"fmt"
	"os"
	"testing"
)

func TestF(t *testing.T) {
	// get running dir path
	fmt.Println(os.Getwd())
}

func TestNewApp(t *testing.T) {
	newApp("helloWorld")
}

func TestNewModule(t *testing.T) {
	newModule("user")
}
