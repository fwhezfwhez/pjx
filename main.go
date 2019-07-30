package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	//"os/exec"
	"os/user"
	"path"
	"runtime/debug"
	"strings"
)

var pj Pjx

func init() {
	if os.Getenv("pjx_path") == "" {
		us, e := user.Current()
		if e != nil {
			panic(e)
		}
		if IfLog(os.Args) {
			logger.Println(fmt.Sprintf("pjx_path not found, pjx has auto use its process-scope env '%s', it will not effect your system env setting.But it's advised to set 'pjx_path' properly to let pjx know where to put packages locally.", PathJoin(us.HomeDir, "pjx_path")))
		}
		SetPjxEnv()
	}
}
func main() {
	Cmd()
}

func Cmd() {
	var e error
	pj.AppPath, e = os.Getwd()
	if e != nil {
		panic(e)
	}
	pj.IfLog = IfLog(os.Args)
	pj.IfForce = IfForce(os.Args)

	args, kv := rmAttach(os.Args[1:])
	pj.KV = kv
	pj.O = kv["o"]

	switch args[0] {
	// doc
	case "version", "--version":
		fmt.Println(pj.Version())
	case "help", "--help":
		fmt.Println(pj.Usage())

		// project generate
	case "new":
		appName := args[1]
		newApp(appName)
	case "module":
		moduleName := args[1]
		newModule(moduleName)

		// package storage and migration
	case "add":
		func() {
			// pjx add db fwhezfwhez master
			// pjx add db [global master]
			directoryName := args[1]
			var namespace string
			var tag string
			switch len(args) {
			// pjx add db fwhezfwhez master
			case 4:
				namespace = args[2]
				tag = args[3]
			case 2:
				namespace = "global"
				tag = "master"
			default:
				fmt.Println(fmt.Sprintf("bad args num,want 'pjx add <pkgName> [namespace] [tagName]' but got '%s'", strings.Join(os.Args, " ")))
				return
			}
			addPkg(directoryName, namespace, tag)
		}()
	case "use":
		func() {
			// pjx use db fwhezfwhez master
			// pjx use db [global master]
			directoryName := args[1]
			var namespace string
			var tag string
			args, _ = rmAttach(args)

			switch len(args) {
			// pjx add db fwhezfwhez master
			case 4:
				namespace = args[2]
				tag = args[3]
			case 2:
				namespace = "global"
				tag = "master"
			default:
				logger.Println(fmt.Sprintf("bad args num,want 'pjx use <pkgName> [namespace] [tagName]' but got '%s'", strings.Join(os.Args, " ")))
				return
			}
			usePkg(directoryName, namespace, tag)
		}()
	case "merge":
		// pjx merge $path fwhezfwhez -f, 合并path下所有package进fwhezfwhez空间里，存在命名冲突时，覆盖
		// pjx merge $path global -u, 合并path下所有package进global里，存在命名冲突时，跳过该包

		src := args[1]
		namespace := args[2]
		mergePackage(src, namespace)

	case "clone":
		// pjx clone github.com/fwhezfwhez/pjx-packages.git fwhefwhez -f  从远程仓库中拉取并合并进fwhezfwhez空间里，命名冲突时，覆盖
		// pjx clone github.com/fwhezfwhez/pjx-packages.git fwhefwhez -u  从远程仓库中拉取并合并进fwhezfwhez空间里，存在命名冲突时，跳过该包
		//src := args[1]
		//namespace := args[2]
		// cloneFrom(src, namespace)
	default:
		fmt.Println(fmt.Sprintf("command '%s' not found", args[0]))
	}
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
	var tmpl struct {
		Package string   `yaml:"package"`
		DirList []string `yaml:"dirList"`
	}
	var tmplKey = "default"
	var m = pj.KV["m"]
	if m != "" {
		if moduleTmplMap[m] != "" {
			tmplKey = m
		}
	}
	e := yaml.Unmarshal([]byte(moduleTmplMap[tmplKey]), &tmpl)
	if e != nil {
		panic(e)
	}

	_, e = os.Stat(PathJoin(pj.AppPath, tmpl.Package))
	if e != nil {
		if os.IsNotExist(e) {
			if e := os.Mkdir(PathJoin(pj.AppPath, tmpl.Package), os.ModePerm); e != nil {
				fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
				return
			}
		} else {
			panic(e)
		}
	}

	if e := os.Mkdir(PathJoin(pj.AppPath, tmpl.Package, moduleName), os.ModePerm); e != nil {
		fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
		return
	}

	for _, v := range tmpl.DirList {
		if e := os.Mkdir(PathJoin(pj.AppPath, tmpl.Package, moduleName, strings.Replace(v, "{$object}", moduleName, -1)), os.ModePerm); e != nil {
			fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
			return
		}
	}
	//if e := os.Mkdir(PathJoin(pj.AppPath, tmpl.Package, moduleName, fmt.Sprintf("%s%s", moduleName, "Pb")), os.ModePerm); e != nil {
	//	fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
	//	return
	//}
	//if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Model")), os.ModePerm); e != nil {
	//	fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
	//	return
	//}
	//if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Router")), os.ModePerm); e != nil {
	//	fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
	//	return
	//}
	//if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Service")), os.ModePerm); e != nil {
	//	fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
	//	return
	//}
	//
	//if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "TestClient")), os.ModePerm); e != nil {
	//	fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
	//	return
	//}
	//
	//if e := os.Mkdir(path.Join(pj.AppPath, "module", moduleName, fmt.Sprintf("%s%s", moduleName, "Export")), os.ModePerm); e != nil {
	//	fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
	//	return
	//}
}

