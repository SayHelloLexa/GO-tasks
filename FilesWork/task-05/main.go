package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ln, _, _ := reader.ReadLine()

	fmt.Println(strings.ToUpper(string(ln)))
}