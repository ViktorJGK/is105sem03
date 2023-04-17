package mycrypt

import "unicode"

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		if indeks == -1 {
			// Hvis symbol ikke finnes i alfabetet, kopier det direkte til den krypterte meldingen
			kryptertMelding[i] = melding[i]
		} else {
			if indeks+chiffer >= len(alphabet) {
				// Hvis indeksen overskrider antall symboler i alfabetet etter å ha lagt til chiffer,
				// bruk modulo-operatoren for å wrap rundt til starten av alfabetet
				kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
			} else {
				kryptertMelding[i] = alphabet[indeks+chiffer]
			}
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		// Bruk unicode.ToLower for å konvertere symbol og alfabet[i] til små bokstaver før sammenligning
		if unicode.ToLower(symbol) == unicode.ToLower(alfabet[i]) {
			return i
		}
	}
	// Returner -1 hvis symbol ikke finnes i alfabetet
	return -1
}
