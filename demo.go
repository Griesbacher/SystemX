package main

import (
	"github.com/griesbacher/SystemX/bin"
	"time"
)

func main() {
	go bin.Server("serverConfig.gcfg", "clientConfig.gcfg", "")
	time.Sleep(time.Duration(1) * time.Second)
	bin.Client("clientConfig.gcfg", "")
}
