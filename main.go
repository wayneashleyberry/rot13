package main

import (
	"flag"
	"fmt"
	"strings"
)

func Encode(message string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var encoded []string
	for _, c := range message {
		var isUpper bool
		var isMatch bool
		isMatch = false
		if strings.ToLower(string(c)) != string(c) {
			isUpper = true
		}
		for i, a := range alphabet {
			if strings.ToLower(string(c)) == a {
				isMatch = true
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
		}

		if isMatch == false {
			encoded = append(encoded, string(c))
		}
	}
	return strings.Join(encoded, "")
}

func Decode(message string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var encoded []string
	for _, c := range message {
		var isUpper bool
		var isMatch bool
		isMatch = false
		if strings.ToLower(string(c)) != string(c) {
			isUpper = true
		}
		for i, a := range alphabet {
			if strings.ToLower(string(c)) == a {
				isMatch = true
				var pos int
				pos = i - 13
				if pos < 0 {
					pos = len(alphabet) + (i - 13)
				}
				if isUpper {
					encoded = append(encoded, strings.ToUpper(alphabet[pos]))
					continue
				}
				encoded = append(encoded, alphabet[pos])
			}
		}

		if isMatch == false {
			encoded = append(encoded, string(c))
		}
	}
	return strings.Join(encoded, "")
}

func main() {
	var decode bool
	flag.BoolVar(&decode, "decode", false, "...")
	flag.Parse()

	if decode {
		fmt.Println(Decode(strings.Join(flag.Args(), " ")))
		return
	}

	fmt.Println(Encode(strings.Join(flag.Args(), " ")))
}
