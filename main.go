package main

import (
	"fmt"

	"github.com/andrei-maslov/ritualpay/internal/domain"
	"github.com/andrei-maslov/ritualpay/internal/parser"
	"github.com/andrei-maslov/ritualpay/internal/utils"
)

func main() {
	fmt.Println("RitaulPay запущен!")

	files, _ := utils.GetOrderFiles()
	var orders []*domain.Order

	for _, file := range files {
		fmt.Println("File: ", file)

		order, err := parser.Parse(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		orders = append(orders, order)
	}

	for _, order := range orders {
		fmt.Println(order)
	}
}