func addPkg(directoryName string, namespace string, tag string) {
	execDir, e := os.Getwd()
	if e != nil {
		panic(e)
	}

	// pkg dir should exist and is not a file
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

	// pjx_path should be a dir and well set in os env and the dir path exist
	pjxPath := os.Getenv("pjx_path")
	if pjxPath == "" {
		logger.Println(fmt.Sprintf("pjxPath not found in path, make sure it's well in system env."))
		return
	}
	fileInfo, e = os.Stat(pjxPath)
	if e != nil {
		if os.IsNotExist(e) {
			logger.Println(fmt.Sprintf("pjx_path dir not exists, '%s' not exist", pjxPath))
			return
		}
		panic(e)
	}

	if !fileInfo.IsDir() {
		logger.Println(fmt.Sprintf("pjx_path is not a  directory, '%s' is not a directory", pjxPath))
		return
	}

	// check lib exist or not
	var libPath string
	if tag == "master" {
		libPath = path.Join(pjxPath, namespace, directoryName)
	} else {
		libPath = path.Join(pjxPath, namespace, directoryName+"@"+tag)
	}
	fileInfo, e = os.Stat(FormatPath(libPath))
	if e == nil {
		// exist
		// if `--force` or `-f`, delete old existing folder.
		if IfForce(os.Args) {
			DelDir(libPath)
		} else {
			logger.Println(fmt.Sprintf("lib folder '%s' already exist no need to add. Or you can add '-f' to forcely add one.The old one will be replaced", FormatPath(libPath)))
			return
		}
	} else {
		if os.IsNotExist(e) {
			// do nothing when not exist
		} else {
			// panic if meet unexpected error
			panic(e)
		}
	}
	// prepare namespace
	fileInfo, e = os.Stat(FormatPath(PathJoin(pjxPath, namespace)))
	if e != nil {
		if os.IsNotExist(e) {
			// create when not exist
			if e := os.Mkdir(PathJoin(pjxPath, namespace), os.ModePerm); e != nil {
				fmt.Println(fmt.Sprintf("%v\n%s", e, debug.Stack()))
				return
			}
		}
	}

	// when tag is master, folder name is itself, or folder name will suffixed by '@tag'
	CopyDir(dirPath, FormatPath(libPath))

	return
}

