package word

import "unicode"

func IsPalindrome(s string) bool {
	// Optimization: Pre-allocate the letters array with a sufficiently large array
	letters := make([]rune, 0, len(s))

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	// Optimization: End the comparisons after passing the middle of the string
	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}
