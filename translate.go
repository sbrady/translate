package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/sbrady/translate/translator"
	"github.com/sbrady/translate/piglatin"
	"github.com/sbrady/translate/russian"
)

func newTranslator(translatorType string) translator.Translator {
	switch translatorType {
		case "piglatin":
			return &piglatin.Translator{}
		case "russian":
			return &russian.Translator{}
	}

	panic("missing translator type")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	translator := newTranslator(os.Args[1])
	
	for scanner.Scan() == true {
		str := scanner.Text()
		fmt.Printf("%s -> %s\n", str, translator.Translate(str))
	}
}