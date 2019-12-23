package translit

// RuneToUA returns Ukrainian equivalent of Latin character.
func RuneToUA(char rune) rune {
	switch char {
	case 'A':
		return 'А'
	case 'B':
		return 'В'
	case 'C':
		return 'С'
	case 'E':
		return 'Е'
	case 'H':
		return 'Н'
	case 'I':
		return 'І'
	case 'K':
		return 'К'
	case 'M':
		return 'М'
	case 'O':
		return 'О'
	case 'P':
		return 'Р'
	case 'T':
		return 'Т'
	case 'X':
		return 'Х'
	default:
		return char
	}
}

// RuneToUA returns Ukrainian equivalent of Latin character.
func RuneToLatin(char rune) rune {
	switch char {
	case 'А':
		return 'A'
	case 'В':
		return 'B'
	case 'С':
		return 'C'
	case 'Е':
		return 'E'
	case 'Н':
		return 'H'
	case 'І':
		return 'I'
	case 'К':
		return 'K'
	case 'М':
		return 'M'
	case 'О':
		return 'O'
	case 'Р':
		return 'P'
	case 'Т':
		return 'T'
	case 'Х':
		return 'X'
	default:
		return char
	}
}

// ToUA translates string from latin to cyrillic equivalent.
func ToUA(lexeme string) string {
	chars := make([]rune, 0)

	for _, v := range lexeme {
		chars = append(chars, RuneToUA(v))
	}

	return string(chars)
}

// ToLatin( translates string from cyrillic to latin equivalent.
func ToLatin(lexeme string) string {
	chars := make([]rune, 0)

	for _, v := range lexeme {
		chars = append(chars, RuneToLatin(v))
	}

	return string(chars)
}
