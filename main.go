package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/andrei-maslov/ritualpay/internal/parser"
)

func getDir() string {
	file, _ := os.Executable()
	return filepath.Dir(file)
}

func findOrderFiles(dir string) ([]string, error) {
	var files []string

	error := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println("\t\t", path, "  ", info.Name())
		if err != nil {
			return err
		}
		if (strings.HasSuffix(strings.ToLower(info.Name()), ".xls") ||
			strings.HasSuffix(strings.ToLower(info.Name()), ".xlsx")) &&
			strings.Contains(strings.ToLower(info.Name()), "заказ") {
			files = append(files, path)
		}

		return nil
	})

	if error != nil {
		return nil, error
	}

	return files, nil
}

func main() {
	fmt.Println("RitaulPay запущен!")

	dir := getDir()
	files, _ := findOrderFiles(dir)

	for _, file := range files {
		fmt.Println("File: ", file)

		parser.Parse(file)
	}
}
