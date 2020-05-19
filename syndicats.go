package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/tealeg/xlsx"
)

// sumSyndicats prend en entrée un struct de type commandlineFlag et produit un
// tableau contenant toutes les lignes d'une plage pour chaque fichier dans le
// dossier d'export. Le fichier est alors sauvé dans un xlsx.
func sumSyndicats(opts commandlineFlags) {
	var (
		err        error
		exportFile xlsxObject
	)

	clrange := rangeCoordinates(opts.rangeStart, opts.rangeEnd)

	f := xlsx.NewFile()

	exportFile.Sheet, err = f.AddSheet("Feuille1")
	checkErr(err)

	// On crée la première ligne du tableau avec les en-têtes
	exportFile.Row = exportFile.Sheet.AddRow()
	var headerBar []string

	switch {
	case opts.syndicatsFlag && !opts.ctsFlag:
		headerBar = header
	case !opts.syndicatsFlag && opts.ctsFlag:
		headerBar = headerSum
	}
	for _, col := range headerBar {
		exportFile.Cell = exportFile.Row.AddCell()
		exportFile.Cell.Value = col

	}

	// Récupère le nom de chaque fichier à traiter
	fileList, err := ioutil.ReadDir(opts.inputFiles)
	checkErr(err)

	for _, file := range fileList {
		if checkIfXlsx, _ := path.Match("*.xlsx", file.Name()); checkIfXlsx {
			switch {
			case opts.csvFlag:
				printCSV(file.Name(), clrange)
			default:
				fmt.Fprintf(os.Stderr, "%s\n", file.Name())
				extractRange(file.Name(), clrange, exportFile)
			}
		}
	}

	err = f.Save(opts.outputFile)

	return
}
