package main

import(
	"fmt"
	"os"
	"strings"
)

func main() {
	var colorflag, substring,fulltext,colorname, colorcode string

	Argcount := len (os.Args)

	switch Argcount{
	case 2:
		fulltext = os.Args[1]
	case 3:
		colorflag = os.Args[1]
		fulltext = os.Args[2]
		substring = os.Args[2]
	case 4:
		colorflag = os.Args[1]
		fulltext = os.Args[3]
		substring = os.Args[2]
	default:
		fmt.Println("Usage: go run main.go <color> <substring> <fulltext>")
		return
	}

	if colorflag != "" {
		if strings.HasPrefix(colorflag, "--color=") {
			colorname = strings.TrimPrefix(colorflag, "--color=")
		}
	} else {
		fmt.Println("Usage: go run main.go <color> <substring> <fulltext>")
		return
	}
	

	reset := "\033[0m"

	switch colorname {
	case "red":
		colorcode = "\033[31m"
	case "green":
		colorcode = "\033[32m"
	case "yellow":
		colorcode = "\033[33m"
	case "blue":
		colorcode = "\033[34m"
	case "magenta":
		colorcode = "\033[35m"
	case "cyan":
		colorcode = "\033[36m"
	default:
		fmt.Println("Unsupported color. Supported colors: red, green, yellow, blue, magenta, cyan.")
		return
	}
	result := mapcolors(substring, fulltext)

	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)
	theSlice := strings.Split(text, "\n")

	
		for row := 0; row < 8; row++ {
			for i, char := range fulltext {
				indexcalc := (int(char) - 32) * 9 + 1 + row
				if result[i] {
				fmt.Print(colorcode + theSlice[indexcalc] + reset)
				} else {
					fmt.Print(theSlice[indexcalc])
					}
		}
		fmt.Println()
	
}
}

func mapcolors(substring string, fulltext string) []bool {
	isColored := make([]bool, len(fulltext))
	if substring == "" {
		return isColored
	}
	index := 0 
	for index < len(fulltext) {
		matchindex := strings.Index(fulltext[index:], substring)
		if matchindex == -1 {
			break
		}
		actualStart := index + matchindex
		for i := actualStart; i < actualStart + len(substring); i++ {
			isColored[i] = true
		}
		index = actualStart + 1

	}
	return isColored
}