package main 

import {
	"flag",
	"fmt",
	"log",
	"net/http",
	"os"
}

func main(){
	PORT := flag.Int("port", 0, "3000")
	ORIGIN := flag.String("origin", "", "")
	CLEAR_CACHE := flag.Bool("clear-cache", false, "Clear the Cache")
	flag.Parse()
	proxy := proxy.NewProxy("http://cache-server.io")

}