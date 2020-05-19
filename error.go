package main

import (
	"fmt"
	"log"
)

// Si une erreur a eu lieu, arrête immédiatement le progamme et affiche une
// erreur
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println(`Usage à 
	développer`)
}
