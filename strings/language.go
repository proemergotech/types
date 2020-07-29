package strings

type Language struct {
	lower
}

func NewLanguage(code string) Language {
	return Language{newLower(code)}
}
