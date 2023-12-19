package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var url string
	fmt.Print("Введите URL для сканирования: ")
	_, err := fmt.Scan(&url)

	if err != nil {
		fmt.Println("Ошибка при вводе URL:", err)
		return
	}

	fmt.Println("[+] Start enumerating", url)

	fileName := "typo3-12.2.0.list"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		fullURL := url + scanner.Text()
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			fmt.Println("[+] ", fullURL)
		}
		if lineCount%60 == 0 {
			fmt.Printf("Текущая строка из списка: %d\n", lineCount)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}
