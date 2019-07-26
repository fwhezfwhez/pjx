package main

import (
	"os"
	"strings"
)

type Pjx struct {
	GOModule bool
	AppPath  string
}

func (p *Pjx) Version() string {
	return "v2.0.0"
}
func (p *Pjx) Usage() string {
	var usages = []string{
		"commands:",

		"\tpjx version                   show pjx version",
		"\tpjx help                    show available usage",
		"\tpjx new <appName>             add a new app project and specific its name",
		"\tpjx module <moduleName>       add a new module in an existed app",
		"example:",

		"\tpjx new helloWorld            new an app named helloWorld",
		"\tcd helloWorld                 cd into helloWorld directory",
		"\tpjx new module user           new a module named user",
	}
	return strings.Join(usages, "\n")
}

func (p *Pjx) ImportPrefix() string {
	if p.GOModule {
		return ""
	}
	return FormatPath(os.Getenv("GOPATH"))
}
