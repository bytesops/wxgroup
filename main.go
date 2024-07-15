package main

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
	localMem "wxgroup/controllers"
	_ "wxgroup/routers"
)

func getNextDir(path string) {
	// 读取文件夹内容
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 遍历文件夹
	for _, entry := range entries {
		// 获取完整的路径
		fullPath := filepath.Join(path, entry.Name())

		// 判断是文件还是文件夹
		if entry.IsDir() {
			fmt.Printf("Dir: %s\n", fullPath)
			getNextDir(fullPath)
		} else {
			localMem.LocalQrcodes[entry.Name()] = fullPath
		}
	}
}

func initLocalQrcodes() {
	rootPath := "./qrcodes"
	// 读取文件夹内容
	entries, err := ioutil.ReadDir(rootPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 遍历文件夹
	for _, entry := range entries {
		// 获取完整的路径
		fullPath := filepath.Join(rootPath, entry.Name())

		// 判断是文件还是文件夹
		if entry.IsDir() {
			fmt.Printf("Dir: %s\n", fullPath)
			getNextDir(fullPath)
		} else {
			localMem.LocalQrcodes[entry.Name()] = fullPath
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initLocalQrcodes()
	var list []string
	for k, _ := range localMem.LocalQrcodes {
		fmt.Println(k)
		list = append(list, k)
	}
	localMem.LocalQrcodeStrs = list
	beego.Run()
}
