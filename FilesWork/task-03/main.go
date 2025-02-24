// 3. Подсчет строк, слов и символов

// Задача: Напишите программу, которая принимает
// имя текстового файла в качестве аргумента и выводит
// количество строк, слов и символов в этом файле.

// Задачи:
// Используйте bufio.Scanner для построчного чтения файла.
// Используйте пакет strings для разделения строк на слова.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func counter(name string) (int, int, int) {
    file, err := os.Open(name)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return 0, 0, 0
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)

    var words int
    var lines int
    var symbols int

    for scanner.Scan() {
        line := scanner.Text()
        lines++
        words += len(strings.Fields(line))
        
        // Проходимся по строке, чтобы считать символы
        // а не байты
        for range line {
            symbols++
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return 0, 0, 0
    }

    return lines, words, symbols
}


func main() {
    lines, words, symbols := counter("input.txt")
    fmt.Println("Lines:", lines)
    fmt.Println("Words:", words)
    fmt.Println("Symbols:", symbols)
}