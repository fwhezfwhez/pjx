package main

import "strings"

var commandPrefix = "- "
var osPrefix = []string{"? linux", "? darwin", "? windows"}
var stopOnFailPrefix = "! "
var stdinPrefix = "  -"

type Command struct {
	Line int

	Stdin   []string
	Command string
}

//func Decode(buf []byte) ([]Command, error) {
//	var rs = make([]Command, 0, 10)
//
//	arr := strings.Split(string(buf), "\n")
//
//	for i, v := range arr {
//		if strings.Trim(v, " ") == "#" {
//			continue
//		}
//		if strings.HasPrefix(v, "//") {
//			continue
//		}
//		if strings.HasPrefix(v, "--") {
//			continue
//		}
//
//		// handle -
//		if n := strings.Index(v, commandPrefix); n != -1 {
//			var command Command
//			command.Line = i
//			command.Command = string(v[n+len(commandPrefix):])
//
//			command.Command = handleReplace(command.Command)
//			command.Stdin, i = stdinOf(arr, i, v)
//			continue
//		}
//		// handle ? <os>
//
//
//	}
//
//	return rs, nil
//}

func eleminateAnotation(buf []byte) []byte {
	tmp := strings.Split(string(buf), "\n")
	var rs []string
	for _, v := range tmp {
		if !strings.HasPrefix(v, `//`) {
			rs = append(rs, v)
		}
	}
	return []byte(strings.Join(rs, "\n"))
}
