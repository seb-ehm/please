package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"please/pkg/command"
	"please/pkg/util"
	"strings"
)

func main() {
	versionPtr := flag.Bool("v", false, "Print the version")
	flag.Parse()

	if *versionPtr {
		fmt.Println("Version 0.0.1")
	} else {
		if util.IsNoPipe(os.Stdin) {
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
			suggestion := command.Suggest(commandHistory)
			fmt.Println(suggestion)
		}

	}
}
