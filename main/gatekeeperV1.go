package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour() // Haal alleen het uur op

	if hour >= 7 && hour < 12 {
		fmt.Println("Goedemorgen! Welkom bij Fonteyn Vakantieparken!")
	} else if hour >= 12 && hour < 18 {
		fmt.Println("Goedemiddag! Welkom bij Fonteyn Vakantieparken!")
	} else if hour >= 18 && hour < 23 {
		fmt.Println("Goedenavond! Welkom bij Fonteyn Vakantieparken!")
	} else {
		fmt.Println("Sorry, de parkeerplaats is gesloten")
	}
}
