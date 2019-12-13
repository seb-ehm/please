package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"please/pkg/command"
	"please/pkg/util"
	"time"
)

func main() {
	versionPtr := flag.Bool("v", false, "Print the version")
	flag.Parse()

	if *versionPtr {
		fmt.Println("Version 0.0.1")
	} else {
		if util.IsNoPipe(os.Stdin) {
			flag.Usage()
			return
		}
		input := make(chan string)
		go func() { //util.IsNoPipe currently cannot correctly identify non-pipe input from Git-for-Windows, so a timeout is necessary
			reader := bufio.NewReader(os.Stdin)

			for {
				data, _, err := reader.ReadLine()
				if err != nil && err == io.EOF {
					break
				} else {
					input <- string(data)
				}
			}
			close(input)
		}()
		var commandHistory []string
		timeout, _ := time.ParseDuration("1s")
		select {
		case line := <-input:
			commandHistory = append(commandHistory, line)
		case <-time.After(timeout):
			flag.Usage()
		}
		if len(commandHistory) > 0 {
			suggestion := command.Suggest(command.New(commandHistory[0]))
			fmt.Println(suggestion)
		}

	}
}
