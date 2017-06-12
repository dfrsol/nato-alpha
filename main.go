package main

import (
	"fmt"
	"strings"
)

var natoAlphabet = map[string]string {
	"a": "Alfa",
	"b": "Bravo",
	"c": "Charlie",
	"d": "Delta",
	"e": "Echo",
	"f": "Foxtrot",
	"g": "Golf",
	"h": "Hotel",
	"i": "India",
	"j": "Juliett",
	"k": "Kilo",
	"l": "Lima",
	"m": "Mike",
	"n": "November",
	"o": "Oscar",
	"p": "Papa",
	"q": "Quebec",
	"r": "Romeo",
	"s": "Sierra",
	"t": "Tango",
	"u": "Uniform",
	"v": "Victor",
	"w": "Whiskey",
	"x": "X-ray",
	"y": "Yankee",
	"z": "Zulu",
}

func main() {
	var word string

	fmt.Println("Enter a word you'd like translated to the Nato Alphabet")
	fmt.Scanln(&word)

	printNatoAlpha(word)
}

func printNatoAlpha(s string) {
	byteSlice := []byte(strings.ToLower(s))

	for _, letter := range byteSlice  {
		fmt.Println(natoAlphabet[string(letter)])
	}
}
