package main

import (
	"fmt"
)

func main() {
	var i string
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scan(&i)

	// Controle of het kenteken in de lijst voorkomt
	if i == "AB-123-CD" || i == "XY-456-ZZ" || i == "LM-789-OP" {
		// Als het i in de lijst staat
		fmt.Println("Goedemorgen! Welkom bij Fonteyn Vakantieparken!", i)
	} else {
		// Als het i niet in de lijst staat
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}
