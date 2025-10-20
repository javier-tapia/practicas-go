package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := range 3 {
		round := PlayRound(i)
		switch i {
		case 0:
			fmt.Println("El jugador eligió PIEDRA")
		case 1:
			fmt.Println("El jugador eligió PAPEL")
		case 2:
			fmt.Println("El jugador eligió TIJERA")
		}

		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("ComputerChoice: %s\n", round.ComputerChoice)
		fmt.Printf("RoundResult: %s\n", round.RoundResult)
		fmt.Printf("ComputerChoiceInt: %d\n", round.ComputerChoiceInt)
		fmt.Printf("ComputerScore: %s\n", round.ComputerScore)
		fmt.Printf("PlayerScore: %s\n", round.PlayerScore)

		fmt.Println("\n=================================================================")
	}
}
