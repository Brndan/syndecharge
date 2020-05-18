package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"

	"github.com/tealeg/xlsx"
)

func main() {

	cwd, err := os.Getwd()
	checkErr(err)

	// valeurs par défaut du début de plage et de fin pour les décharges locales
	rangeStart, rangeEnd := "A25", "J43"
	outputFile := "export.xlsx"
	inputFiles := cwd
	csvFlag := false
	xlsxFlag := false

	flag.StringVar(&rangeStart, "d", rangeStart, "Début de la plage")
	flag.StringVar(&rangeEnd, "f", rangeEnd, "Fin de la plage")
	flag.StringVar(&outputFile, "o", outputFile, "Fichier d’export au format Excel")
	flag.StringVar(&inputFiles, "i", cwd, "Tableurs en entrée ; peut être un dossier")
	flag.BoolVar(&csvFlag, "csv", false, "Export CSV")
	flag.BoolVar(&xlsxFlag, "xlsx", true, "Export XLSX (défaut)")

	flag.Parse()

	if csvFlag == true {
		xlsxFlag = false
	}
	clrange := rangeCoordinates(rangeStart, rangeEnd)

	f := xlsx.NewFile()
	var exportFile xlsxObject

	exportFile.Sheet, err = f.AddSheet("Feuille1")
	checkErr(err)

	//header := []string{"Syndicat", "Civilité", "Prénom", "Nom", "Heures décharge", "Min décharge", "Heures ORS", "Min ORS", "Corps", "RNE"}

	exportFile.Row = exportFile.Sheet.AddRow()
	for _, col := range header {
		exportFile.Cell = exportFile.Row.AddCell()
		exportFile.Cell.Value = col

	}

	// Affiche le nom de chaque fichier à traiter
	fileList, err := ioutil.ReadDir(inputFiles)

	checkErr(err)
	for _, file := range fileList {
		if checkIfXlsx, _ := path.Match("*.xlsx", file.Name()); checkIfXlsx == true {
			switch {
			case csvFlag == true:
				printCSV(file.Name(), clrange)
			default:
				extractRange(file.Name(), clrange, exportFile)
			}
		}
	}

	if xlsxFlag == true {
		err = f.Save(outputFile)
		checkErr(err)
	}

}
