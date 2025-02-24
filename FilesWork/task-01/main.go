package main

import (
	"fmt"
	"os"
)

func main() {
	filedata, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	rs := []rune(string(filedata))

	for i := len(rs) - 1; i >= 0; i-- {
		file.WriteString(string(rs[i]))
	}

	filedata, err = os.ReadFile("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(filedata))
}