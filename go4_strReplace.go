package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	fmt.Println(broken) // output: G# r#cks!
	replacer := strings.NewReplacer("#", "o")
	broken2 := replacer.Replace(broken)
	fmt.Println(broken2) // output: Go rocks!

}
