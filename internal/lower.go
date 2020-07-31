package internal

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Lower string

func NewLower(s string) Lower {
	return Lower(strings.ToLower(s))
}

func (l Lower) String() string {
	return string(l)
}

func (l Lower) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Lower) UnmarshalText(b []byte) error {
	*l = NewLower(string(b))

	return nil
}

func (l Lower) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(l.String())), nil
}

func (l *Lower) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	*l = NewLower(str)

	return nil
}

func (l *Lower) UnmarshalJSONWithValidate(b []byte, validate func(string) error) error {
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	if err := validate(str); err != nil {
		return err
	}

	*l = NewLower(str)

	return nil
}

func (l Lower) MarshalBinary() ([]byte, error) {
	return l.MarshalText()
}

func (l *Lower) UnmarshalBinary(b []byte) error {
	return l.UnmarshalText(b)
}

func (l Lower) Value() (driver.Value, error) {
	if l.String() == "" {
		return nil, nil
	}

	return l.String(), nil
}

func (l *Lower) Scan(src interface{}) error {
	if src == nil {
		*l = ""
		return nil
	}

	if src, ok := src.(string); ok {
		*l = NewLower(src)
	} else {
		return fmt.Errorf("cannot convert %T to target type", src)
	}

	return nil
}

func (l *Lower) ScanWithValidate(src interface{}, validate func(string) error) error {
	if src == nil {
		*l = ""
		return nil
	}

	if src, ok := src.(string); ok {
		if err := validate(src); err != nil {
			return err
		}
		*l = NewLower(src)
	} else {
		return fmt.Errorf("cannot convert %T to target type", src)
	}

	return nil
}
