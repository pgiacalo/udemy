package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printUsage() {
	fmt.Println("Usage: go run main.go <read> <file_path>")
	fmt.Println("Usage: go run main.go <write> <file_path> [contents] [append bool]")
}

func writeToFile(filePath string, contents string, append bool, perm os.FileMode) error {
	if !append {
		return ioutil.WriteFile(filePath, []byte(contents+"\n"), perm)
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(contents + "\n"); err != nil {
		return err
	}

	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch args[1] {
	case "read":
		if len(args) < 3 {
			printUsage()
			os.Exit(1)
		}
		filePath := args[2]
		b, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error during read attempt:", err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	case "write":
		if len(args) < 4 {
			printUsage()
			os.Exit(1)
		}
		append := true
		if len(args) == 5 && args[4] == "false" {
			append = false
		}
		file := args[2]
		err := writeToFile(file, args[3], append, 0644)
		if err != nil {
			fmt.Println("Error during write attempt:", err)
			os.Exit(1)
		}
	default:
		printUsage()
		os.Exit(1)
	}
}
