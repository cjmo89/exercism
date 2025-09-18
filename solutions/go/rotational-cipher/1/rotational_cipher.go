package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	newStr := strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			newRune := r + rune(shiftKey)%26
			if newRune > 'Z' {
				newRune = newRune - 'Z' + 'A' - 1
			}
			return newRune
		}
		if r >= 'a' && r <= 'z' {
			newRune := r + rune(shiftKey)%26
			if newRune > 'z' {
				newRune = newRune - 'z' + 'a' - 1
			}
			return newRune
		}
		return r
	}, plain)
	return newStr
}
