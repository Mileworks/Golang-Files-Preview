package main

import (
	"github.com/astaxie/beego"
	_ "github.com/mileworks/plm-files-preview/routers"
)

// /Users/mac/go/bin/bee pack -be GOOS=linux -be GOARCH=amd64
func main() {
	beego.Run()
}
