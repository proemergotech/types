package strings

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ugorji/go/codec"
)

func TestLowerNew(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue lower
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
			text:          "FoO",
			expectedValue: "foo",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			result := newLower(test.text)
			if !strings.EqualFold(string(result), string(test.expectedValue)) {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}

func TestLowerString(t *testing.T) {
	for index, test := range []struct {
		lo            lower
		expectedValue string
	}{
		{
			lo:            "",
			expectedValue: "",
		},
		{
			lo:            "foo",
			expectedValue: "foo",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.lo, test.expectedValue), func(t *testing.T) {
			result := test.lo.String()
			if result != test.expectedValue {
				t.Fatalf("expected: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}

func TestLowerMsgPack(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError error
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
			text:          "FoO",
			expectedValue: "foo",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			handle := &codec.MsgpackHandle{}

			var textB []byte
			err := codec.NewEncoderBytes(&textB, handle).Encode(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var lo lower
			err = codec.NewDecoderBytes(textB, handle).Decode(&lo)
			if err != nil {
				t.Fatal(err)
			}

			var b []byte
			err = codec.NewEncoderBytes(&b, handle).Encode(lo)
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

func TestLowerJSON(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError error
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
			text:          "FoO",
			expectedValue: "foo",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			textB, err := json.Marshal(test.text)
			if err != nil {
				t.Fatal(err)
			}

			var lo lower
			err = json.Unmarshal(textB, &lo)
			if err != nil {
				t.Fatal(err)
			}

			b, err := json.Marshal(lo)
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

func TestLowerSql(t *testing.T) {
	for index, test := range []struct {
		text          string
		expectedValue string
		expectedError error
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
			text:          "FoO",
			expectedValue: "foo",
		},
	} {
		t.Run(fmt.Sprintf("Case %d: %v -> %v", index+1, test.text, test.expectedValue), func(t *testing.T) {
			origLower := lower(test.text)
			driverValue, err := origLower.Value()
			if err != nil {
				t.Fatal(err)
			}

			s, ok := driverValue.(string)
			if !ok && test.text != "" {
				t.Fatalf("value does not returned with a string, returned: %T", driverValue)
			}

			var scanValue lower

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
