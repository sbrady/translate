package piglatin

func Translate(english string) (string) {
	return translateToPig(english)
}

func translateToPig(english string) (string) {
	return addCorrectLatinSoundBasedOnLastLetter(moveLettersToEnd(removeDuplicateLettersFromBeginningAfterFirstLetter(english)))
}

func addCorrectLatinSoundBasedOnLastLetter(english string) (string){
	if isVowel(rune(english[len(english)-1])) {
		return english + "way"
	}
	return english + "ay"
}

func moveLettersToEnd(english string) (string) {
	if isVowel(rune(english[0])){
		return string(english[1:len(english)]) + string(english[0])
	}

	return moveConsonantClusterToEnd(english, extractBeginningConsonantCluster(english))
}

func moveConsonantClusterToEnd(english string, consonantsCluster string)(string){
	return englishWithOutConsonantCluster(english, consonantsCluster) + consonantsCluster
}

func removeDuplicateLettersFromBeginningAfterFirstLetter(english string)(string){
	if (english[1] == english[2]) {
		return english[0:1] + english[2:len(english)]
	}
	return english
}

func englishWithOutConsonantCluster(english string, consonants string) (string) {
	return english[len(consonants):len(english)]
}

func extractBeginningConsonantCluster(word string) (string) {
	consonants := []rune{}
	for _, letter := range word {
		if isConsonant(letter) {
			consonants = append(consonants, letter)
		} else {
			break;
		}
	}
	return string(consonants)
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