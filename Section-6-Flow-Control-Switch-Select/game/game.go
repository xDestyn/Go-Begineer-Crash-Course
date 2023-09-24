package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input in channels
	// Print to screen
	// Keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber += round
			g.RoundChan <- 0
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			g.DisplayChan <- ""
		}
	}
}

// ClearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	g.DisplayChan <- fmt.Sprintf(`
Welcome to Rock, Paper, Scissors!
---------------------------------
Game is played for three rounds, and best of three wins!
`)
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	g.DisplayChan <- fmt.Sprintf("Round %d", g.Round.RoundNumber)
	<-g.DisplayChan

	fmt.Print("Please enter rock, paper, or scissors ->")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose rock"
		<-g.DisplayChan
	case PAPER:
		g.DisplayChan <- "Computer chose paper"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer chose scissors"
		<-g.DisplayChan
	default:
		fmt.Println("Computer chose something else")
	}

	g.DisplayChan <- ""
	<-g.DisplayChan

	if playerValue == computerValue {
		g.DisplayChan <- "It's a tie!"
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
		default:
			g.DisplayChan <- "Invalid choice"
			<-g.DisplayChan
			return false
		}
	}

	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	g.DisplayChan <- ""
	<-g.DisplayChan

	g.DisplayChan <- fmt.Sprintf(`
	Final score
	-----------
	Player: %d/3, Computer %d/3:`, g.Round.PlayerScore, g.Round.ComputerScore,
	)
	<-g.DisplayChan

	//fmt.Println()
	// fmt.Println("Final score")
	// fmt.Println("-----------")
	// fmt.Printf("Player: %d/3, Computer %d/3:", g.Round.PlayerScore, g.Round.ComputerScore)
	// fmt.Println()

	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins the game!"
		<-g.DisplayChan
		// fmt.Println("Player wins the game!")
	} else {
		g.DisplayChan <- "Computer wins the game!"
		<-g.DisplayChan
		//fmt.Println("Computer wins the game!")
	}
}
