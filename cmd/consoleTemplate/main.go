package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "restTemplate", log.LstdFlags)
	logger.Println("Starting console application...")

	// Int flag example
	var nFlag *int
	nFlag = flag.Int("n", 1234, "help message for flag n")

	// String flag example
	var strFlag *string
	strFlag = flag.String("str", "default", "help message for flag str")

	// Boolean flag example
	var boolFlag *bool
	boolFlag = flag.Bool("bool", false, "help message for flag bool")

	// Parse flags
	flag.Parse()

	logger.Printf("Flag n: %d\n", *nFlag)
	logger.Printf("Flag str: %s\n", *strFlag)
	logger.Printf("Flag bool: %t\n", *boolFlag)
}
