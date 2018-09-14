package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = clipboard.WriteAll(wd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Copied current working directory to the clipboard: %s\n", wd))
}
