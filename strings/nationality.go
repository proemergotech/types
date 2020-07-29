package strings

type Nationality struct {
	lower
}

func NewNationality(code string) Nationality {
	return Nationality{newLower(code)}
}
