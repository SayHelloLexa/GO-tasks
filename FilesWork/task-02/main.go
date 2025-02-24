package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Открываем входной файл
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()

	// Создаем экземпляр ридера
	reader := bufio.NewReader(input)

	// Создаем и открываем выходной файл
	output, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer output.Close()

	// Копируем содержимое из входного в выходной
	_, err = io.Copy(output, reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Читаем содержимое выходного файла
	outputData, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(outputData))
}