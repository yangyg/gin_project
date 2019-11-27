package main

import (
	"flag"
	_ "gin_project/config"
	"gin_project/routers"
)

var (
	h    bool
	port string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&port, "p", ":9999", "port")
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	router := routers.SetupRouter()
	router.Run(port)
}
