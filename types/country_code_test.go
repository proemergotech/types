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
			expectedValue: "",
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
			text:          "t1",
			expectedValue: "t1",
		},
		{
			text:          "Foo",
			expectedError: "invalid country code",
		},
		{
			text:          "12",
			expectedError: "invalid country code",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			result, err := NewCountryCode(test.text)
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

func TestCountryCodeString(t *testing.T) {
	for index, test := range []struct {
		code          CountryCode
		expectedValue string
	}{
		{
			code:          "",
			expectedValue: "",
		},
		{
			code:          "foo",
			expectedValue: "foo",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.code, test.expectedValue), func(t *testing.T) {
			result := test.code.String()
			if result != test.expectedValue {
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
			expectedValue: "",
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
			expectedError: "invalid country code",
		},
		{
			text:          "12",
			expectedError: "invalid country code",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			handle := &codec.MsgpackHandle{}

			var textB []byte
			err := codec.NewEncoderBytes(&textB, handle).Encode(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var code CountryCode
			err = codec.NewDecoderBytes(textB, handle).Decode(&code)
			if err != nil {
				if test.expectedError != "" && strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			} else if test.expectedError != "" {
				t.Errorf("expected error: %s, got none", test.expectedError)
			}

			var b []byte
			err = codec.NewEncoderBytes(&b, handle).Encode(&code)
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
			expectedValue: "",
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
			expectedError: "invalid country code",
		},
		{
			text:          "12",
			expectedError: "invalid country code",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			textB, err := json.Marshal(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var code CountryCode
			err = json.Unmarshal(textB, &code)
			if err != nil {
				if test.expectedError != "" && strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			} else if test.expectedError != "" {
				t.Errorf("expected error: %s, got none", test.expectedError)
			}

			b, err := json.Marshal(code)
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
			expectedValue: "",
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
			expectedError: "invalid country code",
		},
		{
			text:          "12",
			expectedError: "invalid country code",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			origCode, err := NewCountryCode(test.text)
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
