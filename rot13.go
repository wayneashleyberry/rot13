package rot13

import "strings"

func indexPosition(s string, arr []string) int {
	for i, c := range arr {
		if c == s {
			return i
		}
	}
	return -1
}

// Encode takes a plain text message and returns the rot13 encoded result.
func Encode(message string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var encoded []string
	for _, c := range message {
		isUpper := false
		if strings.ToLower(string(c)) != string(c) {
			isUpper = true
		}

		i := indexPosition(strings.ToLower(string(c)), alphabet)
		if i == -1 {
			encoded = append(encoded, string(c))
			continue
		}

		var pos int
		pos = i + 13

		if pos >= len(alphabet) {
			pos = (i + 13) - len(alphabet)
		}

		if isUpper {
			encoded = append(encoded, strings.ToUpper(alphabet[pos]))
			continue
		}
		encoded = append(encoded, alphabet[pos])
	}
	return strings.Join(encoded, "")
}

// Decode takes a rot13 encoded message and returns the unencoded result.
func Decode(message string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var decoded []string
	for _, c := range message {
		isUpper := false
		if strings.ToLower(string(c)) != string(c) {
			isUpper = true
		}

		i := indexPosition(strings.ToLower(string(c)), alphabet)
		if i == -1 {
			decoded = append(decoded, string(c))
			continue
		}

		var pos int
		pos = i - 13

		if pos < 0 {
			pos = len(alphabet) + (i - 13)
		}

		if isUpper {
			decoded = append(decoded, strings.ToUpper(alphabet[pos]))
			continue
		}
		decoded = append(decoded, alphabet[pos])
	}
	return strings.Join(decoded, "")
}
