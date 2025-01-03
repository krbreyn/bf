package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/krbreyn/bf_go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: bfgo <input_file.bf>")
		return
	}

	programPath := os.Args[1]

	switch filepath.Ext(programPath) {
	case ".b", ".bf":
		// nothing
	default:
		fmt.Println("file did not end in .b or .bf")
		return
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error reading file", err)
	}

	fmt.Println(bf_go.RunProgram(string(file)))

}
