package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"math/rand"
	"os"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

var currentQrcode = ""
var currentQrcodePath = ""

func mvToScanned() {
	dest := "scanned/" + currentQrcode
	err := os.Rename(currentQrcodePath, dest)
	if err != nil {
		fmt.Println(err)
		return
	}
	delete(LocalQrcodes, currentQrcode)
	var list []string
	for k, _ := range LocalQrcodes {
		fmt.Println(k)
		list = append(list, k)
	}
	LocalQrcodeStrs = list
	fmt.Println(len(LocalQrcodes))
}

func (c *MainController) Post() {
	fmt.Println(len(LocalQrcodes))
	body := make(map[string]string)
	if len(LocalQrcodes) == 0 {
		c.Abort("404")
		return
	}
	if currentQrcode != "" {
		mvToScanned()
		currentQrcode = ""
	}
	index := rand.Intn(len(LocalQrcodes))
	body["img"] = LocalQrcodeStrs[index]
	currentQrcode = LocalQrcodeStrs[index]
	currentQrcodePath = LocalQrcodes[LocalQrcodeStrs[index]]
	c.Data["json"] = &body
	err := c.ServeJSON()
	if err != nil {
		return
	}
}
