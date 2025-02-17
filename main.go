package main

import (
	"fmt"

	"github.com/andrei-maslov/ritualpay/internal/parser"
	"github.com/andrei-maslov/ritualpay/internal/utils"
)

func main() {
	fmt.Println("RitaulPay запущен!")

	files, _ := utils.GetOrderFiles()

	for _, file := range files {
		fmt.Println("File: ", file)

		parser.Parse(file)
	}
}
