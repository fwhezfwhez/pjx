package main

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
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
	_, e := os.Stat(dest)
	if e == nil {
		logger.Println(fmt.Sprintf("dest '%s' exist, do nothing", dest))
		return
	}
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

// same as copyDir, but will forcely replace the old if dest existed.
func CopyDirF(src string, dest string) {
	_, e := os.Stat(dest)
	if e != nil {
		if os.IsNotExist(e) {
			CopyDir(src, dest)
			return
		}
		panic(e)
	}
	DelDir(dest)
	CopyDir(src, dest)
}

// same as copyDir, but will do nothing if dest exists
func CopyDirU(src string, dest string) {
	info, e := os.Stat(dest)
	if e != nil {
		if os.IsNotExist(e) {
			CopyDir(src, dest)
			return
		}
		panic(e)
	}

	if !info.IsDir() {
		logger.Println(fmt.Sprintf("dest '%s' is not a dir", dest))
		return
	}
	if IfLog(os.Args) {
		logger.Println(fmt.Sprintf("dest '%s' exist, do nothing", dest))
	}
	return
}

func DelDir(src string) {
	src = FormatPath(src)

	_, e := os.Stat(src)
	if e != nil {
		if os.IsNotExist(e) {
			if IfLog(os.Args) {
				logger.Println(fmt.Sprintf("del dir '%s' not found, do nothing", src))
			}
			return
		}
		panic(e)
	}
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

// -l
func IfLog(Arg []string) bool {
	for _, v := range Arg {
		if v == "-l" || v == "-L" || v == "--log" || v == "--Log" {
			return true
		}
	}
	return false
}

// -u
func IfU(Arg []string) bool {
	for _, v := range Arg {
		if v == "-u" || v == "-U" {
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
		if v == "--help" || v=="--version"{
			newArr = append(newArr, v)
		}
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
			if v == "-file" {
				i += 1
				kv["file"] = arr[i]
				continue
			}
			if v == "-secret" {
				i += 1
				kv["secret"] = arr[i]
				continue
			}

			if v == "-d" {
				i += 1
				kv["d"] = arr[i]
				continue
			}

			if v == "-p" {
				i += 1
				kv["p"] = arr[i]
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

// get srv from https://github.com/xxx/srv.git
func GetGitName(src string) string {
	arr := strings.Split(src, "/")

	return strings.Split(arr[(len(arr) - 1)], ".")[0]
}

// Split 增强型Split，对  a,,,,,,,b,,c     以","进行切割成[a,b,c]
func Split(s string, sub string) []string {
	var rs = make([]string, 0, 20)
	tmp := ""
	Split2(s, sub, &tmp, &rs)
	return rs
}

// Split2 附属于Split，可独立使用
func Split2(s string, sub string, tmp *string, rs *[]string) {
	s = strings.Trim(s, sub)
	if !strings.Contains(s, sub) {
		*tmp = s
		*rs = append(*rs, *tmp)
		return
	}
	for i := range s {
		if string(s[i]) == sub {
			*tmp = s[:i]
			*rs = append(*rs, *tmp)
			s = s[i+1:]
			Split2(s, sub, tmp, rs)
			return
		}
	}
}

func Encrypt(content []byte, key string) (string, error) {
	if key == "" {
		panic("secret emtpy")
	}
	key = MD5(key)[:8]
	src := []byte(content)
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func Decrypt(decrypted string, keyStr string) ([]byte, error) {
	keyStr = MD5(keyStr)[0:8]
	key := []byte(keyStr)
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return nil, err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}

	out = ZeroUnPadding(out)
	return out, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

// md5 加密
func MD5(rawMsg string) string {
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str1)
}

// get dir of a file
func DirOf(fileName string) string {
	return path.Dir(fileName)
}

func IfReg(Arg []string) bool {
	return true
	for _, v := range Arg {
		if v == "-r" || v == "-R" || v == "-P" || v == "-p" || v == "reg" || v == "regex" {
			return true
		}
	}
	return false
}

func IfPartern(Arg []string) bool {
	return IfReg(Arg)
}
