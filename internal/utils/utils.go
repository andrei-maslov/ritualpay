package utils

import (
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

func ReportDir() string {
	return getOrderLocationDir() + "/Отчеты"
}

func ClearReportDir() error {
	return os.RemoveAll(ReportDir())
}

func CreateReportDir() error {
	return os.MkdirAll(ReportDir(), 0755)
}
