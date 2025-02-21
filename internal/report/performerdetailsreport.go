package report

import (
	"fmt"
	"strings"

	"github.com/andrei-maslov/ritualpay/internal/domain"
)

type _order struct {
	Number           string
	CustomerFullName string
	Services         []_service
}

type _service struct {
	Name            string
	PerformerPayotu int
}

// PerformerDetailsReport — отчет по конкретному исполнителю
type PerformerDetailsReport struct {
	PerformerName string // Имя исполнителя, для которого формируем отчет

	orders []_order
}

func (r *PerformerDetailsReport) Generate(orders []*domain.Order) string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("Отчет по исполнителю: %s\n", r.PerformerName))

	for _, order := range orders {
		var services []_service
		for _, s := range order.Services {
			for _, performer := range s.Performers {
				if performer == r.PerformerName {
					service := _service{Name: s.Name, PerformerPayotu: s.PerformerPayout}
					services = append(services, service)
				}
			}
		}
		if len(services) > 0 {
			o := _order{Number: order.OrderNumber,
				CustomerFullName: order.CustomerFullName,
				Services:         services}
			r.orders = append(r.orders, o)
		}
	}

	for _, order := range r.orders {

		result.WriteString(fmt.Sprintf("├─ Заказ №%s для %s\n", order.Number, order.CustomerFullName))
		for i, service := range order.Services {
			prefix := ""
			if i != len(order.Services)-1 {
				prefix = "├"
			} else {
				prefix = "└"
			}
			result.WriteString(fmt.Sprintf("│  %s- %s:  %d\n", prefix, service.Name, service.PerformerPayotu))
		}
	}

	if !strings.Contains(result.String(), "Заказ №") {
		result.WriteString("└─ Исполнитель не участвовал в заказах.\n")
	}

	return result.String()
}
