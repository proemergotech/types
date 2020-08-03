package types

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"gitlab.com/proemergotech/dliver-types/internal"
)

var languageValidator = regexp.MustCompile(`^[A-Za-z]{2}$`)

// ISO 639-1 representation of language codes
type Language struct {
	internal.Lower
}

func NewLanguage(code string) (Language, error) {
	if code == "" {
		return Language{}, nil
	}

	if !languageValidator.MatchString(code) {
		return Language{}, fmt.Errorf("invalid language: %s", code)
	}

	return Language{internal.NewLower(code)}, nil
}

func (la Language) MarshalText() ([]byte, error) {
	return la.Lower.MarshalText()
}

func (la *Language) UnmarshalText(b []byte) error {
	lang, err := NewLanguage(string(b))
	if err != nil {
		return err
	}

	*la = lang

	return nil
}

func (la Language) MarshalJSON() ([]byte, error) {
	return la.Lower.MarshalJSON()
}

func (la *Language) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	lang, err := NewLanguage(str)
	if err != nil {
		return err
	}

	*la = lang

	return nil
}

func (la Language) MarshalBinary() ([]byte, error) {
	return la.Lower.MarshalBinary()
}

func (la *Language) UnmarshalBinary(b []byte) error {
	return la.UnmarshalText(b)
}

func (la *Language) Scan(src interface{}) error {
	if src == nil {
		*la = Language{}
		return nil
	}

	if src, ok := src.(string); ok {
		var err error
		*la, err = NewLanguage(src)

		return err
	}

	return fmt.Errorf("cannot convert %T to Language", src)
}
