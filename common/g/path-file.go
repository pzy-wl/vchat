package g

import (
	"errors"
	//  "fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	//fmt.Println("path111:", path)
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	//fmt.Println("path222:", path)
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.New(`Can't find "/" or "\".`)
	}
	//fmt.Println("path333:", path)
	return string(path[0 : i+1]), nil
}
