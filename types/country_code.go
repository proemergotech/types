package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var countryCodeValidator = regexp.MustCompile(`(^[A-Za-z]{2}$)|(^[tT]1$)`)

// ISO 3166-1 Alpha-2 representation of country codes. T1 represents tor exit node
type CountryCode string

func NewCountryCode(code string) (CountryCode, error) {
	if code == "" {
		return "", nil
	}

	if !countryCodeValidator.MatchString(code) {
		return "", fmt.Errorf("invalid country code: %s", code)
	}

	return CountryCode(strings.ToLower(code)), nil
}

func (c CountryCode) String() string {
	return string(c)
}

func (c CountryCode) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *CountryCode) UnmarshalText(b []byte) error {
	code, err := NewCountryCode(string(b))
	if err != nil {
		return err
	}

	*c = code

	return nil
}

func (c CountryCode) MarshalJSON() ([]byte, error) {
	if c.String() == "" {
		return []byte("null"), nil
	}

	return []byte(strconv.Quote(c.String())), nil
}

func (c *CountryCode) UnmarshalJSON(b []byte) error {
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

	*c = code

	return nil
}

func (c CountryCode) MarshalBinary() ([]byte, error) {
	return c.MarshalText()
}

func (c *CountryCode) UnmarshalBinary(b []byte) error {
	return c.UnmarshalText(b)
}

func (c CountryCode) Value() (driver.Value, error) {
	if c.String() == "" {
		return nil, nil
	}

	return c.String(), nil
}

func (c *CountryCode) Scan(src interface{}) error {
	if src == nil {
		*c = ""
		return nil
	}

	if src, ok := src.(string); ok {
		var err error
		*c, err = NewCountryCode(src)

		return err
	}

	return fmt.Errorf("cannot convert %T to CountryCode", src)
}
