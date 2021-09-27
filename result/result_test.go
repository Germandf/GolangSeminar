package result

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResult(test *testing.T) {
	var cases = []struct {
		Input   string
		Success bool
	}{
		{"", false},
		{"NNAAA", false},
		{"NN32AAA", false},
		{"TX03AAAA", true},
		{"TX04AAAA", true},
	}
	for _, testData := range cases {
		_, err := NewResult(testData.Input)
		if testData.Success {
			assert.Nil(test, err)
		} else {
			assert.NotNil(test, err)
		}
	}
}

func TestParseStringToResult(test *testing.T) {
	var cases = []struct {
		Input   string
		Success bool
		Type    string
		Length  int
		Value   string
	}{
		{"", false, "", 0, ""},
		{"NNAAA", false, "", 0, ""},
		{"NN32AAA", false, "", 0, ""},
		{"TX03AAAA", true, "TX", 3, "AAA"},
		{"TX04AAAA", true, "TX", 4, "AAAA"},
	}
	for _, testData := range cases {
		t, l, v, err := ParseStringToResult(testData.Input)
		assert.EqualValues(test, testData.Type, t)
		assert.EqualValues(test, testData.Length, l)
		assert.EqualValues(test, testData.Value, v)
		if testData.Success {
			assert.Nil(test, err)
		} else {
			assert.NotNil(test, err)
		}
	}
}
