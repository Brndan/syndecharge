package main

import (
	"fmt"
	"os"
)

func main() {

	cwd, err := os.Getwd()
	checkErr(err)

	var options commandlineFlags

	// valeurs par défaut du début de plage et de fin pour les décharges locales
	options.rangeStart, options.rangeEnd = "A25", "J44"
	options.outputFile = "export.xlsx"
	options.inputFiles = cwd
	// récupération des paramètres
	getOpts(&options)

	switch {
	case options.versionFlag:
		fmt.Printf("Date de compilation : %s\nIdentifiant de version : %s", buildTime, sha1ver)
		os.Exit(0)
	case (options.syndicatsFlag || options.ctsFlag) && !(options.syndicatsFlag && options.ctsFlag):
		sumSyndicats(options)
	default:
		fmt.Printf("Vous ne devez utiliser qu'une option parmi -cts et -syndicats.\nSi vous n’entrez aucune de ces trois options -syndicats est utilisé par défaut.")
		usage()
		os.Exit(1)
	}

}
