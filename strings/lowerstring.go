package strings

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type lower string

func newLower(s string) lower {
	return lower(strings.ToLower(s))
}

func (l lower) String() string {
	return string(l)
}

func (l lower) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *lower) UnmarshalText(b []byte) error {
	*l = newLower(string(b))

	return nil
}

func (l lower) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(l.String())), nil
}

func (l *lower) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	*l = newLower(str)

	return nil
}

func (l lower) MarshalBinary() ([]byte, error) {
	return l.MarshalText()
}

func (l *lower) UnmarshalBinary(b []byte) error {
	return l.UnmarshalText(b)
}

func (l lower) Value() (driver.Value, error) {
	if l.String() == "" {
		return nil, nil
	}

	return l.String(), nil
}

func (l *lower) Scan(src interface{}) error {
	if src == nil {
		*l = ""
		return nil
	}

	if src, ok := src.(string); ok {
		*l = newLower(src)
	} else {
		return fmt.Errorf("cannot convert %T to lower", src)
	}

	return nil
}
