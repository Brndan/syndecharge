package main

import "flag"

// Un type qui permet de faire la liste de toutes les options de la ligne de
// commande
type commandlineFlags struct {
	rangeStart     string // début de la plage de données ex : A24
	rangeEnd       string // fin de la plage de donnée, ex : J56
	outputFile     string // nom du fichier d'export
	inputFiles     string // dossier où se trouve les fichiers à traiter
	csvFlag        bool   // préférer un export CSV
	xlsxFlag       bool   // préférer un export XLSX
	versionFlag    bool   // afficher la version du programme
	ctsFlag        bool   // synthèse du CTS
	syndicatsFlag  bool   // synthèse des syndicats ou mandatés
	sommeFlag      bool   // somme de toutes les décharges par syndicat
	structureCoord string // coordonnées du nom du syndicat dans le fichier pour les CTS
	tempsCtsCoord  string // coordonnées du CTS utilisé par syndicat
}

// Récupérer toutes les options de la ligne de commande
func getOpts(opts *commandlineFlags) {
	flag.BoolVar(&opts.versionFlag, "v", false, "Print version info and exit.")
	flag.StringVar(&opts.rangeStart, "begin", opts.rangeStart, "Début de la plage")
	flag.StringVar(&opts.rangeEnd, "end", opts.rangeEnd, "Fin de la plage")
	flag.StringVar(&opts.outputFile, "o", opts.outputFile, "Fichier d’export au format Excel")
	flag.StringVar(&opts.inputFiles, "i", opts.inputFiles, "Tableurs en entrée ; peut être un dossier")
	flag.BoolVar(&opts.csvFlag, "csv", false, "Export CSV")
	flag.BoolVar(&opts.xlsxFlag, "xlsx", true, "Export XLSX (défaut)")
	flag.BoolVar(&opts.ctsFlag, "cts", false, "Export du CTS")
	flag.BoolVar(&opts.syndicatsFlag, "syndicats", true, "Exporter les décharges locales ou fédérales")
	flag.BoolVar(&opts.sommeFlag, "somme", false, "Exporter la synthèse du temps consommé par le syndicat")
	flag.StringVar(&opts.structureCoord, "structure", "", "Coordonnées du nom du syndicat dans les fichiers importés")
	flag.StringVar(&opts.tempsCtsCoord, "temps", "", "Coordonnées du CTS du syndicat dans les fichiers importés")

	flag.Parse()
	return
}
