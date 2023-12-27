package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func insertToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var msg string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		msg += string(line) + "\n"
	}

	return msg, nil
}

func main() {
	// createNewFile("sample.log", "this is sample log")
	// result, _ := readFile("sample.log")
	// fmt.Println(result)
	insertToFile("sample.log", "\nthis is add message")
}
