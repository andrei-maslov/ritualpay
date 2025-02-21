package parser

import (
	"fmt"
	"strconv"
	"strings"

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
	fmt.Printf("Версия шаблона: %d\n\n", v)

	order := domain.Order{}

	err = parseOrderInfo(f, &order)
	if err != nil {
		fmt.Println("ERR: Ошибка чтения информации о заказе")
		return nil, err
	}

	err = parseServices(f, &order)
	if err != nil {
		fmt.Println("ERR: Ошибка чтения услуг")
		return nil, err
	}

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

func parseServices(f *excelize.File, o *domain.Order) error {
	for _, serviceRow := range cellAdrs.ServicesRowNumbers() {
		service := domain.Service{}

		serviceNameCellAdrs := cellAdrs.ServiceNameColumn() + serviceRow
		costCellAdrs := cellAdrs.CostColumn() + serviceRow
		noteCellAdrs := cellAdrs.PerformerPayoutColumn() + serviceRow
		performerPayoutCellAdrs := cellAdrs.PerformerPayoutColumn() + serviceRow

		name, err := f.GetCellValue(cellAdrs.OrderSheetName(), serviceNameCellAdrs)
		if err != nil {
			fmt.Println("ERR: Не удалось прочитать Наименование услуги")
			continue
		}
		service.Name = name

		value, err := f.GetCellValue(cellAdrs.OrderSheetName(), costCellAdrs)
		if err != nil {
			fmt.Println("ERR: Не удалось прочитать Цену услуги")
			continue
		}
		if strings.TrimSpace(value) != "" {
			cost, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("ERR: Не удалось преобразовать Цену услуги в число: " + value)
				continue
			}
			service.Cost = cost
		}

		note, err := f.GetCellValue(cellAdrs.OrderSheetName(), noteCellAdrs)
		if err != nil {
			fmt.Println("ERR: Не удалось прочитать примечание")
			continue
		}
		service.Note = note

		value, err = f.GetCellValue(cellAdrs.OrderSheetName(), performerPayoutCellAdrs)
		if err != nil {
			fmt.Println("ERR: Не удалось прочитать Стоимость выполнения работ")
			continue
		}
		if strings.TrimSpace(value) != "" {
			performerPayout, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("ERR: Не удалось преобразовать Стоимость выполнения работ в число: " + value)
				continue
			}
			service.PerformerPayout = performerPayout

			for i, performerColumn := range cellAdrs.PerformersColumns() {
				performerCellAdrs := performerColumn + serviceRow

				value, err = f.GetCellValue(cellAdrs.OrderSheetName(), performerCellAdrs)
				if err != nil {
					fmt.Println("ERR: Не удалось прочитать исполнителя ", i+1)
					continue
				}

				trimmedValue := strings.TrimSpace(value)
				if trimmedValue == "" {
					continue
				}
				service.Performers = append(service.Performers, trimmedValue)
			}
		}

		if service.Cost > 0 || service.PerformerPayout > 0 && len(service.Performers) > 0 {
			o.Services = append(o.Services, service)
		}
	}

	return nil
}
