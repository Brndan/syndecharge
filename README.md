# syndecharge

Compile les décharges individuelles des tableaux des syndicats.



## Utilisation

Usage : `syndecharge [--cts] [--csv] [--begin A25] [--end J44] [-i dossier] [-o fichier]`
		
*syndecharge* est un programme qui compile les données de plages de fichiers
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
			
	--csv	Au lieu d’exporter un fichier Excel, produit une sortie CSV
			sur la sortie standard
		
	-v		affiche la version et la date de compilation.
	
	-h		affiche cette aide.



*Exemples :*
	

	syndecharge --begin A25 --end J44 -i tableaux -o ../synthèse.xlsx

→ Génère un fichier syntèse.xlsx de la plage A25 à J44 de tous les fichiers
		situés dans le répertoire courant.
	
	syndecharge --cts --begin A74 --end A74

→ Génère un fichier export.xlsx dans le répertoire courant contenant la
synthèse de tous le temps utilisé par les syndicats à partir des fichiers
situés dans le répertoire courant.

	synddecharge --cts --begin A25 --end J44 > export.csv

→ Redirige la sortie vers le fichier export.csv. Le programme affiche sur la sortie d'erreurs une liste des fichiers traités pour vérification.

La redirection risque de poser des problèmes d’encodage sous Windows.

## Compilation

Le programme est codé en [Go](https://golang.org/dl/).

Pour construire le programme, lancer le script `build.sh` sous Linux, `build.ps1` sous Windows.