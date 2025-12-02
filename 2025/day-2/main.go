package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("day-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

