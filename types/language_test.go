package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ugorji/go/codec"
)

func TestLanguageNew(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue Language
		expectedError string
	}{
		{
			text:          "",
			expectedError: "invalid language",
		},
		{
			text:          "fo",
			expectedValue: Language{"fo"},
		},
		{
			text:          "Fo",
			expectedValue: Language{"fo"},
		},
		{
			text:          "t1",
			expectedValue: Language{"t1"},
		},
		{
			text:          "Foo",
			expectedError: "invalid language",
		},
		{
			text:          "12",
			expectedError: "invalid language",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			result, err := NewLanguage(test.text)
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

func TestLanguageMsgPack(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedError: "invalid language",
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
			expectedError: "invalid language",
		},
		{
			text:          "12",
			expectedError: "invalid language",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			handle := &codec.MsgpackHandle{}

			var textB []byte
			err := codec.NewEncoderBytes(&textB, handle).Encode(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var la Language
			err = codec.NewDecoderBytes(textB, handle).Decode(&la)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			}

			var b []byte
			err = codec.NewEncoderBytes(&b, handle).Encode(&la)
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

func TestLanguageJSON(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedError: "invalid language",
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
			expectedError: "invalid language",
		},
		{
			text:          "12",
			expectedError: "invalid language",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			textB, err := json.Marshal(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var la Language
			err = json.Unmarshal(textB, &la)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Fatal(err)
			}

			b, err := json.Marshal(la)
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

func TestLanguageSql(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError string
	}{
		{
			text:          "",
			expectedError: "invalid language",
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
			expectedError: "invalid language",
		},
		{
			text:          "12",
			expectedError: "invalid language",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			origCode, err := NewLanguage(test.text)
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

			var scanValue Language

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
