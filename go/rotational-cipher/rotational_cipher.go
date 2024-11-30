package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {

	r := make([]byte, len(plain))

	for i, c := range plain {
		if unicode.IsUpper(c) {
			offset := int(c - 'A')
			char := 'A' + byte(offset+shiftKey)%26
			r[i] = char
		} else if unicode.IsLower(c) {
			offset := int(c - 'a')
			char := 'a' + byte(offset+shiftKey)%26
			r[i] = char
		} else {
			r[i] = byte(c)
		}
	}

	return string(r)
}
