package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	versionPtr := flag.Bool("v", false, "Print the version")
	flag.Parse()
	if *versionPtr {
		fmt.Println("Version 0.0.1")
	} else {
		args := flag.Args()
		fmt.Println(strings.Join(args, " "))
	}

}
