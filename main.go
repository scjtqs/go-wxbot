package main

import (
	"flag"
	"fmt"
	"github.com/scjtqs/go-wxbot/utils"
	// "log"
	// "net/http"
)

const (
	//Version 当前版本
	Version = "0.1.1"
)

func main() {

	flag.Parse()

	fmt.Printf("debug mode %s\n", *utils.Debug)

	wx := utils.Wxweb{}
	wx.Start()
}
