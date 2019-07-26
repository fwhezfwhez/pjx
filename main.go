package main

import (
	"fmt"
	"os"
	"path"
	"runtime/debug"
	"strings"
)

var pj Pjx

func main() {
	fmt.Println(1)
	//// set pj gomodule
	//flag.BoolVar(&pj.GOModule, "gomodule", false, "pjx -gomodule true")
	//flag.Parse()
	//// set appPath
	//var e error
	//pj.AppPath, e = os.Getwd()
	//if e != nil {
	//	panic(e)
	//}
	//
	//args := os.Args[1:]
	//
	//switch args[0] {
	//case "version", "--version":
	//	fmt.Println(pj.Version())
	//case "help", "--help":
	//	fmt.Println(pj.Usage())
	//case "new":
	//	appName := args[1]
	//	newApp(appName)
	//case "module":
	//	moduleName := args[1]
	//	newModule(moduleName)
	//case "add":
	//	// pjx add db fwhezfwhez master
	//	// pjx add db [global master]
	//	directoryName := args[1]
	//	var namespace string
	//	var tag string
	//	switch len(args) {
	//	// pjx add db fwhezfwhez master
	//	case 4:
	//		namespace = args[2]
	//		tag = args[3]
	//	case 2:
	//		namespace = "global"
	//		tag = "master"
	//	default:
	//		fmt.Println(fmt.Sprintf("bad args num,want 'pjx add <pkgName> [namespace] [tagName]' but got '%s'", strings.Join(os.Args, " ")))
	//		return
	//	}
	//	addPkg(directoryName, namespace, tag)
	//}

}

func newApp(appName string) {
	dirPath, e := os.Getwd()
	if e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}

	if e := os.Mkdir(path.Join(dirPath, appName), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}

	appPath := path.Join(dirPath, appName)
	if e := os.Mkdir(path.Join(appPath, "config"), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
	if e := os.Mkdir(path.Join(appPath, "module"), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
	if e := os.Mkdir(path.Join(appPath, "dependence"), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
	if e := os.Mkdir(path.Join(appPath, "independence"), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
	if f, e := os.Create(path.Join(appPath, "main.go")); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	} else {
		f.Close()
	}

	if f, e := os.OpenFile(path.Join(appPath, "main.go"), os.O_CREATE|os.O_RDWR, os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	} else {
		defer f.Close()
		_, e := f.Write(tmplMainGo)
		if e != nil {
			fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
			return
		}
	}

	if f, e := os.OpenFile(path.Join(appPath, "config", "local.yaml"), os.O_CREATE|os.O_RDWR, os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	} else {
		defer f.Close()
		_, e := f.Write(tmplConfigYAML)
		if e != nil {
			fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
			return
		}
	}

	if f, e := os.OpenFile(path.Join(appPath, "config", "init.go"), os.O_CREATE|os.O_RDWR, os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	} else {
		defer f.Close()
		_, e := f.Write(GetBasicTmplConfigYAMLGO(strings.Replace(path.Join(appPath, "config", "local.yaml"), "\\", "/", -1)))
		if e != nil {
			fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
			return
		}
	}
}

func newModule(moduleName string) {
	if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}

	if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Pb")), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
	if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Model")), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
	if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Router")), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
	if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Service")), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}

	if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "TestClient")), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}

	if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Export")), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}
}

func addPkg(directoryName string, namespace string, tag string) {
	execDir, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	dirPath := FormatPath(path.Join(execDir, directoryName))
	fileInfo, e := os.Stat(dirPath)
	if e != nil {
		if os.IsNotExist(e) {
			logger.Println(fmt.Sprintf("pkg not found, '%s' is not found", dirPath))
			return
		}
		panic(e)
	}
	if !fileInfo.IsDir() {
		logger.Println(fmt.Sprintf("only support directory, '%s' is not a directory", dirPath))
		return
	}

	pjxPath := os.Getenv("pjx_path")
	if pjxPath == "" {
		logger.Println(fmt.Sprintf("pkgPath not found in path, make sure it's well in system env."))
		return
	}
	fileInfo, e = os.Stat(dirPath)
	if !fileInfo.IsDir() {
		logger.Println(fmt.Sprintf("pjx_path is not a  directory, '%s' is not a directory", pjxPath))
		return
	}

	if tag == "master" {
		copyDir(dirPath, FormatPath(path.Join(pjxPath, namespace, directoryName)))
	} else {
		copyDir(dirPath, FormatPath(path.Join(pjxPath, namespace, directoryName+"@"+tag)))
	}
    return
}

// sync all pkg and pjx.json in os env 'pjx_path'.
func syncAll() {

}

// sync one pkg and pjx.json in os env 'pjx_path'.
func syncOne() {
}
