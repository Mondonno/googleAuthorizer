package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings",
	"rec"
)

func start(key string) {
	rec.StartLooping(key)
}

func main() {
	key := flag.String("key", "", "Gauth key used to generate new auth 6 digit key")
	flag.Parse();

	if(key == "") {
		flag.PrintDeafults()
		return
	}

	go start(key)
}