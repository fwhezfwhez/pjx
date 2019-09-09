package main

import (
	"fmt"
	"os"
	"strings"
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

func TestCopyDirF(t *testing.T) {
    // CopyDirU(`G:\go_workspace\GOPATH\src\pjx\helo`, `G:\go_workspace\GOPATH\src\pjx\helo2`)
}
func TestGetGitName(t *testing.T) {
	fmt.Println(GetGitName("https://ffff/ffffff/fff.git"))

	fmt.Println(strings.HasPrefix(`// hello`, `//`))
}

func TestProfileCommands(t *testing.T) {
	rs :=profileCommand([]string{
		"hello1 \\",
		"hello2",
		"hello3",
		"hello4 \\",
		"hello5 \\",
		"hello6",
	})
	fmt.Println(len(rs))
	fmt.Println(rs)
}
