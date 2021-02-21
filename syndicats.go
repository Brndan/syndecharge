package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/Brndan/syndecharge/problem"

	"github.com/tealeg/xlsx"
)

// sumSyndicats prend en entrée un struct de type commandlineFlag et produit un
// tableau contenant toutes les lignes d'une plage pour chaque fichier dans le
// dossier d'export. Le fichier est alors sauvé dans un xlsx.
func exportSyndicats(opts commandlineFlags) {
	var (
		err        error
		exportFile xlsxObject
	)

	clrange := rangeCoordinates(opts.rangeStart, opts.rangeEnd)
	if clrange.BeginX > clrange.EndX || clrange.BeginY > clrange.EndY {
		fmt.Fprintf(os.Stderr, "Vous avez spécifié un début de plage %s plus grand que la fin de plage %s.\nImpossible de continuer.\n", opts.rangeStart, opts.rangeEnd)
		os.Exit(1)
	}

	var headerBar []string

	switch {
	case opts.syndicatsFlag && !opts.ctsFlag:
		headerBar = header
	case !opts.syndicatsFlag && opts.ctsFlag:
		headerBar = headerSum
	}

	// Récupère le nom de chaque fichier à traiter
	folderPath, _ := filepath.Abs(opts.inputFiles)
	fileList, err := os.ReadDir(folderPath)
	problem.CheckErr(err)

	switch {
	case opts.csvFlag:
		for i, col := range headerBar {
			fmt.Printf("\"%s\"", col)
			if i < len(headerBar)-1 {
				fmt.Printf(",")
			}
		}
		fmt.Printf("\n")

		var fileFullPath string
		for _, file := range fileList {
			fileFullPath = filepath.Join(folderPath, file.Name())
			if checkIfXlsx, _ := path.Match("*.xlsx", file.Name()); checkIfXlsx {
				fmt.Fprintf(os.Stderr, "%s\n", file.Name())
				printCSV(fileFullPath, clrange)

			}
		}

	default:
		f := xlsx.NewFile()
		exportFile.Sheet, err = f.AddSheet("Feuille1")
		problem.CheckErr(err)
		// On crée la première ligne du tableau avec les en-têtes
		exportFile.Row = exportFile.Sheet.AddRow()
		for _, col := range headerBar {
			exportFile.Cell = exportFile.Row.AddCell()
			exportFile.Cell.Value = col

		}

		var fileFullPath string
		for _, file := range fileList {
			fileFullPath = filepath.Join(folderPath, file.Name())
			if checkIfXlsx, _ := path.Match("*.xlsx", file.Name()); checkIfXlsx {
				fmt.Fprintf(os.Stderr, "%s : ", file.Name())
				err := extractRange(fileFullPath, clrange, exportFile)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				} else {
					fmt.Fprintln(os.Stderr, "OK")
				}
			}
		}
		outputFullPath, _ := filepath.Abs(opts.outputFile)
		err = f.Save(outputFullPath)
		problem.CheckErr(err)
	}

	return
}
