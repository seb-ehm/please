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
	f, err := os.OpenFile("pleaseoutput.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if *versionPtr {
		fmt.Println("Version 0.0.1")
	} else {
		f.WriteString("Input: \n")
		if isNoPipe(os.Stdin) {
			f.WriteString("No Pipe \n")
			args := flag.Args()
			fmt.Println(strings.Join(args, " "))
			return
		}

		reader := bufio.NewReader(os.Stdin)
		var output []string
		f.WriteString("Pipe: \n")
		for {
			input, _, err := reader.ReadLine()
			if err != nil && err == io.EOF {
				break
			}
			f.WriteString(string(input))
			output = append(output, string(input))
		}
		f.WriteString("Output: \n")
		for j := 0; j < len(output); j++ {
			f.WriteString(output[j])
			fmt.Printf("%s\n", output[j])
		}
		f.WriteString("\n")
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
