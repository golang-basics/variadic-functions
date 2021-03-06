package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first ")
	r2 := strings.NewReader("second ")
	r3 := strings.NewReader("third\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
