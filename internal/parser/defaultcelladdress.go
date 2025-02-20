package parser

type DefaultCellAddress struct {
}

func (c DefaultCellAddress) OrderSheetName() string {
	return "Заказ"
}

func (c DefaultCellAddress) TemplateVersionCell() string {
	return "M2"
}

func (c DefaultCellAddress) OrderNumberCell() string {
	return "AO13"
}

func (c DefaultCellAddress) CustomerFullNameCell() string {
	return "O16"
}

func (c DefaultCellAddress) HomePhomeCell() string {
	return "R17"
}

func (c DefaultCellAddress) MobilePhoneCell() string {
	return "AY17"
}

func (c DefaultCellAddress) AddressCell() string {
	return "H18"
}

func (c DefaultCellAddress) DeceasedFullNameCell() string {
	return "O19"
}

func (c DefaultCellAddress) DeceasedAgeCell() string {
	return "I20"
}

func (c DefaultCellAddress) DeceasedHeightCell() string {
	return "Z20"
}

func (c DefaultCellAddress) DeceasedClothingSizeCell() string {
	return "BA20"
}

func (c DefaultCellAddress) DeceasedBirthDateCell() string {
	return "N21"
}

func (c DefaultCellAddress) DeceasedDeathDateCell() string {
	return "AY21"
}

func (c DefaultCellAddress) ServicesRowNumbers() []string {
	return []string{
		"24", "25", "26", "27", "28", "29", "30", "31", "32", "33",
		"34", "35", "36", "37", "38", "39", "40", "41", "42", "43",
		"44", "45", "46", "47", "48", "49", "50", "51", "52",
		"57", "58", "59", "60", "61", "62", "63", "64", "65", "66",
		"67", "68", "69", "70", "71", "72", "73", "74", "75", "76",
		"77", "78",
	}
}

func (c DefaultCellAddress) ServiceNameColumn() string {
	return "E"
}

func (c DefaultCellAddress) CostColumn() string {
	return "AL"
}

func (c DefaultCellAddress) NoteColumn() string {
	return "BE"
}

func (c DefaultCellAddress) PerformerPayoutColumn() string {
	return "CC"
}

func (c DefaultCellAddress) PerformersColumns() []string {
	return []string{"CM", "CW", "DG", "DQ", "EA", "EK", "EU"}
}
