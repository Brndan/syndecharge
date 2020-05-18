package main

import (
	"github.com/tealeg/xlsx"
)

var header = []string{"Syndicat", "Civilité", "Prénom", "Nom", "Heures décharge", "Min décharge", "Heures ORS", "Min ORS", "Corps", "RNE"}

// Une plage de cellules qui reçoit les coordonnées x,y
// pour le début et la fin de la plage
type cellRange struct {
	BeginX int
	EndX   int
	BeginY int
	EndY   int
}

// un type qui permet de passer aisément une référence vers un tableur en
// mémoire. Permet d'utiliser des méthodes pour ajouter des feuilles, des
// lignes, et des cellules
type xlsxObject struct {
	Row   *xlsx.Row
	Cell  *xlsx.Cell
	Sheet *xlsx.Sheet
}

// Prend en argument une cellule de début de plage (ex "A26")
// et une cellule de fin de plage (ex "JJ80") et retourne clr, un struct
// cellRange qui contient les coordonnées de début et de fin de plage sous
// forme d'un axe x, y
func rangeCoordinates(rangeStart, rangeEnd string) (clr cellRange) {
	var err error
	clr.BeginX, clr.BeginY, err = xlsx.GetCoordsFromCellIDString(rangeStart)
	checkErr(err)
	clr.EndX, clr.EndY, err = xlsx.GetCoordsFromCellIDString(rangeEnd)
	return
}

// Ouvre un fichier file. Extrait la plage de données cellRange, et l’ajoute à
// la suite d'un objet destXlsx déjà en mémoire
func extractRange(file string, clr cellRange, destXlsx xlsxObject) (err error) {
	fileSlice, err := xlsx.FileToSlice(file)
	checkErr(err)
	for row := clr.BeginY; row <= clr.EndY; row++ {
		if fileSlice[0][row][clr.BeginX+1] != "" {
			destXlsx.Row = destXlsx.Sheet.AddRow()
			for col := clr.BeginX; col <= clr.EndX; col++ {
				destXlsx.Cell = destXlsx.Row.AddCell()
				destXlsx.Cell.Value = fileSlice[0][row][col]
			}
		}
	}

	return
}
