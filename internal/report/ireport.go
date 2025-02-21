package report

import "github.com/andrei-maslov/ritualpay/internal/domain"

type Reporter interface {
	Generate(orders []*domain.Order) string
}
