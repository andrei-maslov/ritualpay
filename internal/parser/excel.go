package parser

import (
	"fmt"

	"github.com/andrei-maslov/ritualpay/internal/domain"

	"github.com/xuri/excelize/v2"
)

const (
	OrderSheet = "Заказ"

	OrderNumberCellAdrs      = "AO13"
	CustomerFullNameCellAdrs = "O16"
	HomePhomeCellAdrs        = "R17"
	MobilePhoneCellAdrs      = "AY17"
	AddressCellAdrs          = "H18"

	DeceasedFullNameCellAdrs     = "O19"
	DeceasedAgeCellAdrs          = "I20"
	DeceasedHeightCellAdrs       = "Z20"
	DeceasedClothingSizeCellAdrs = "BA20"
	DeceasedBirthDateCellAdrs    = "N21"
	DeceasedDeathDateCellAdrs    = "AY21"
)

func Parse(filepath string) (*domain.Order, error) {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println("ERR: Ошибка открытия файла")
		return nil, nil
	}

	order := domain.Order{}

	err = parseOrderInfo(f, &order)
	if err != nil {
		// TODO Ругнуться
	}

	// err := parseServices(f, &order)

	return &order, err
}

func parseOrderInfo(f *excelize.File, order *domain.Order) error {
	if order == nil {
		fmt.Println("ERR: Получен пустой объект заказа!")
		return nil // TODO написать свою ошибку входного аргумента
	}

	cells := map[string]struct {
		cellAddress string
		field       *string
		title       string
	}{
		"order_number":           {OrderNumberCellAdrs, &order.OrderNumber, "Номер счет-заказа"},
		"customer_full_name":     {CustomerFullNameCellAdrs, &order.CustomerFullName, "ФИО заказчика"},
		"home_phone":             {HomePhomeCellAdrs, &order.HomePhone, "Добашний телефон"},
		"mobile_phone":           {MobilePhoneCellAdrs, &order.MobilePhone, "Мобильный телефон"},
		"address":                {AddressCellAdrs, &order.Address, "Адрес"},
		"deceased_full_name":     {DeceasedFullNameCellAdrs, &order.DeceasedFullName, "ФИО умершего"},
		"deceased_age":           {DeceasedAgeCellAdrs, &order.DeceasedAge, "Возраст"},
		"deceased_height":        {DeceasedHeightCellAdrs, &order.DeceasedHeight, "Рост"},
		"deceased_clothing_size": {DeceasedClothingSizeCellAdrs, &order.DeceasedClothingSize, "Размер одежды"},
		"deceased_birth_date":    {DeceasedBirthDateCellAdrs, &order.DeceasedBirthDate, "Дата рождения"},
		"deceased_death_date":    {DeceasedDeathDateCellAdrs, &order.DeceasedDeathDate, "Дата смерти"},
	}

	for key, cell := range cells {
		value, err := f.GetCellValue(OrderSheet, cell.cellAddress)
		if err != nil {
			return fmt.Errorf("ERR: Не удалось прочитать поле %s(%s) [%s] - %w", key, cell.title, cell.cellAddress, err)
		}
		*cell.field = value
	}

	return nil
}
