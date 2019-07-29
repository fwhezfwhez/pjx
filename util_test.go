package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFormatPath(t *testing.T) {
	fmt.Println(os.Getenv("pkg_path"))
}

func TestRmAttach(t *testing.T) {
	fmt.Println(rmAttach([]string{"pjx", "use", "hello", "-o", "hello2"}))
}

func TestDelDir(t *testing.T) {
	DelDir("G:\\go_workspace\\GOPATH\\src\\pjx\\hello")
}

func TestPathJoin(t *testing.T) {
	fmt.Println(PathJoin([]string{"1","empty","2"}...))
}

func TestSetPjxEnv(t *testing.T) {
	SetPjxEnv()
	fmt.Println(os.Getenv("pjx_path"))
}