func usePkg(directoryName string, namespace string, tag string) {
	var addedDirectoryName string
	if pj.O != "" {
		addedDirectoryName = pj.O
	} else {
		addedDirectoryName = directoryName
	}

	execDir, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	dirPath := FormatPath(path.Join(execDir, addedDirectoryName))
	fileInfo, e := os.Stat(dirPath)

	if e == nil {
		logger.Println(fmt.Sprintf("pkg exists, '%s' already exist", dirPath))
		return
	}

	pjxPath := os.Getenv("pjx_path")
	if pjxPath == "" {
		logger.Println(fmt.Sprintf("pkgPath not found in path, make sure it's well in system env."))
		return
	}

	fileInfo, e = os.Stat(pjxPath)
	if e != nil {
		if os.IsNotExist(e) {
			logger.Println(fmt.Sprintf("pjx_path dir not exist, '%s' not exist", pjxPath))
			return
		}
		panic(e)
	}

	if !fileInfo.IsDir() {
		logger.Println(fmt.Sprintf("pjx_path is not a  directory, '%s' is not a directory", pjxPath))
		return
	}

	// check lib has this folder or not
	var libPath string
	if tag == "master" {
		libPath = path.Join(pjxPath, namespace, directoryName)
	} else {
		libPath = path.Join(pjxPath, namespace, directoryName+"@"+tag)
	}
	fileInfo, e = os.Stat(FormatPath(libPath))

	if e != nil {
		if os.IsNotExist(e) {
			logger.Println(fmt.Sprintf("pjx_path not found folder '%s'", libPath))
			return
		}
		panic(e)
	}

	CopyDir(libPath, dirPath)
	return
}

func initEnv() {
	// pjx_path should be a dir and well set in os env and the dir path exist
	pjxPath := os.Getenv("pjx_path")
	if pjxPath == "" {
		SetPjxEnv()
		logger.Println(fmt.Sprintf("pjxPath not found in path, the default 'pjx_path' has been set by default, you might need to reopen command windows and type pjx env to see its detail."))
		return
	}
	fileInfo, e := os.Stat(pjxPath)
	if e != nil {
		if os.IsNotExist(e) {
			logger.Println(fmt.Sprintf("pjx_path dir not exists, '%s' not exist", pjxPath))
			return
		}
		panic(e)
	}

	if !fileInfo.IsDir() {
		logger.Println(fmt.Sprintf("pjx_path is not a  directory, '%s' is not a directory", pjxPath))
		return
	}
}

func mergePackage(src string, namespace string) {
	if src == "." {
		var e error
		src, e = os.Getwd()
		if e != nil {
			panic(e)
		}
	}

	src = FormatPath(src)
	var state = 3 // 1. -f , 2. -u , 3. stop when exist error, privilege  -u > -f > null
	if IfForce(os.Args) {
		state = 1
	}
	if IfU(os.Args) {
		state = 2
	}
	srcInfo, e := os.Stat(src)
	if e != nil {
		if os.IsNotExist(e) {
			logger.Println(fmt.Sprintf("'%s' not found", src))
			return
		}
		panic(e)
	}
	if !srcInfo.IsDir() {
		logger.Println(fmt.Sprintf("'%s' is not dir", src))
		return
	}

	rd, e := ioutil.ReadDir(src)
	if e != nil {
		panic(e)
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			continue
		}
		src := PathJoin(src, fi.Name())
		dest := PathJoin(os.Getenv("pjx_path"), namespace, fi.Name())

		namespacePath := PathJoin(os.Getenv("pjx_path"), namespace)
		_,e:=os.Stat(namespace)
		if e!=nil {
			if os.IsNotExist(e) {
				if e:= os.Mkdir(namespacePath, os.ModePerm);e!=nil {
					panic(e)
				}
			}else{
				panic(e)
			}
		}

		switch state {
		case 1:
			CopyDirF(src, dest)
		case 2:
			CopyDirU(src, dest)
		case 3:
			CopyDir(src, dest)
		}
	}
	return
}
