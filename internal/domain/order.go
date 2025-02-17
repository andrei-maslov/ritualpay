package domain

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
	Services []Service
}
