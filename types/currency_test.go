package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ugorji/go/codec"
)

func TestCurrencyNew(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue Currency
		expectedError string
	}{
		{
			text:          "",
			expectedValue: "",
		},
		{
			text:          "foo",
			expectedValue: "foo",
		},
		{
			text:          "Foo",
			expectedValue: "foo",
		},
		{
			text:          "Fo1",
			expectedError: "invalid currency",
		},
		{
			text:          "Fo",
			expectedError: "invalid currency",
		},
		{
			text:          "123",
			expectedError: "invalid currency",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			result, err := NewCurrency(test.text)
			if err != nil {
				if test.expectedError != "" && strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			} else if test.expectedError != "" {
				t.Errorf("expected error: %s, got none", test.expectedError)
			}
			if !strings.EqualFold(result.String(), test.expectedValue.String()) {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}

func TestCurrencyString(t *testing.T) {
	for index, test := range []struct {
		currency      Currency
		expectedValue string
	}{
		{
			currency:      "",
			expectedValue: "",
		},
		{
			currency:      "foo",
			expectedValue: "foo",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.currency, test.expectedValue), func(t *testing.T) {
			result := test.currency.String()
			if result != test.expectedValue {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}

func TestCurrencyMsgPack(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedValue: "",
		},
		{
			text:          "foo",
			expectedValue: "foo",
		},
		{
			text:          "Foo",
			expectedValue: "foo",
		},
		{
			text:          "Fo1",
			expectedError: "invalid currency",
		},
		{
			text:          "Fo",
			expectedError: "invalid currency",
		},
		{
			text:          "123",
			expectedError: "invalid currency",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			handle := &codec.MsgpackHandle{}

			var textB []byte
			err := codec.NewEncoderBytes(&textB, handle).Encode(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var currency Currency
			err = codec.NewDecoderBytes(textB, handle).Decode(&currency)
			if err != nil {
				if test.expectedError != "" && strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			} else if test.expectedError != "" {
				t.Errorf("expected error: %s, got none", test.expectedError)
			}

			var b []byte
			err = codec.NewEncoderBytes(&b, handle).Encode(&currency)
			if err != nil {
				t.Fatal(err)
			}

			var str string
			err = codec.NewDecoderBytes(b, handle).Decode(&str)
			if err != nil {
				t.Fatal(err)
			}

			if str != test.expectedValue {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, str)
			}
		})
	}
}

func TestCurrencyJSON(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedValue: "",
		},
		{
			text:          "foo",
			expectedValue: "foo",
		},
		{
			text:          "Foo",
			expectedValue: "foo",
		},
		{
			text:          "Fo1",
			expectedError: "invalid currency",
		},
		{
			text:          "Fo",
			expectedError: "invalid currency",
		},
		{
			text:          "123",
			expectedError: "invalid currency",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			textB, err := json.Marshal(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var currency Currency
			err = json.Unmarshal(textB, &currency)
			if err != nil {
				if test.expectedError != "" && strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			} else if test.expectedError != "" {
				t.Errorf("expected error: %s, got none", test.expectedError)
			}

			b, err := json.Marshal(currency)
			if err != nil {
				t.Fatal(err)
			}

			var str string
			err = json.Unmarshal(b, &str)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.EqualFold(str, test.expectedValue) {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, str)
			}
		})
	}
}

func TestCurrencySql(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedValue: "",
		},
		{
			text:          "foo",
			expectedValue: "foo",
		},
		{
			text:          "Foo",
			expectedValue: "foo",
		},
		{
			text:          "Fo1",
			expectedError: "invalid currency",
		},
		{
			text:          "Fo",
			expectedError: "invalid currency",
		},
		{
			text:          "123",
			expectedError: "invalid currency",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			origCode, err := NewCurrency(test.text)
			if err != nil {
				if test.expectedError != "" && strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			} else if test.expectedError != "" {
				t.Errorf("expected error: %s, got none", test.expectedError)
			}
			driverValue, err := origCode.Value()
			if err != nil {
				t.Fatal(err)
			}

			s, ok := driverValue.(string)
			if !ok && test.text != "" {
				t.Fatalf("value does not returned with a string, returned: %T", driverValue)
			}

			var scanValue Currency

			if s == "" {
				err = scanValue.Scan(nil)
			} else {
				err = scanValue.Scan(s)
			}

			if err != nil {
				t.Fatal(err)
			}

			if scanValue.String() != test.expectedValue {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, scanValue.String())
			}
		})
	}
}
