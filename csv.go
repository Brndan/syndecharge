package main

import (
	"fmt"
	"regexp"
	"github.com/Brndan/syndecharge/problem"

	"github.com/tealeg/xlsx"
)

// Ouvre un fichier file. Extrait la plage de données cellRange, et l’imprime
// sur la sortie standard
// TODO : error-handling
func printCSV(file string, clRange cellRange) (err error) {

	fileSlice, err := xlsx.FileToSlice(file)
	problem.CheckErr(err)
	for row := clRange.BeginY; row <= clRange.EndY; row++ {
		if fileSlice[0][row][clRange.BeginX+1] != "" {
			for col := clRange.BeginX; col <= clRange.EndX; col++ {
				if match, _ := regexp.MatchString(`^[-+]?[0-9]*\.?[0-9]+$`, fileSlice[0][row][col]); match {
					fmt.Printf("%s", fileSlice[0][row][col])
				} else {
					fmt.Printf("\"%s\"", fileSlice[0][row][col])
				}
				if col < clRange.EndX {
					fmt.Printf(",")
				}

			}
			fmt.Printf("\n")
		}
	}
	return
}
