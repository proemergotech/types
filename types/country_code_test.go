package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ugorji/go/codec"
)

func TestCountryCodeNew(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue CountryCode
		expectedError string
	}{
		{
			text:          "",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "fo",
			expectedValue: CountryCode{"fo"},
		},
		{
			text:          "Fo",
			expectedValue: CountryCode{"fo"},
		},
		{
			text:          "Foo",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "12",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			result, err := NewCountryCode(test.text)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			}
			if !strings.EqualFold(result.String(), test.expectedValue.String()) {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}

func TestCountryCodeMsgPack(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "fo",
			expectedValue: "fo",
		},
		{
			text:          "Fo",
			expectedValue: "fo",
		},
		{
			text:          "Foo",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "12",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			handle := &codec.MsgpackHandle{}

			var textB []byte
			err := codec.NewEncoderBytes(&textB, handle).Encode(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var cc CountryCode
			err = codec.NewDecoderBytes(textB, handle).Decode(&cc)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			}

			var b []byte
			err = codec.NewEncoderBytes(&b, handle).Encode(&cc)
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

func TestCountryCodeJSON(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "fo",
			expectedValue: "fo",
		},
		{
			text:          "Fo",
			expectedValue: "fo",
		},
		{
			text:          "Foo",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "12",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			textB, err := json.Marshal(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var cc CountryCode
			err = json.Unmarshal(textB, &cc)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			}

			b, err := json.Marshal(cc)
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

func TestCountryCodeSql(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "fo",
			expectedValue: "fo",
		},
		{
			text:          "Fo",
			expectedValue: "fo",
		},
		{
			text:          "Foo",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
		{
			text:          "12",
			expectedError: "country code must contain alphabetic characters only with a length of 2",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			origCode, err := NewCountryCode(test.text)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			}
			driverValue, err := origCode.Value()
			if err != nil {
				t.Fatal(err)
			}

			s, ok := driverValue.(string)
			if !ok && test.text != "" {
				t.Fatalf("value does not returned with a string, returned: %T", driverValue)
			}

			var scanValue CountryCode

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
