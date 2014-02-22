package russian

import (
	"testing"
	// "fmt"
)

func TestItTranslateAnEnglishWordToRussian(t  *testing.T){
	translation := translator.Translate("egg")
	assert (translation == "яйцо", "expected яйцо", t)
}

//Helpers
func assert(cond bool, error_message string, t *testing.T) {
	if !cond {
		t.Error(error_message)
	}
}