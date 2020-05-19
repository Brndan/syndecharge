package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/tealeg/xlsx"
)

func main() {

	cwd, err := os.Getwd()
	checkErr(err)

	var options commandlineFlags

	// valeurs par défaut du début de plage et de fin pour les décharges locales
	options.rangeStart, options.rangeEnd = "A25", "J43"
	options.outputFile = "export.xlsx"
	options.inputFiles = cwd
	// récupération des paramètres
	getOpts(&options)

	if options.versionFlag == true {
		fmt.Printf("Date de compilation : %s\nIdentifiant de version : %s", buildTime, sha1ver)
		os.Exit(0)
	}

	if options.csvFlag == true {
		options.xlsxFlag = false
	}

	clrange := rangeCoordinates(options.rangeStart, options.rangeEnd)

	f := xlsx.NewFile()
	var exportFile xlsxObject

	exportFile.Sheet, err = f.AddSheet("Feuille1")
	checkErr(err)

	exportFile.Row = exportFile.Sheet.AddRow()
	for _, col := range header {
		exportFile.Cell = exportFile.Row.AddCell()
		exportFile.Cell.Value = col

	}

	// Récupère le nom de chaque fichier à traiter
	fileList, err := ioutil.ReadDir(options.inputFiles)

	checkErr(err)
	for _, file := range fileList {
		if checkIfXlsx, _ := path.Match("*.xlsx", file.Name()); checkIfXlsx == true {
			switch {
			case options.csvFlag == true:
				printCSV(file.Name(), clrange)
			default:
				extractRange(file.Name(), clrange, exportFile)
			}
		}
	}

	if options.xlsxFlag == true {
		err = f.Save(options.outputFile)
		checkErr(err)
	}

}
