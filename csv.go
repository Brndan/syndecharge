package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func printCSV(file string, clRange cellRange) (err error) {

	fileSlice, err := xlsx.FileToSlice(file)
	checkErr(err)
	for row := clRange.BeginY; row <= clRange.EndY; row++ {
		if fileSlice[0][row][clRange.BeginX+1] != "" {
			for col := clRange.BeginX; col <= clRange.EndX; col++ {
				fmt.Printf("\"%s\"", fileSlice[0][row][col])
				if col < clRange.EndX {
					fmt.Printf(",")
				}

			}
			fmt.Printf("\n")
		}
	}
	return
}
