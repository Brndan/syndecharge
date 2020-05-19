package main

import (
	"flag"
	"fmt"
)

// Un type qui permet de faire la liste de toutes les options de la ligne de
// commande
type commandlineFlags struct {
	rangeStart    string // début de la plage de données ex : A24
	rangeEnd      string // fin de la plage de donnée, ex : J56
	outputFile    string // nom du fichier d'export
	inputFiles    string // dossier où se trouve les fichiers à traiter
	csvFlag       bool   // préférer un export CSV
	xlsxFlag      bool   // préférer un export XLSX
	versionFlag   bool   // afficher la version du programme
	ctsFlag       bool   // synthèse du CTS
	syndicatsFlag bool   // synthèse des syndicats ou mandatés
}

// Récupérer toutes les options de la ligne de commande
func getOpts(opts *commandlineFlags) {

	flag.Usage = func() {
		fmt.Println(`Usage : syndecharge [--cts] [--begin A25] [--end J44] [-i dossier] [-o fichier]
		
syndecharge est un programme qui compile les données de plages de fichiers
Excel contenant les déclarations de décharge des syndicats.
	
	--cts	crée un fichier Excel en sortie contenant les synthèses des syndicats :
	Syndicat | ETP attribué au syndicat | Mutualisation | "ETP disponibles | Consommé | Crédit d'Heures (CHS)
			Si --cts n'est pas renseigné, le défaut est une compilation des
			mandats des décharges locales ou fédérales.
		
	--begin	début de la plage de données. Par défaut, la valeur A25 est
			attribuée à ce paramètre.
		
	--end	fin de la plage de données. Par défaut, la valeur J44 est
			attribuée à ce paramètre.
		
	-i		dossier dans lequel se trouvent les fichiers à compiler. Si ce
			paramètre est omis, le répertoire courant est utilisé.
		
	-o		fichier en sortie. Par défaut, le fichier "export.xlsx" est
			généré dans le répertoire courant.
		
	-v		affiche la version et la date de compilation.
	
	-h		affiche cette aide.
	
Exemples :
	
	syndecharge --begin A25 --end J44 -o ../synthèse.xlsx
	
→ Génère un fichier syntèse.xlsx de la plage A25 à J44 de tous les fichiers
		situés dans le répertoire courant.
	
	syndecharge --cts --begin A74 --end A74
	
→ Génère un fichier export.xlsx dans le répertoire courant contenant la
synthèse de tous le temps utilisé par les syndicats à partir des fichiers
situés dans le répertoire courant.
		`)
	}

	flag.BoolVar(&opts.versionFlag, "v", false, "Print version info and exit.")
	flag.StringVar(&opts.rangeStart, "begin", opts.rangeStart, "Début de la plage")
	flag.StringVar(&opts.rangeEnd, "end", opts.rangeEnd, "Fin de la plage")
	flag.StringVar(&opts.outputFile, "o", opts.outputFile, "Fichier d’export au format Excel")
	flag.StringVar(&opts.inputFiles, "i", opts.inputFiles, "Tableurs en entrée ; peut être un dossier")
	flag.BoolVar(&opts.csvFlag, "csv", false, "Export CSV")
	flag.BoolVar(&opts.xlsxFlag, "xlsx", true, "Export XLSX (défaut)")
	flag.BoolVar(&opts.ctsFlag, "cts", false, "Export du CTS")
	flag.BoolVar(&opts.syndicatsFlag, "syndicats", true, "Exporter les décharges locales ou fédérales")

	flag.Parse()

	if opts.ctsFlag {
		opts.syndicatsFlag = false
	}
	return
}
