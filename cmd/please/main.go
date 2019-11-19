package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	versionPtr := flag.String("v", "no", "Request the version")
	flag.Parse()
	if *versionPtr != "no" {
		fmt.Println("Version 0.0.1")
	} else {
		args := flag.Args()
		fmt.Println(strings.Join(args, "\\ "))
	}

}
