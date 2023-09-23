package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	argsLength := len(args)

	// if having correct arguments (number of files)
	if argsLength == 2 {
		var arg interface{} = os.Args[1]

		// check if the argument is string
		input, isString := arg.(string)
		if !isString {
			fmt.Printf("Please input valid string as the argument.")
			os.Exit(1)
		}

		// if the string is only "" then return nothing
		if os.Args[1] == "" {
			return
		}

		// detect escape on command line as escape
		input = strings.ReplaceAll(args[1], "\\n", "\n")

		// read the text file and split each ascii-art to slice
		asciibyte, _ := os.ReadFile("standard.txt")
		artSlice := strings.Split(string(asciibyte), "\n\n")
		resultByNL := strings.Split(input, "\n")

		for _, lc := range resultByNL {

			// deal with "" (more than double \n)
			if lc == "" {
				fmt.Println()
				continue
			}

			// look for and store corresponding word to the result
			artList := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
			resultToPrint := []string{}
			for _, ib := range lc {
				for ai, ab := range artList {
					if ib == ab {
						resultToPrint = append(resultToPrint, artSlice[ai])
					}
				}
			}

			// prepare the string slice's slice to store result by lines
			wordNumber := len(resultToPrint)
			MAXLINES := 8
			ResultByLines := make([][]string, wordNumber)
			for i := 0; i < wordNumber; i++ {
				ResultByLines[i] = make([]string, MAXLINES)
			}

			// split art by lines
			for h := 0; h < wordNumber; h++ {
				resultSplit := strings.Split(resultToPrint[h], "\n")
				for ln := 0; ln < MAXLINES; ln++ {
					ResultByLines[h][ln] = resultSplit[ln]
				}
			}

			// print result
			for ln := 0; ln < MAXLINES; ln++ {
				for wn := 0; wn < wordNumber; wn++ {
					if ln == 0 && ResultByLines[wn][ln] == "" {
						fmt.Print("      ")
						continue
					}
					fmt.Print(ResultByLines[wn][ln])
				}
				fmt.Println()
			}
		}

	} else { // not enough or too many arguments
		fmt.Println("Please input a string as the argument.")
	}
}
