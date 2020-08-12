package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var languageValidator = regexp.MustCompile(`^[A-Za-z]{2}$`)

// ISO 639-1 representation of language langs
type Language string

func NewLanguage(lang string) (Language, error) {
	if lang == "" {
		return "", nil
	}

	if !languageValidator.MatchString(lang) {
		return "", fmt.Errorf("invalid language: %s", lang)
	}

	return Language(strings.ToLower(lang)), nil
}

func (l Language) String() string {
	return string(l)
}

func (l Language) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Language) UnmarshalText(b []byte) error {
	lang, err := NewLanguage(string(b))
	if err != nil {
		return err
	}

	*l = lang

	return nil
}

func (l Language) MarshalJSON() ([]byte, error) {
	if l.String() == "" {
		return []byte("null"), nil
	}

	return []byte(strconv.Quote(l.String())), nil
}

func (l *Language) UnmarshalJSON(b []byte) error {
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

	*l = lang

	return nil
}

func (l Language) MarshalBinary() ([]byte, error) {
	return l.MarshalText()
}

func (l *Language) UnmarshalBinary(b []byte) error {
	return l.UnmarshalText(b)
}

func (l Language) Value() (driver.Value, error) {
	if l.String() == "" {
		return nil, nil
	}

	return l.String(), nil
}

func (l *Language) Scan(src interface{}) error {
	if src == nil {
		*l = ""
		return nil
	}

	if src, ok := src.(string); ok {
		var err error
		*l, err = NewLanguage(src)

		return err
	}

	return fmt.Errorf("cannot convert %T to Language", src)
}
