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
	OrderNumber      string
	CustomerFullName string
	HomePhome        string
	MobilePhone      string
	Address          string

	DeceasedFullName     string
	DeceasedAge          int8
	DeceasedHeight       string // тип string на случай если в excel файле запишут "1м 65см"
	DeceasedClothingSize string
	DeceasedBirthDate    string // тип так же string так как не знаю в каком формате будут писать дату
	DeceasedDeathDate    string // аналогично :point_up:

	Services []Service
}
