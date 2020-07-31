package types

import (
	"errors"
	"regexp"

	"gitlab.com/proemergotech/dliver-types/internal"
)

var countryCodeValidator = regexp.MustCompile(`^[A-Za-z]{2}$`)

// ISO 3166-1 Alpha-2 representation of country codes
type CountryCode struct {
	internal.Lower
}

func NewCountryCode(code string) (CountryCode, error) {
	if err := validateCountryCode(code); err != nil {
		return CountryCode{}, err
	}

	return CountryCode{internal.NewLower(code)}, nil
}

func (cc CountryCode) MarshalText() ([]byte, error) {
	if err := validateCountryCode(cc.String()); err != nil {
		return nil, err
	}

	return cc.Lower.MarshalText()
}

func (cc *CountryCode) UnmarshalText(b []byte) error {
	if err := validateCountryCode(string(b)); err != nil {
		return err
	}

	return cc.Lower.UnmarshalText(b)
}

func (cc CountryCode) MarshalJSON() ([]byte, error) {
	if err := validateCountryCode(cc.String()); err != nil {
		return nil, err
	}

	return cc.Lower.MarshalJSON()
}

func (cc *CountryCode) UnmarshalJSON(b []byte) error {
	return cc.Lower.UnmarshalJSONWithValidate(b, validateCountryCode)
}

func (cc CountryCode) MarshalBinary() ([]byte, error) {
	if err := validateCountryCode(cc.String()); err != nil {
		return nil, err
	}

	return cc.Lower.MarshalBinary()
}

func (cc *CountryCode) UnmarshalBinary(b []byte) error {
	if err := validateCountryCode(string(b)); err != nil {
		return err
	}

	return cc.Lower.UnmarshalBinary(b)
}

func (cc *CountryCode) Scan(src interface{}) error {
	return cc.Lower.ScanWithValidate(src, validateCountryCode)
}

func validateCountryCode(code string) error {
	if !countryCodeValidator.MatchString(code) {
		return errors.New("country code must contain alphabetic characters only with a length of 2")
	}

	return nil
}
