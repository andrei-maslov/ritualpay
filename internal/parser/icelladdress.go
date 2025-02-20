package parser

type ICellAddress interface {
	OrderSheetName() string

	// Версия шаблона excel шаблона
	TemplateVersionCell() string
	// номер счет-заказа
	OrderNumberCell() string
	// ФИО заказчика
	CustomerFullNameCell() string
	// Домашний номер
	HomePhomeCell() string
	// Мобильный номер
	MobilePhoneCell() string
	// Адрес
	AddressCell() string
	// Данные об умершем
	// ФИО
	DeceasedFullNameCell() string
	// Возраст
	DeceasedAgeCell() string
	// Рост
	DeceasedHeightCell() string
	// Размер одежды
	DeceasedClothingSizeCell() string
	// Дата рождения
	DeceasedBirthDateCell() string
	// Дата смерти
	DeceasedDeathDateCell() string

	// Номера строка с товарами и услугами
	ServicesRowNumbers() []string
	// Наименование
	ServiceNameColumn() string
	// Цена
	CostColumn() string
	// Примечание
	NoteColumn() string
	// Оплата одному исполнителю
	PerformerPayoutColumn() string
	// Исполнители
	PerformersColumns() []string
}
