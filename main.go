package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"time"
)

// 打卡目录
const path = "/Users/li/Downloads/punch/"

func getTimeStr() string {
	timeFormat := "2006-01-02 15:04:05.000"
	now := time.Now()
	return now.Format(timeFormat)
}

func punch(text string) {
	now := time.Now()
	filename := path + now.Format("2006-01-02") + ".out"

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

	color.Blue(finalText)
}

func punchHistory() {
	now := time.Now()
	filename := path + now.Format("2006-01-02") + ".out"

	//2、逐行读取
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, os.ModePerm) //打开
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file) //关闭

	line := bufio.NewReader(file)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		text := string(content)
		color.Blue(text)
	}
}

func main() {
	switch len(os.Args) {
	case 1:
		punchHistory()
	case 2:
		punch(os.Args[1])

	}
}
