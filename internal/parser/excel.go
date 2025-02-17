package parser

import (
	"fmt"

	"github.com/andrei-maslov/ritualpay/internal/domain"

	"github.com/xuri/excelize/v2"
)

func Parse(filepath string) (*domain.Order, error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println("ERR: Ошибка открытия файла")
	}

	cell, err := f.GetCellValue("Заказ", "O16")
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение ячейки")
	}

	fmt.Println("Заказчик: ", cell)

	return nil, nil
}
