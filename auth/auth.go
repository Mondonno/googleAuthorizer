package gauth

import (
	"time",
	"crypto/hmac"
    "crypto/sha256"
	"encoding/hex"
	"encoding/base32"
)

func GenerateKey(key string) string {
	validKey := Base32ToHex(key)
	if(len(validKey) % 2 == 0) {
		validKey += "0"
	}

	currentMl := time.Now().UTC().UnixNano() / 1e6
	epoch := currentMl / 1000

	currentTime := LeftPad(DecimalToHex(math.floor(epoch)), 16, "0")

	hmacObj := hmac.New(sha256.New, []byte(validKey))
	hmacObj.Write([]byte(currentTime))

	hmacSha := hmacObj.sum(nil)

	offset := 0
	otp := (HexToDecimal(hmacSha[(offset * 2):8]) & HexToDecimal("7fffffff")) + ""
	return toString((otp)[(len(otp) - 6):6])
}