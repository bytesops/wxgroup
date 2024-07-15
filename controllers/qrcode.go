package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"os"
)

type QrcodeController struct {
	beego.Controller
}

func (c *QrcodeController) Get() {
	c.Ctx.Output.Context.ResponseWriter.Header().Set("Content-Type", "image/jpeg")
	key := c.GetString("img")
	path := LocalQrcodes[key]
	fmt.Println(path)
	// 读取文件内容
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
	c.Ctx.Output.ContentType("image/jpeg")
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = c.Ctx.Output.Body(all)
	if err != nil {
		return
	}
}
