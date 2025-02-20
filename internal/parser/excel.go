package parser

import (
	"fmt"
	"strconv"

	"github.com/andrei-maslov/ritualpay/internal/domain"
	// "github.com/andrei-maslov/ritaulpay/internal/parser/"

	"github.com/xuri/excelize/v2"
)

var cellAdrs ICellAddress = DefaultCellAddress{}

func Parse(filepath string) (*domain.Order, error) {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println("ERR: Ошибка открытия файла")
		return nil, nil
	}

	v, err := parseTemplateVersion(f)
	if err != nil {
		fmt.Println("ERR: Ошибка чтения версии файла")
		return nil, err
	}
	fmt.Printf("Версия шаблона: %d", v)

	order := domain.Order{}

	err = parseOrderInfo(f, &order)
	if err != nil {
		fmt.Println("ERR: Ошибка чтения информации о заказе")
		return nil, err
	}

	err = parseServices(f, &order)

	return &order, err
}

func parseTemplateVersion(f *excelize.File) (int, error) {
	strV, err := f.GetCellValue(cellAdrs.OrderSheetName(), cellAdrs.TemplateVersionCell())
	if err != nil {
		return -1, err
	}

	v, err := strconv.Atoi(strV)
	if err != nil {
		return -1, err
	}
	return v, nil
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
		"order_number":           {cellAdrs.OrderNumberCell(), &order.OrderNumber, "Номер счет-заказа"},
		"customer_full_name":     {cellAdrs.CustomerFullNameCell(), &order.CustomerFullName, "ФИО заказчика"},
		"home_phone":             {cellAdrs.HomePhomeCell(), &order.HomePhone, "Добашний телефон"},
		"mobile_phone":           {cellAdrs.MobilePhoneCell(), &order.MobilePhone, "Мобильный телефон"},
		"address":                {cellAdrs.AddressCell(), &order.Address, "Адрес"},
		"deceased_full_name":     {cellAdrs.DeceasedFullNameCell(), &order.DeceasedFullName, "ФИО умершего"},
		"deceased_age":           {cellAdrs.DeceasedAgeCell(), &order.DeceasedAge, "Возраст"},
		"deceased_height":        {cellAdrs.DeceasedHeightCell(), &order.DeceasedHeight, "Рост"},
		"deceased_clothing_size": {cellAdrs.DeceasedClothingSizeCell(), &order.DeceasedClothingSize, "Размер одежды"},
		"deceased_birth_date":    {cellAdrs.DeceasedBirthDateCell(), &order.DeceasedBirthDate, "Дата рождения"},
		"deceased_death_date":    {cellAdrs.DeceasedDeathDateCell(), &order.DeceasedDeathDate, "Дата смерти"},
	}

	for key, cell := range cells {
		value, err := f.GetCellValue(cellAdrs.OrderSheetName(), cell.cellAddress)
		if err != nil {
			return fmt.Errorf("ERR: Не удалось прочитать поле %s(%s) [%s] - %w", key, cell.title, cell.cellAddress, err)
		}
		*cell.field = value
	}

	return nil
}
