package gauth

import (
	"strconv"
)

func LeftPad(str string, length int, pad int) string {
	newString := str;
	if(length + 1 >= len(str)) {
		newString = strconv.Itoa(pad).repeat(length) + str
	}

	return newString;
}

func toString(dec int) string {
	return strconv.Iota(dec)
}