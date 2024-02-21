package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if check(os.Args[1]) {
		fmt.Println("File not found")
	} else {
		read(os.Args[1])
	}
}

func read(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	res := fmt.Sprintf("%s", data)
	fmt.Print(res)
}

func check(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return true
	}
	return false
}
