package types

import (
	"errors"
	"regexp"

	"gitlab.com/proemergotech/dliver-types/internal"
)

var languageValidator = regexp.MustCompile(`^[A-Za-z]{2}$`)

// ISO 639-1 representation of language codes
type Language struct {
	internal.Lower
}

func NewLanguage(code string) (Language, error) {
	if err := validateLanguage(code); err != nil {
		return Language{}, err
	}

	return Language{internal.NewLower(code)}, nil
}

func (la Language) MarshalText() ([]byte, error) {
	if err := validateLanguage(la.String()); err != nil {
		return nil, err
	}

	return la.Lower.MarshalText()
}

func (la *Language) UnmarshalText(b []byte) error {
	if err := validateLanguage(string(b)); err != nil {
		return err
	}

	return la.Lower.UnmarshalText(b)
}

func (la Language) MarshalJSON() ([]byte, error) {
	if err := validateLanguage(la.String()); err != nil {
		return nil, err
	}

	return la.Lower.MarshalJSON()
}

func (la *Language) UnmarshalJSON(b []byte) error {
	return la.Lower.UnmarshalJSONWithValidate(b, validateLanguage)
}

func (la Language) MarshalBinary() ([]byte, error) {
	if err := validateLanguage(la.String()); err != nil {
		return nil, err
	}

	return la.Lower.MarshalBinary()
}

func (la *Language) UnmarshalBinary(b []byte) error {
	if err := validateLanguage(string(b)); err != nil {
		return err
	}

	return la.Lower.UnmarshalBinary(b)
}

func (la *Language) Scan(src interface{}) error {
	return la.Lower.ScanWithValidate(src, validateLanguage)
}

func validateLanguage(code string) error {
	if !languageValidator.MatchString(code) {
		return errors.New("language must contain alphabetic characters only with a length of 2")
	}

	return nil
}
