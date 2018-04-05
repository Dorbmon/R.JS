package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	sum int64
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("countline [Dir] [extention]")
		return
	}
	filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if getExt(path) == os.Args[2] {
			fmt.Println(path)
			n, e := countFileLine(path)
			if e == nil {
				fmt.Println("lines: ", n)
				sum = sum + n
			}
		}
		return nil
	})
	fmt.Println("total lines: ", sum)
}

//获取得文件的扩展名，最后一个.后面的内容
func getExt(f string) (ext string) {
	//  fmt.Println("ext:", f)
	index := strings.LastIndex(f, ".")
	data := []byte(f)
	for i := index + 1; i < len(data); i++ {
		ext = ext + string([]byte{data[i]})
	}
	return
}

func countFileLine(name string) (count int64, err error) {
	fmt.Println("count file:", name)
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	count = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	fmt.Println("line:", count)
	return
}