package main

import (
	"fmt"
	"os"
)

func walkDir(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())

		if file.IsDir() {
			subDir := path + "/" + file.Name()

			if err := walkDir(subDir); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	walkDir(".")
}