package main

import (
	"flag"
	"fmt"
	// "log"
	// "net/http"
	"github.com/scjtqs/go-wxbot/utils"
)

const (
	//Version 当前版本
	Version = "0.1.1"
)

var debug = flag.String("d", "on", "if on debug mode")

func main() {

	flag.Parse()

	fmt.Printf("debug mode %s\n", *debug)

	wx := utils.Wxweb{}
	wx.start()
}
