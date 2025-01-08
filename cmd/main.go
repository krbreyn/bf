package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/krbreyn/bf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: bfgo <input_file.bf> [compress/decompress (optional)]")
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

	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "compress":
			fmt.Println(bf.CompressProgram(string(file)))
		case "decompress":
			fmt.Println(bf.DecompressProgram(string(file)))
		default:
			fmt.Println("invalid command")
		}
	} else {
		fmt.Println(bf.RunProgram(string(file)))
	}
}
