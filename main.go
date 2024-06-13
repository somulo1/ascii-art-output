package main

import (
	"ascii-art-output/functions"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define flag that will be used to specify the output file
	output := flag.String("output", "output.txt", "File that stores the output.")
	flag.Parse()

	if len(flag.Args()) > 2 || len(flag.Args()) < 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	stringInput := flag.Args()[0]

	// Variable to track if the flag was set
	var nameSet bool
	var flagSet bool = false

	// Enforce the flag format to be used to be --output=<filename.txt>
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "output" {
			flagSet = true
			result := strings.Replace(os.Args[1], *output, "", 1)
			if !(result == "--output=") {
				nameSet = true
			}
		}
	})
	// defining usage and error handling
	if nameSet {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	if !(strings.HasSuffix(*output, ".txt")) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	BannerFile := "standard.txt"

	if len(flag.Args()) == 2 {
		banner := strings.Replace(flag.Args()[1], ".txt", "", 1)
		BannerFile = banner + ".txt"
	}

	// read banner file specified
	file, err := os.ReadFile(BannerFile)
	if err != nil {
		fmt.Println("Error openning", BannerFile, err)
		return
	}
	file = []byte(strings.Replace(string(file), "\r\n", "\n", -1))

	fileLine := strings.Split(string(file), "\n")

	link := ""
	switch BannerFile {
	case "standard.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt"
	case "shadow.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt"
	case "thinkertoy.txt":
		link = "https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt"
	}

	if len(fileLine) != 856 {
		fmt.Println("The file", BannerFile, "is not correctly formated, please use the correct version", link, "!!!")
		return
	}

	//Write the results to the output file specified by user then print the results.
	asciiOutput := functions.AsciiArt(stringInput, fileLine)
	error := os.WriteFile(*output, []byte(asciiOutput), 0644)
	if error != nil {
		fmt.Println("Error:", error)
	} else if flagSet == false {
		fmt.Print(functions.AsciiArt(stringInput, fileLine))
	}
}
