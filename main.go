package main

import (
	"ascii-art-fs/functions"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 4 || len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	var stringInput string
	if len(os.Args) > 2 {
		stringInput = os.Args[2] // Reading the string argument entered
	} else if len(os.Args) == 2 {
		stringInput = os.Args[1]
	}

	if stringInput == "" {
		return
	}

	BannerFile := "standard.txt"

	if len(os.Args) == 4 {
		banner := strings.Replace(os.Args[3], ".txt", "", 1)
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

	//Define flag that will be used to specify the output file
	output := flag.String("output", "output.txt", "File that stores the output.")
	flag.Parse()

	//Enforce the specified double dash format
	for i, arg := range os.Args {
		if i == 1 {
			if strings.HasPrefix(arg, "-output") || strings.HasPrefix(arg, "--output= ") || strings.HasPrefix(arg, "--output") && !strings.Contains(arg, "--output=") ||
				strings.HasPrefix(arg, "-output") && !strings.HasSuffix(arg, ".txt") || strings.HasPrefix(arg, "--output") && !strings.HasSuffix(arg, ".txt") {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
		}
	}
	asciiOutput := functions.AsciiArt(stringInput, fileLine)
	error := os.WriteFile(*output, []byte(asciiOutput), 0644)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Print(functions.AsciiArt(stringInput, fileLine))
	}

}
