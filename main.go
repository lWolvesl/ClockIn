package main

import (
	"os"
	"time"
)

func getTimeStr() string {
	timeFormat := "2006-01-02 15:04:05.000"
	now := time.Now()
	return now.Format(timeFormat)
}

func punch(text string) {
	now := time.Now()
	filename := "./" + now.Format("2006-01-02") + ".out"

	fp, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm) // 读写方式打开

	// defer延迟调用
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {

		}
	}(fp) //关闭文件，释放资源。

	finalText := getTimeStr() + " " + text + "\n"

	_, err := fp.WriteString(finalText)

	if err != nil {
		return
	}
}

func main() {
	punch(os.Args[1])
}
