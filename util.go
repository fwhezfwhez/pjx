package main

import (
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
)

// format path, different between windows and linux/mac
func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
		//return strings.Replace(s, , "/", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		logger.Println("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}

// copy folder from src to dest
func CopyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)

	if IfLog(os.Args) {
		logger.Println("src:", src)
		logger.Println("dest:", dest)
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
		if pj.IfLog {
			logger.Println("exec `xcopy " + src + " " + dest + " " + "/I" + " " + "/E`")

		}
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
		if pj.IfLog {
			logger.Println("exec `cp -R " + src + " " + dest + "`")
		}
	}

	outPut, e := cmd.Output()
	if e != nil {
		logger.Println(e.Error())
		return
	}
	if pj.IfLog {
		logger.Println(string(outPut))
	}
}

func DelDir(src string) {
	src = FormatPath(src)
	switch runtime.GOOS {
	case "windows":
		os.RemoveAll(src)
		//cmdStr := fmt.Sprintf(`rd /s /q %s` , SingleRightDownPath(src))
		//logger.Println(SingleRightDownPath(src))
		//cmdList := strings.Split(cmdStr, " ")
		//cmd = exec.Command(cmdList[0], cmdList[1:]...)
		//fmt.Println("exec `rd" + " /s " + "/q " + src + "`")
	case "darwin", "linux":
		//cmd = exec.Command("rm", "-rf", src)
		////cmd = exec.Command("sudo rm", "-rf", src)
		//fmt.Println("sudo exec `rm -rf " + src)
		//outPut, e := cmd.Output()
		//if e != nil {
		//	logger.Println(e.Error())
		//	return
		//}
		//logger.Println(string(outPut))
		os.RemoveAll(src)
	}

}

// check os.Args contains -f or not
func IfForce(Arg []string) bool {
	for _, v := range Arg {
		if v == "-f" || v == "-F" || v == "--force" || v == "--Force" {
			return true
		}
	}
	return false
}

func IfLog(Arg []string) bool {
	for _, v := range Arg {
		if v == "-l" || v == "-L" || v == "--log" || v == "--Log" {
			return true
		}
	}
	return false
}

func PathJoin(args ...string) string {
	var tmp = make([]string, 0, 10)
	for _, v := range args {
		if v == "nil" || v == "empty" || v == "" {
			continue
		}
		tmp = append(tmp, v)
	}
	return FormatPath(path.Join(tmp...))
}

// rm arg with - or -- prefix
// when meet specific value lime -o, it will save key'o' and its value into kv map as part of return.
func rmAttach(arr []string) ([]string, map[string]string) {
	var newArr = make([]string, 0, 10)
	var kv = make(map[string]string, 0)
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		if strings.HasPrefix(v, "-") || strings.HasPrefix(v, "--") {
			if v == "-o" {
				i += 1
				kv["o"] = arr[i]
				continue
			}
			if v == "-m" {
				i += 1
				kv["m"] = arr[i]
				continue
			}
			continue
		} else {
			newArr = append(newArr, v)
		}
	}
	return newArr, kv
}

// This is temporary value, as soon as process stops, env loses its effect.
func SetPjxEnv() {
	if os.Getenv("pjx_path") != "" {
		logger.Println("pjx_path already exist, no need to set again")
		return
	}

	var key = "pjx_path"
	var value string
	usInfo, e := user.Current()
	if e != nil {
		panic(e)
	}
	value = PathJoin(usInfo.HomeDir, "pjx_path")
	_, e = os.Stat(value)
	if e != nil {
		if os.IsNotExist(e) {
			if e := os.Mkdir(value, os.ModePerm); e != nil {
				logger.Println(e.Error())
				return
			}
		} else {
			panic(e)
		}
	}

	if e := os.Setenv(key, value); e != nil {
		panic(e)
	}
}
