package main

import (
	"fmt"
	"time"
)

type DeterministicDice struct {
	state int
}

func (d *DeterministicDice) Roll() int {
	result := d.state
	d.state = mod1(d.state+1, 100)
	return result
}

type Player struct {
	position int
	distance int
	score    int
}

func NewPlayer(position int) Player {
	return Player{position, 0, 0}
}

type Game struct {
	dice    DeterministicDice
	players []Player
	turn    int
}

func NewGame(players []Player) Game {
	return Game{DeterministicDice{1}, players, 0}
}

func (g *Game) RollDice() int {
	result := 0
	for i := 0; i < 3; i++ {
		result += g.dice.Roll()
	}
	return result
}

func (g *Game) Move(playerId int, distance int) {
	position := g.players[playerId].position
	position = mod1(position+distance, 10)

	g.players[playerId].distance += distance
	g.players[playerId].position = position
	g.players[playerId].score += position
}

func (g *Game) Turn() int {
	playerId := g.turn % len(g.players)

	distance := g.RollDice()
	g.Move(playerId, distance)
	// fmt.Printf("Player %d rolls %d and moves to space %d for a total score of %d\n", playerId+1, distance, g.players[playerId].position, g.players[playerId].score)
	g.turn++
	return g.players[playerId].score
}

func (g *Game) Play() int {
	for {
		score := g.Turn()
		if score >= 1000 {
			break
		}
	}

	playerId := g.turn % len(g.players)
	playerScore := g.players[playerId].score

	return playerScore * g.turn * 3
}

func mod1(x, m int) int {
	return (x-1)%m + 1
}

func main() {
	example := NewGame([]Player{NewPlayer(4), NewPlayer(8)})
	fmt.Printf("Example: %v\n", example.Play())

	game := NewGame([]Player{NewPlayer(6), NewPlayer(10)})
	result1 := game.Play()
	fmt.Printf("Puzzle 1: %v\n", result1)

	fmt.Println("========")
	start := time.Now()
	total := 0
	for i := 0; i < 444356092776315; i++ {
		// total++
	}
	fmt.Println(total)
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)

	// result2 :=
	// fmt.Printf("Puzzle 2: %v\n", result2)
}
