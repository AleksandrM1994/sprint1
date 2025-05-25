package helpers

import "unicode"

// RemoveControlCharacters удаляет все управляющие символы из строки
func RemoveControlCharacters(s string) string {
	result := []rune{}
	for _, r := range s {
		if !unicode.IsControl(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
