package piglatin

func Translate(english string) (string) {
	englishRunes := []rune(english)
	return translateToPig(englishRunes)
}

func translateToPig(english []rune) (string) {
	return addCorrectLatinSoundBasedOnLastLetter(moveLettersToEnd(removeDuplicateLettersFromBeginningAfterFirstLetter(english)))
}

func addCorrectLatinSoundBasedOnLastLetter(english []rune) (string){
	if isVowel(english[len(english) - 1]) {
		return string(english) + "way"
	}
	return string(english) + "ay"
}

func moveLettersToEnd(english []rune) ([]rune) {
	if isVowel(english[0]) { 
		f := english[1:len(english)]
		f = append(f, english[0])
		return f
	}

	return moveConsonantClusterToEnd(english, extractBeginningConsonantCluster(english))
}

func moveConsonantClusterToEnd(english []rune, consonantsCluster []rune)([]rune){
	return append(englishWithOutConsonantCluster(english, consonantsCluster), consonantsCluster...)
}

func removeDuplicateLettersFromBeginningAfterFirstLetter(english []rune)([]rune){
	if (english[1] == english[2]) {
		return append(english[0:1], english[2:len(english)]...)
	}
	return english
}

func englishWithOutConsonantCluster(english []rune, consonants []rune) ([]rune) {
	return english[len(consonants):len(english)]
}

func extractBeginningConsonantCluster(word []rune) ([]rune) {
	consonants := []rune{}
	for _, letter := range word {
		if isConsonant(letter) {
			consonants = append(consonants, letter)
		} else {
			break;
		}
	}
	return consonants
}

func isVowel(letter rune) (bool) {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}

func isConsonant(letter rune) (bool) {
	return !isVowel(letter)
}