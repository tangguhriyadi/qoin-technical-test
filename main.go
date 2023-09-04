package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id    int
	dice  []int
	score int
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var numPlayers, numDice int
	fmt.Print("Enter the number of players: ")
	fmt.Scan(&numPlayers)
	fmt.Print("Enter the number of dice: ")
	fmt.Scan(&numDice)

	players := initializePlayers(numPlayers, numDice)

	round := 1
	for len(players) > 1 {
		fmt.Printf("==================\n")
		fmt.Printf("Round %d: Rolling the dice...\n", round)
		rollDice(players, r)
		evaluateDice(players)
		fmt.Println("After evaluation:")
		for i := range players {
			for j := range players[i].dice {
				if players[i].dice[j] == 99999 {
					players[i].dice[j] = 1
				}
			}
			fmt.Printf("Player #%d (Score: %d): %v\n", players[i].id, players[i].score, players[i].dice)
		}
		removePlayersWithoutDice(&players)
		round++
	}

	winner := getWinner(players)

	fmt.Printf("==================\n")
	fmt.Println("The game ends because only Player #", players[0].id, "has dice left.")
	fmt.Println("The game is won by Player #", winner.id, "for having the most points among all players.")
}

func initializePlayers(numPlayers, numDice int) []Player {
	players := make([]Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = Player{id: i + 1, dice: make([]int, numDice)}
	}
	return players
}

func rollDice(players []Player, r *rand.Rand) {
	for i := range players {
		for j := range players[i].dice {
			players[i].dice[j] = r.Intn(7)
		}
		fmt.Printf("Player #%d (Score: %d): %v\n", players[i].id, players[i].score, players[i].dice)
	}
}

func evaluateDice(players []Player) {
	for i := range players {
		for j := 0; j < len(players[i].dice); j++ {
			if players[i].dice[j] == 6 {
				players[i].score++
				players[i].dice = append(players[i].dice[:j], players[i].dice[j+1:]...)
				j--
			} else if players[i].dice[j] == 1 {
				nextPlayerIndex := (i + 1) % len(players)
				players[nextPlayerIndex].dice = append(players[nextPlayerIndex].dice, 99999)
				players[i].dice = append(players[i].dice[:j], players[i].dice[j+1:]...)
				j--
			}
		}

	}
}

func removePlayersWithoutDice(players *[]Player) {
	newPlayers := []Player{}
	for _, player := range *players {
		if len(player.dice) > 0 {
			newPlayers = append(newPlayers, player)
		}
	}
	*players = newPlayers

}

func getWinner(players []Player) Player {
	winner := players[0]
	for _, player := range players {
		if player.score > winner.score {
			winner = player
		}
	}
	return winner
}
