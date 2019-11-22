package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	versionPtr := flag.Bool("v", false, "Print the version")
	flag.Parse()

	if *versionPtr {
		fmt.Println("Version 0.0.1")
	} else {
		if isNoPipe(os.Stdin) {
			args := flag.Args()
			fmt.Println(strings.Join(args, " "))
			return
		}

		reader := bufio.NewReader(os.Stdin)
		var commandHistory []string
		for {
			input, _, err := reader.ReadLine()
			if err != nil && err == io.EOF {
				break
			}
			commandHistory = append(commandHistory, string(input))
		}

	}
}

func isNoPipe(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	mode := info.Mode()
	isPipe := mode&os.ModeNamedPipe != 0
	isCharDevice := mode&os.ModeDevice != 0 && mode&os.ModeCharDevice != 0

	return !isPipe || isCharDevice
}
