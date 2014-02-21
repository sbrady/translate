package piglatin

import (
	"testing"
	"fmt"
)

func TestTheLengthIsCorrect(t *testing.T){
	translation := Translate("happy")
	assert (len(translation) == 7, "the length is wrong", t)
}

func TestEndsInAy(t *testing.T) {
	translation := Translate("happy")
	lastTwo := translation[(len(translation) -2):len(translation)]
	assert (lastTwo == "ay", "ay is missing", t)
}

func TestItMovesTheConsonantToTheEnd(t  *testing.T) {
	translation := Translate("happy")
	the_consonant := translation[4:(len(translation) -2)]
	assert (the_consonant == "h", "the consonant is in the wrong place", t)
}

func TestItMovesTheConsonantClusterToTheEnd(t  *testing.T) {
	translation := Translate("glove")
	expected_consonant_cluster := translation[3:(len(translation) -2)]
	assert (expected_consonant_cluster == "gl", "the cluster is in the wrong place", t)
}

//Functional Tests

func TestWhenItBeginsWithAConsonant(t  *testing.T){
	translation := Translate("happy")
	assert (translation == "appyhay", "expected appyhay", t)
}

func TestWhenItBeginsWithAVowel(t  *testing.T){
	translation := Translate("inbox")
	fmt.Println(translation)

	assert (translation == "nboxiway", "expected nboxiway", t)
}

func TestWhenItBeginsWithAConstantCluster(t  *testing.T){
	translation := Translate("glow")
	assert (translation == "owglay", "expected appyhay owglay", t)
}

func TestWhenItHasADoubleLetterAfterTheFirstLetterAndTheFirstLetterIsConsonant(t  *testing.T){
	translation := Translate("hoop")
	assert (translation == "ophay", "expected ophay", t)
}

func TestWhenItHasADoubleLetterAfterTheFirstLetterIsAVowel(t  *testing.T){
	translation := Translate("egg")
	assert (translation == "geway", "expected geway", t)
}



//Helpers
func assert(cond bool, error_message string, t *testing.T) {
	if !cond {
		t.Error(error_message)
	}
}
