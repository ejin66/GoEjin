package util

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

func InArray(ary []string, v string) bool {
	for _, item := range ary {
		if item == v {
			return true
		}
	}
	return false
}

func Substring(s string, from int, length int) string {
	if len(s) < length+from {
		length = len(s) - from
	}
	return s[from : length+from]
}

func Replace(s string, new string, regex string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(s, new)
}

func PrintError(msg interface{}) {
	//其中0x1B是标记，[开始定义颜色，1代表高亮，39代表黑色背景，31代表红色前景，0代表恢复默认颜色。
	fmt.Printf("%c[0;39;31m%s%s%c[0m\n", 0x1B, "Error: ", msg, 0x1B)
}

func Error404() string {
	return "404 no found"
}

func Error404Web() string {
	return "404 no found"
}

func Web(path string) string {
	//TODO path的位置
	fmt.Print("src/web/" + path)
	file, err := os.OpenFile("src/web/"+path, os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("open web page err:", err.Error())
		return Error404Web()
	}
	buf := bytes.NewBufferString("")
	for {
		var buffer = make([]byte, 1024*10)
		i, err := file.Read(buffer)
		if i <= 0 {
			break
		}
		if err != nil {
			fmt.Println("read web page err:", err.Error(), i)
			break
		}
		buf.Write(buffer[:i])
	}
	return buf.String()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
