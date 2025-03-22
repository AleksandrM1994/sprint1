package helpers

import "unicode"

func RemoveControlCharacters(s string) string {
	result := []rune{}
	for _, r := range s {
		if !unicode.IsControl(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
