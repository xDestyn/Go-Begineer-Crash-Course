package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
	//PLAYERWINS   = 1
	//COMPUTERWINS = 2
	//DRAW         = 3
)

type Round struct {
	//Winner         int
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

/**
 * Created three slices of strings, each with exactly three entries.
 */
var winMessages = []string{
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket!",
}

var loseMessages = []string{
	"Too bad!",
	"Try again!",
	"This is just not your day!",
}

var drawMessages = []string{
	"Great minds think alike!",
	"Uh oh. Try again!",
	"Nobody wins, but you can try again!",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	// Generate a random number between 0 and 2, inclusive.
	// Which message to use depends on the value of the random number.
	messageInt := rand.Intn(3)
	// Declare a variable to hold the message.
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins!"
		message = loseMessages[messageInt]
	}

	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult

	return result
}
