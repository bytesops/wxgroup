package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type GroupController struct {
	beego.Controller
}

var fileList []string

func getGroups(path string) {
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
		ok := strings.Contains(entry.Name(), ".DS_Store")
		if ok {
			continue
		}

		// 判断是文件还是文件夹
		if entry.IsDir() {
			fmt.Printf("Dir: %s\n", fullPath)
			getGroups(fullPath)
		} else {
			getMd5(fullPath)
			result := append(fileList, fullPath)
			fmt.Printf("File: %s\n", fullPath)
			fmt.Printf("append result: %s\n", result)
		}
	}
}

func getMd5(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 计算MD5值
	md5Hash := md5.Sum(data)
	md5Str := hex.EncodeToString(md5Hash[:])
	// 输出MD5值
	fmt.Println("MD5:", md5Str)

	err = ioutil.WriteFile("qrcodes/"+md5Str+".png", data, 0755)
	if err != nil {
		return
	}
}

func (c *GroupController) Get() {
	rootPath := "./groups"
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
		ok := strings.Contains(entry.Name(), ".DS_Store")
		if ok {
			continue
		}

		// 判断是文件还是文件夹
		if entry.IsDir() {
			fmt.Printf("Dir: %s\n", fullPath)
			getGroups(fullPath)
		} else {
			getMd5(fullPath)
			result := append(fileList, fullPath)
			fmt.Printf("File: %s\n", fullPath)
			fmt.Printf("append result: %s\n", result)
		}
	}

	for _, f := range fileList {
		fmt.Printf("File: %s\n", f)
	}
	body := make(map[string]string)
	body["code"] = "0"
	body["msg"] = "success"
	c.Data["json"] = &body
	err = c.ServeJSON()
	if err != nil {
		return
	}
}
