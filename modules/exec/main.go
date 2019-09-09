// Do not use.
// On dev.
package main

import (
	"fmt"
	"io"

	//"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("G:\\go_workspace\\GOPATH\\src\\pjx\\modules\\exec\\exec")

	stdin, e := cmd.StdinPipe()
	if e != nil {
		panic(e)
	}

	stdout, e := cmd.StdoutPipe()
	if e != nil {
		panic(e)
	}
	if e := cmd.Start(); e != nil {
		panic(e)
	}
	_, e = stdin.Write([]byte("hello\n"))
	_, e = stdin.Write([]byte("world\n"))
	if e != nil {
		panic(e)
	}
	stdin.Close()

	//out, _ := ioutil.ReadAll(stdout)
	// or you can use a loop
	for {
		var buf = make([]byte, 512)
		n, e := stdout.Read(buf)
		if e == io.EOF {
			break
		}
		if e != nil {
			panic(e)
		}
		fmt.Println(string(buf[:n]))
	}
	stdout.Close()
}
