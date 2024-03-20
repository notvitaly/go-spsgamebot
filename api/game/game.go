package game

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Game struct {
	InitiatorID   int64
	playerOneName string
	playerTwoName string
	RoundsTotal   int
	rounds        [][]int
	Turn          int
	Finished      bool
}

func (g *Game) FindWinner() int {
	var playerOneWins int = 0
	var playerTwoWins int = 0

	for _, round := range g.rounds {
		if IsPlayerWinner(round[0], round[1]) {
			playerOneWins++
		}
		if IsPlayerWinner(round[1], round[0]) {
			playerTwoWins++
		}
	}

	if playerOneWins > playerTwoWins {
		return 1
	} else if playerTwoWins > playerOneWins {
		return 2
	} else {
		return 0
	}
}

func IsPlayerWinner(playerOneTurn int, playerTwoTurn int) bool {
	return (playerOneTurn+1)%3 == playerTwoTurn
}

func limitString(str string, length int) string {
	lastIdx := 0

	if len(str) > length {
		lastIdx = length
	} else {
		lastIdx = len(str)
	}

	var limited string = str[0:lastIdx]

	if len(str) > length {
		limited += ".."
	}

	return limited
}

func stringToArrays(input string) [][]int {
	var arrays [][]int

	for i := 0; i < len(input); i += 2 {
		num1 := int(input[i] - '0')
		arrayToAppend := []int{num1}
		var num2 int
		if i+1 < len(input) {
			num2 = int(input[i+1] - '0')
			arrayToAppend = append(arrayToAppend, num2)
		}
		arrays = append(arrays, arrayToAppend)
	}

	return arrays
}

func arraysToString(input [][]int) string {
	var roundsString string

	for i := range input {
		for _, turn := range input[i] {
			roundsString += strconv.Itoa(turn)
		}
	}
	return roundsString
}

func (g *Game) ToString() string {
	roundsCount := len(g.rounds)
	var str string

	if len(g.rounds[roundsCount-1]) == 2 && !g.Finished {
		roundsCount++
	}

	if g.Finished {
		str = "<b><i>The game is over</i></b>\n"
	} else {
		str = fmt.Sprintf("<b>Round %d/%d</b>\n", roundsCount, g.RoundsTotal)
	}

	turns := []string{
		"🗿", "✂️", "🧻",
	}

	for _, round := range g.rounds {
		if len(round) == 1 && g.playerTwoName == "" {
			str += fmt.Sprintf("%s 🤫 – 💭", g.playerOneName)
			break
		} else if len(round) == 1 {
			str += fmt.Sprintf("%s 🤫 – %s 💭", g.playerOneName, g.playerTwoName)
			break
		}

		str += fmt.Sprintf("%s %s – %s  %s\n", g.playerOneName, turns[round[0]], turns[round[1]], g.playerTwoName)
	}

	if g.Finished {
		winner := g.FindWinner()

		log.Default().Print("found winner: ")
		log.Default().Println(winner)

		if winner == 0 {
			str += fmt.Sprintf("\n<b><i>%s 🤝 %s</i></b>", g.playerOneName, g.playerTwoName)
		} else {
			players := []string{g.playerOneName, g.playerTwoName}
			str += fmt.Sprintf("\n<b><i>%s won!</i></b>", players[winner-1])
		}
	} else if len(g.rounds[len(g.rounds)-1]) == 2 {
		str += fmt.Sprintf("%s 💭", g.playerOneName)
	}

	return str
}

func DecodeString(encodedString string) Game {
	var game Game
	parts := strings.Split(encodedString, "|")

	game.InitiatorID, _ = strconv.ParseInt(parts[0], 10, 64)
	game.RoundsTotal, _ = strconv.Atoi(parts[1])
	game.playerOneName = strings.Split(parts[2], "-")[0]
	game.playerTwoName = strings.Split(parts[2], "-")[1]

	if parts[3] == "" {
		game.rounds = [][]int{}
	} else {
		game.rounds = stringToArrays(parts[3])
	}

	game.Turn, _ = strconv.Atoi(parts[4])

	return game
}

func (g *Game) EncodeGame() string {
	var roundsString []string

	for i := range g.rounds {
		for _, round := range g.rounds[i] {
			roundsString = append(roundsString, strconv.Itoa(round))
		}
	}

	str := fmt.Sprintf("%d|%d|%s-%s|%s|%d", g.InitiatorID, g.RoundsTotal, g.playerOneName, g.playerTwoName, strings.Join(roundsString, ""), g.Turn)
	return str
}

func (g *Game) MakeNewTurn(id int64, name string, turn int) bool {
	name = limitString(name, 14)
	g.Turn = turn

	if g.InitiatorID == -1 {
		g.InitiatorID = id
		g.playerOneName = name

		g.rounds = append(g.rounds, []int{turn})

		return true
	}

	if id == g.InitiatorID && len(g.rounds[len(g.rounds)-1]) == 2 {
		g.playerOneName = name
		g.rounds = append(g.rounds, []int{turn})

		return true
	} else if id != g.InitiatorID && len(g.rounds[len(g.rounds)-1]) == 1 {
		g.playerTwoName = name
		g.rounds[len(g.rounds)-1] = append(g.rounds[len(g.rounds)-1], turn)

		g.Finished = len(g.rounds) == g.RoundsTotal && len(g.rounds[len(g.rounds)-1]) == 2

		return true
	}

	return false
}

func Handler(w http.ResponseWriter, r *http.Request) {
	return
}
