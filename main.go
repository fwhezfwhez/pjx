package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/howeyc/gopass"
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
	//fmt.Println(args)
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
		// pjx merge path fwhefwhez -f
		// pjx merge path global -u
		src := args[1]
		namespace := args[2]
		mergePackage(src, namespace)
	case "clone":
		src := args[1]
		namespace := args[2]
		cloneFrom(src, namespace)
	case "test":
		fmt.Println("test module are developing")
	case "exec":
		// pjx exec print-hello.pjx
		// pjx exec print-hello
		src := args[1]
		if !strings.HasSuffix(src, ".pjx") {
			src = src + ".pjx"
		}
		execFileCommand(src)
	case "encrypt":
		secret := pj.KV["secret"]
		if secret == "" {
			fmt.Print("Please input secret key:")
			pw, e := gopass.GetPasswdMasked()
			if e != nil {
				panic(e)
			}
			secret = string(pw)
		}

		file := pj.KV["file"]

		if file == "" {
			file = strings.Join(args[1:], ",")
		}
		handleEncrypt(file, secret)
	case "decrypt":
		secret := pj.KV["secret"]
		if secret == "" {
			fmt.Print("Please input secret key:")
			pw, e := gopass.GetPasswdMasked()
			if e != nil {
				panic(e)
			}
			secret = string(pw)
		}

		file := pj.KV["file"]

		if file == "" {
			file = strings.Join(args[1:], ",")
		}
		handleDecrypt(file, secret)

	default:
		fmt.Println(fmt.Sprintf("command '%s' not found", args[0]))
	}
}

func handleEncrypt(fileNames string, secret string) {
	var IfReg = IfReg(os.Args)
	var dir = pj.KV["d"]
	if dir == "" {
		var e error
		dir, e = os.Getwd()
		if e != nil {
			panic(e)
		}
	}
	dir = FormatPath(dir)

	fileNameArr := strings.Split(fileNames, ",")
	for _, v := range fileNameArr {
		v = strings.TrimSpace(v)
		if !IfReg {
			encryptConfigFile(v, secret)
		} else {
			//path.Base()
			if strings.Contains(v, "/") || strings.Contains(v, "\\") {
				panic("regex can't decode filenames containing \\ or /")
			}
			// path.Match(v, )
			files, err := ioutil.ReadDir(dir) //读取目录下文件
			if err != nil {
				return
			}

			for _, file := range files {
				if file.IsDir() {
					continue
				}
				matched, e := path.Match(v, file.Name())
				if e != nil {
					panic(e)
				}

				if matched {
					encryptConfigFile(path.Join(dir, file.Name()), secret)
				}
			}
		}
	}
}

func handleDecrypt(fileNames string, secret string) {
	var IfReg = IfReg(os.Args)
	var dir = pj.KV["d"]
	if dir == "" {
		var e error
		dir, e = os.Getwd()
		if e != nil {
			panic(e)
		}
	}
	dir = FormatPath(dir)

	fileNameArr := strings.Split(fileNames, ",")
	for _, v := range fileNameArr {
		v = strings.TrimSpace(v)
		if !IfReg {
			decryptConfigFile(v, secret)
		} else {
			//path.Base()
			if strings.Contains(v, "/") || strings.Contains(v, "\\") {
				panic("regex can't decode filenames containing \\ or /")
			}
			// path.Match(v, )
			files, err := ioutil.ReadDir(dir) //读取目录下文件
			if err != nil {
				return
			}

			for _, file := range files {
				if !strings.HasSuffix(file.Name(), ".crt") {
					continue
				}
				if file.IsDir() {
					continue
				}
				matched, e := path.Match(v, file.Name())
				if e != nil {
					panic(e)
				}

				if matched {
					decryptConfigFile(path.Join(dir, file.Name()), secret)
				}
			}
		}
	}
}

func encryptConfigFile(fileName string, secret string) {
	path := FormatPath(fileName)
	_, e := os.Stat(path)
	if e == os.ErrNotExist {
		panic(fmt.Errorf("file '%s' not found", path))
	}
	f, e := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if e != nil {
		panic(fmt.Errorf("open '%s' err '%v'", fileName, e))
	}
	defer f.Close()

	b, e := ioutil.ReadAll(f)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))

	newFilePath := path + ".crt"

	f2, e := os.OpenFile(newFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer f2.Close()

	crtContent, e := Encrypt(b, secret)
	if e != nil {
		panic(e)
	}
	f2.Write([]byte(crtContent))
}

func decryptConfigFile(fileName string, secret string) {
	path := FormatPath(fileName)
	_, e := os.Stat(path)
	if e == os.ErrNotExist {
		panic(fmt.Errorf("file '%s' not found", path))
	}

	if !strings.HasSuffix(fileName, ".crt") {
		panic(fmt.Errorf("only decrypt file ended with '.crt' but got '%s'", fileName))
	}

	f, e := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	b, e := ioutil.ReadAll(f)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))

	newFilePath := strings.TrimSuffix(path, ".crt")

	f2, e := os.OpenFile(newFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer f2.Close()

	crtContent, e := Decrypt(string(b), secret)
	if e != nil {
		panic(e)
	}
	f2.Write([]byte(crtContent))
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
	var p = pj.KV["p"]
	if p != "" {
		tmpl.Package = p
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

	if f, e := os.OpenFile(path.Join(pj.AppPath, tmpl.Package, moduleName, "main.go"), os.O_CREATE|os.O_RDWR, os.ModePerm); e != nil {
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
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		if !fi.IsDir() {
			continue
		}
		src := PathJoin(src, fi.Name())
		dest := PathJoin(os.Getenv("pjx_path"), namespace, fi.Name())

		namespacePath := PathJoin(os.Getenv("pjx_path"), namespace)
		_, e := os.Stat(namespacePath)
		if e != nil {
			if os.IsNotExist(e) {
				if e := os.Mkdir(namespacePath, os.ModePerm); e != nil {
					panic(e)
				}
			} else {
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

func cloneFrom(src, namespace string) {
	pjName := GetGitName(src)
	cmdList := []string{"clone", src, pjName}
	cmd := exec.Command("git", cmdList...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if e := cmd.Run(); e != nil {
		panic(e)
	}
	if IfLog(os.Args) {
		logger.Println(out.String())
	}
	out.Reset()
	currentDir, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	var f string
	var u string
	if IfForce(os.Args) {
		f = "-f"
	}
	if IfU(os.Args) {
		u = "-u"
	}
	cmdList = []string{"merge", PathJoin(currentDir, pjName), namespace, f, u}
	cmd = exec.Command("pjx", cmdList...)
	cmd.Stdout = &out
	cmd.Run()
	if IfLog(os.Args) {
		logger.Println(out.String())
	}
	out.Reset()
	DelDir(PathJoin(currentDir, pjName))
}

func execFileCommand(src string) {
}
