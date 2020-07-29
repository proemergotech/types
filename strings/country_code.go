package strings

type CountryCode struct {
	lower
}

func NewCountryCode(code string) CountryCode {
	return CountryCode{newLower(code)}
}
