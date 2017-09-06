package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Read(path string) string {
	return string(ReadByte(path))
}

func ReadByte(path string) []byte {
	fi, err := os.Open(path)
	if err != nil {
		println("Files Logs Error:" + err.Error())
		panic(err)
	}
	defer fi.Close()
	buf, err := ioutil.ReadAll(fi)
	if err != nil {
		println("Files Logs Error:" + err.Error())
		panic(err)
	}
	return buf
}

//系统分隔符
func Separator() string {
	var path string
	//前边的判断是否是系统的分隔符
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	return path
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

//当前目录
func CurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func JsonPath(params ...string) string {
	cp := CurrentDirectory()
	for _, p := range params {
		cp += Separator() + p
	}
	return cp
}

func JsonParentPath(params ...string) string {
	cp := getParentDirectory(CurrentDirectory())
	for _, p := range params {
		cp += Separator() + p
	}
	return cp
}
