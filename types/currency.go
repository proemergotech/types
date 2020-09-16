package types

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var currencyValidator = regexp.MustCompile(`^[A-Za-z]{3}$`)

// ISO 4217 alphabetical currency code
type Currency string

func NewCurrency(currency string) (Currency, error) {
	if currency == "" {
		return "", nil
	}

	if !currencyValidator.MatchString(currency) {
		return "", fmt.Errorf("invalid currency: %s", currency)
	}

	return Currency(strings.ToLower(currency)), nil
}

func (c Currency) String() string {
	return string(c)
}

func (c Currency) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *Currency) UnmarshalText(b []byte) error {
	currency, err := NewCurrency(string(b))
	if err != nil {
		return err
	}

	*c = currency

	return nil
}

func (c Currency) MarshalJSON() ([]byte, error) {
	if c.String() == "" {
		return []byte("null"), nil
	}

	return []byte(strconv.Quote(c.String())), nil
}

func (c *Currency) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	currency, err := NewCurrency(str)
	if err != nil {
		return err
	}

	*c = currency

	return nil
}

func (c Currency) MarshalBinary() ([]byte, error) {
	return c.MarshalText()
}

func (c *Currency) UnmarshalBinary(b []byte) error {
	return c.UnmarshalText(b)
}

func (c Currency) Value() (driver.Value, error) {
	if c.String() == "" {
		return nil, nil
	}

	return c.String(), nil
}

func (c *Currency) Scan(src interface{}) error {
	if src == nil {
		*c = ""
		return nil
	}

	if src, ok := src.(string); ok {
		var err error
		*c, err = NewCurrency(src)

		return err
	}

	return fmt.Errorf("cannot convert %T to Currency", src)
}
