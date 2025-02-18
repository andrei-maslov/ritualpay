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
	MobilePhoneCellAdrs      = "AW17"
	AddressCellAdrs          = "H18"

	DeceasedFullNameCellAdrs     = "O19"
	DeceasedAgeCellAdrs          = "I20"
	DeceasedHeightCellAdrs       = "AR20"
	DeceasedClothingSizeCellAdrs = "AR20" // Такой же адрес как и у роста
	DeceasedBirthDateCellAdrs    = "W21"
	DeceasedDeathDateCellAdrs    = "W21" // TODO разделить в excel шаблоне эти поля
)

func Parse(filepath string) (*domain.Order, error) {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println("ERR: Ошибка открытия файла")
		return nil, nil
	}

	order := domain.Order{}

	err = parseOrderInfo(f, &order)

	return &order, err
}

func parseOrderInfo(f *excelize.File, order *domain.Order) error {
	if order == nil {
		fmt.Println("ERR: Получен пустой объект заказа!")
		return nil // TODO написать свою ошибку входного аргумента
	}

	orderNumber, err := f.GetCellValue(OrderSheet, OrderNumberCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение номера счета заказа")
		return err
	}

	customerFullName, err := f.GetCellValue(OrderSheet, CustomerFullNameCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение ФИО заказчика")
		return err
	}

	homePhone, err := f.GetCellValue(OrderSheet, HomePhomeCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Добашний телефон")
		return err
	}

	mobilePhone, err := f.GetCellValue(OrderSheet, MobilePhoneCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Мобильный телефон")
		return err
	}

	address, err := f.GetCellValue(OrderSheet, AddressCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Адрес")
		return err
	}

	deceasedFullName, err := f.GetCellValue(OrderSheet, DeceasedFullNameCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение ФИО умершего")
		return err
	}

	deceasedAge, err := f.GetCellValue(OrderSheet, DeceasedAgeCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Возраст умершего")
		return err
	}

	deceasedHeight, err := f.GetCellValue(OrderSheet, DeceasedHeightCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Рост умершего")
		return err
	}

	deceasedClothingSize, err := f.GetCellValue(OrderSheet, DeceasedClothingSizeCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Размер одежды умершего")
		return err
	}

	deceasedBirthDate, err := f.GetCellValue(OrderSheet, DeceasedBirthDateCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Дата рождения умершего")
		return err
	}

	deceasedDeathDate, err := f.GetCellValue(OrderSheet, DeceasedDeathDateCellAdrs)
	if err != nil {
		fmt.Println("ERR: Не удалось получить значение Дата смерти умершего")
		return err
	}

	fmt.Println("Номер закак: ", orderNumber)
	fmt.Println("ФИО заказчика: ", customerFullName)
	fmt.Println("Домашний телефон: ", homePhone)
	fmt.Println("Мобильный телефон: ", mobilePhone)
	fmt.Println("Адрес: ", address)
	fmt.Println("ФИО умершего: ", deceasedFullName)
	fmt.Println("Возраст умершего: ", deceasedAge)
	fmt.Println("Рост умершего: ", deceasedHeight)
	fmt.Println("Размер одежды умершего: ", deceasedClothingSize)
	fmt.Println("Дата рождения умершего: ", deceasedBirthDate)
	fmt.Println("Дата смерти умершего: ", deceasedDeathDate)

	order.OrderNumber = orderNumber
	order.CustomerFullName = customerFullName
	order.HomePhome = homePhone
	order.MobilePhone = mobilePhone
	order.Address = address
	order.DeceasedFullName = deceasedFullName
	// order.DeceasedAge = deceasedAge  // TODO convert to int8
	order.DeceasedHeight = deceasedHeight
	order.DeceasedClothingSize = deceasedClothingSize
	order.DeceasedBirthDate = deceasedBirthDate
	order.DeceasedDeathDate = deceasedDeathDate

	return nil
}
