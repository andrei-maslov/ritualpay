package report

import (
	"fmt"
	"sort"
	"strings"

	"github.com/andrei-maslov/ritualpay/internal/domain"
)

// PerformerSummaryReport — отчет по суммам для каждого исполнителя
type PerformerSummaryReport struct {
	preformers []string
}

func (r *PerformerSummaryReport) Generate(orders []*domain.Order) string {
	// Считаем общую сумму для каждого исполнителя
	performerEarnings := make(map[string]int)

	for _, order := range orders {
		for _, service := range order.Services {
			for _, performer := range service.Performers {
				performerEarnings[performer] += service.PerformerPayout
			}
		}
	}

	// Сортируем исполнителей по убыванию суммы
	performers := make([]string, 0, len(performerEarnings))
	for performer := range performerEarnings {
		performers = append(performers, performer)
	}
	sort.Slice(performers, func(i, j int) bool {
		return performerEarnings[performers[i]] > performerEarnings[performers[j]]
	})

	// Формируем отчет
	var result strings.Builder
	result.WriteString("Отчет по суммам для исполнителей:\n")
	for _, performer := range performers {
		result.WriteString(fmt.Sprintf("├─ %s: %d руб.\n", performer, performerEarnings[performer]))
	}
	result.WriteString("└─ Всего исполнителей: " + fmt.Sprint(len(performers)) + "\n")

	keys := []string{}
	for k := range performerEarnings {
		keys = append(keys, k)
	}
	r.preformers = keys

	return result.String()
}

func (r *PerformerSummaryReport) GetPerformers() []string {
	return r.preformers
}
