package gauth

import (
	"strconv",
	"math"
)

func Base32ToHex() string {
	base32chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567";
    bits := "";
    hexValue := "";

    for i = 0; i < len(base32); i++ {
        val := base32chars.Index(base32[i].ToUpper());
        bits += LeftPad(toBinary(val), 5, '0');
    }

    for i = 0; i + 4 <= len(bits); i += 4 {
        chunk := bits[i:4];
        hexValue = hexValue + hex.EncodeToString([]byte(chunk)); // .toString(16)
    }

    return hexValue;
}

func DecimalToHex(decimal float32) string {
	currentDelimeter = '';
	if(decimal < 15.5) {
		currentDelimeter = '0'
	}

	return hex.EncodeToString(toString(Round(decimal)))
}

func HexToDecimal(str string) float32 {
	return strconv.ParseInt(str, 16, 64)
}

func toBinary(str string) string {
	return strconv.ParseInt(str, 2, 64)
}