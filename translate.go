package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/pivotal/translate/piglatin"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() == true {
		str := scanner.Text()
		fmt.Printf("%s -> %s\n", str, piglatin.Translate(str))
	}
}