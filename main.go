package main

import (
	"flag"
	"log"
	"relay-chain/INIT"
)

var (
	ConfPath = ""
	UnLog    = false
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//从命令行读取参数（例如：-conf=...）
	flag.StringVar(&ConfPath, "conf", ConfPath, "config path")
	flag.BoolVar(&UnLog, "unLog", UnLog, "disable logs(default: false)")
	flag.Parse()

	INIT.Init(ConfPath)
	INIT.Start(UnLog)
}
