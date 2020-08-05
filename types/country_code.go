package types

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"gitlab.com/proemergotech/dliver-types/internal"
)

var countryCodeValidator = regexp.MustCompile(`(^[A-Za-z]{2}$)|(^[tT]1$)`)

// ISO 3166-1 Alpha-2 representation of country codes
type CountryCode struct {
	internal.Lower
}

func NewCountryCode(code string) (CountryCode, error) {
	if code == "" {
		return CountryCode{}, nil
	}

	if !countryCodeValidator.MatchString(code) {
		return CountryCode{}, fmt.Errorf("invalid country code: %s", code)
	}

	return CountryCode{internal.NewLower(code)}, nil
}

func (cc CountryCode) MarshalText() ([]byte, error) {
	return cc.Lower.MarshalText()
}

func (cc *CountryCode) UnmarshalText(b []byte) error {
	code, err := NewCountryCode(string(b))
	if err != nil {
		return err
	}

	*cc = code

	return nil
}

func (cc CountryCode) MarshalJSON() ([]byte, error) {
	return cc.Lower.MarshalJSON()
}

func (cc *CountryCode) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	code, err := NewCountryCode(str)
	if err != nil {
		return err
	}

	*cc = code

	return nil
}

func (cc CountryCode) MarshalBinary() ([]byte, error) {
	return cc.Lower.MarshalBinary()
}

func (cc *CountryCode) UnmarshalBinary(b []byte) error {
	return cc.UnmarshalText(b)
}

func (cc *CountryCode) Scan(src interface{}) error {
	if src == nil {
		*cc = CountryCode{}
		return nil
	}

	if src, ok := src.(string); ok {
		var err error
		*cc, err = NewCountryCode(src)

		return err
	}

	return fmt.Errorf("cannot convert %T to CountryCode", src)
}
