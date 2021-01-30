package main

import (
	"fmt"
	"strings"
)

type Kumpulan map[string]string

func main() {
	var morse = Kumpulan{
		"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".", "F": "..-.",
		"G": "--.", "H": "....", "I": "..", "J": ".---", "K": "-.-", "L": ".-..", "M": "--",
		"N": "-.", "O": "---", "P": ".--.", "Q": "--.-", "R": ".-.", "S": "...", "T": "-",
		"U": "..-", "V": "...-", "W": ".--", "X": "-..-", "Y": "-.--", "Z": "--..", "": "/",
	}
	var input string
	fmt.Print("masukan latin: ")
	fmt.Scanf("%s", &input)
	var hasil []string
	translateMorseToString(morse, input, &hasil)
	fmt.Println("morse nya:", hasil)
}

func translateMorseToString(kode Kumpulan, kata string, tampung *[]string) {
	misah := strings.Split(kata, "")
	for _, isikata := range misah {
		for key, isi := range kode {
			if isikata == strings.ToLower(key) {
				*tampung = append(*tampung, isi)
			} else if isikata == strings.ToUpper(key) {
				*tampung = append(*tampung, isi)
			}
		}
	}
}
