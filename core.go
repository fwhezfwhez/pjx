package main

import (
	"os"
	"strings"
)

type Pjx struct {
	GOModule bool
	AppPath  string

	// receive args with value like -o xxx, -m xxx
	KV map[string]string

	// receive -l, whether open log
	IfLog    bool
	// receive -f, whether do command by force
	IfForce  bool

	// receive -o flag, in different command, it refers to different meaning, like pjx use db -o db2,
	// fetch a db from global master and put it into current path and name it db2.
	O string
}

func (p *Pjx) Version() string {
	return "v2.2.6"
}
func (p *Pjx) Usage() string {
	var usages = []string{
		"commands:",

		"# project directory:",
		"\tpjx version                   show pjx version",
		"\tpjx help                      show available usage",
		"\tpjx new <appName>             add a new app project and specific its name",
		"\tpjx module <moduleName>       add a new module in an existed app",
		"",
		"# package management:",
		"\tpjx add <pkgName> [namespace] [tag]            add a current package into pjx local repo with specific namespace and tag, if not set,use global master by default",
		"\tpjx use <pkgName> [namespace] [tag]            use a package from pjx local repo and pu into current dir, if not set namespace and tag, use global master by default",
		"\tpjx merge <path> <namespace> [-f/-u]           merge packages from path into repo/namespace",
		"\tpjx clone <url.git> <namespace> [-f/-u]        clone and merge from remote git repo",
		"",
		"example:",
		"\tpjx new helloWorld                         new an app named helloWorld",
		"\tcd helloWorld                              cd into helloWorld directory",
		"\tpjx new module user                        new a module named user",
		"\tpjx add helloWord                          add helloWord package into repo tag global master",
		"\tpjx use helloWord -o helloWorld2           fetch helloWorld package from repo and rename it helloWorld2",
		"\tpjx merge /home/web/repo/ fwhezfwhez -u    merge pkg from /home/web/repo/ into ${pjx_path}/fwhezfwhez",
		"\tpjx clone https://github.com/fwhezfwhez/pjx-repo.git fwhezfwhez -u clone from url.git and merge into repo",
	}
	return strings.Join(usages, "\n")
}

func (p *Pjx) ImportPrefix() string {
	if p.GOModule {
		return ""
	}
	return FormatPath(os.Getenv("GOPATH"))
}
