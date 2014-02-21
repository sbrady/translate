package main

import (
	"fmt"
	"os"
	"github.com/pivotal/translate/piglatin"
)

func main() {
	fmt.Printf("%s", piglatin.Translate(os.Args[1]))
}