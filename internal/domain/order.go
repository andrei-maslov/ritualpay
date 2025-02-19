package domain

import (
	"fmt"
	"strings"
)

// Исполнитель работы / услуги
type Preformer struct {
	Name string
}

// Услуга
type Service struct {
	Name       string
	Cost       float64
	Performers []Preformer
}

// Заказ
type Order struct {
	OrderNumber      string
	CustomerFullName string
	HomePhone        string
	MobilePhone      string
	Address          string

	DeceasedFullName     string
	DeceasedAge          string
	DeceasedHeight       string // тип string на случай если в excel файле запишут "1м 65см"
	DeceasedClothingSize string
	DeceasedBirthDate    string // тип так же string так как не знаю в каком формате будут писать дату
	DeceasedDeathDate    string // аналогично :point_up:

	Services []Service
}

func (o Order) String() string {
	var sb strings.Builder
	sb.WriteString("Order Details:\n")

	sb.WriteString(fmt.Sprintf("├─ Номер счета:      %s\n", o.OrderNumber))
	sb.WriteString(fmt.Sprintf("├─ Заказчик:         %s\n", o.CustomerFullName))
	sb.WriteString(fmt.Sprintf("├─ Контакты:         %s / %s\n", o.HomePhone, o.MobilePhone))
	sb.WriteString(fmt.Sprintf("├─ Адрес:            %s\n", o.Address))
	sb.WriteString(fmt.Sprintf("├─ Умерший:          %s\n", o.DeceasedFullName))
	sb.WriteString(fmt.Sprintf("├─ Возраст/Рост:     %s лет / %s см\n", o.DeceasedAge, o.DeceasedHeight))
	sb.WriteString(fmt.Sprintf("├─ Размер одежды:    %s\n", o.DeceasedClothingSize))
	sb.WriteString(fmt.Sprintf("├─ Даты:             %s - %s\n", o.DeceasedBirthDate, o.DeceasedDeathDate))
	sb.WriteString(fmt.Sprintf("└─ Услуги (%d):\n", len(o.Services)))

	for i, s := range o.Services {
		sb.WriteString(fmt.Sprintf("   %d. %s (%.2f руб)\n", i+1, s.Name, s.Cost))
	}

	return sb.String()
}
