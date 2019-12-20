package main

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestF(t *testing.T) {
	// get running dir path
	fmt.Println(os.Getwd())
}

func TestNewApp(t *testing.T) {
	newApp("helloWorld")
}

func TestNewModule(t *testing.T) {
	newModule("user")
}

func TestAddPackage(t *testing.T) {
	addPkg("example", "fwhezfwhez", "master")
}

func TestEncrypt(t *testing.T) {
	encryptConfigFile("G:\\go_workspace\\GOPATH\\src\\pjx\\xx.json", "hello")
}
func TestDecrypt(t *testing.T) {
	decryptConfigFile("xx.json.crt", "hello")
}

func TestBase(t *testing.T) {
	fmt.Println(path.Base("/home/kk.json"))
	fmt.Println(os.Getwd())
}

func TestEncryptFiles(t *testing.T) {
	os.Args = []string{"-r"}
	pj.KV = map[string]string{"d": "G:\\go_workspace\\GOPATH\\src\\pjx\\modules\\encrypt-decrypt\\config"}
	handleEncrypt("*.json", "hello")
}
