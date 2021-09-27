package result

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(tlv string) (Result, error) {
	t, l, v, err := ParseStringToResult(tlv)
	if err != nil {
		return Result{}, errors.New(err.Error())
	}
	return Result{Type: t, Length: l, Value: v}, nil
}

func ParseStringToResult(tlv string) (string, int, string, error) {
	if len(tlv) < 5 {
		return "", 0, "", errors.New("Text entered is too short, it needs to be at least 5 characters long")
	}
	t := tlv[0:2]
	lString := tlv[2:4]
	l, err := strconv.Atoi(lString)
	if err != nil {
		return "", 0, "", errors.New("Can't parse length value to int. Make sure 3rd and 4th digits are numbers")
	}
	if len(tlv) < 4+l {
		return "", 0, "", errors.New("Value length is longer than actual value length")
	}
	v := tlv[4 : 4+l]
	return t, l, v, nil
}
