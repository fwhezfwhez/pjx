package main

import "strings"

func FormatPath(s string) string {
	return strings.Replace(s, "\\", "/", -1)
}
