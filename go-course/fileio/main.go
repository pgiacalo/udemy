package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage: go run main.go <read or write> <file_path> [contents]")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch args[1] {
	case "read":
		filePath := args[2]
		b, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	case "write":
		filePath := args[2]
		err := os.WriteFile(filePath, []byte(args[3]), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		printUsage()
		os.Exit(1)
	}
}
