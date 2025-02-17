package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getOrderLocationDir() string {
	file, _ := os.Executable()
	return filepath.Dir(file)
}

func GetOrderFiles() ([]string, error) {
	var files []string

	error := filepath.Walk(getOrderLocationDir(), func(path string, info os.FileInfo, err error) error {
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
