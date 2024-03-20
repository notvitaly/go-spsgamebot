package game

import (
	"testing"
)

func TestStringToArrays(t *testing.T) {
	var strs []string = []string{"123456", "1234", "12", "", "1", "123", "12345"}

	for _, str := range strs {
		arrays := stringToArrays(str)
		strFromArrays := arraysToString(arrays)

		if str != strFromArrays {
			t.Fail()
		}
	}
}

func TestParams(t *testing.T) {
	paramsStr := "4231|3|Vitaly-Ivan|123456|3"

	game := DecodeString(paramsStr)
	str := game.EncodeGame()

	if paramsStr != str {
		t.Fail()
	}
}

func TestLimitString(t *testing.T) {
	strLong := "Hello How Are you?"

	if strLong != limitString(strLong, len(strLong)) {
		t.Errorf("%s -> %s", strLong, limitString(strLong, len(strLong)))
	}
	if limitString(strLong, 5) != "Hello…" {
		t.Errorf("%s -> %s", strLong, limitString(strLong, 5))
	}
}

func TestMakeNewTurn(t *testing.T) {
	var player1id int64 = 423215995
	var player2id int64 = 112344543
	player1Name := "Vitaly"
	player2Name := "Ivan"

	str := "-1|1|-||0"
	g1 := DecodeString(str)

	g1.MakeNewTurn(player1id, player1Name, 0)

	g2 := DecodeString(g1.EncodeGame())
	g2.MakeNewTurn(player2id, player2Name, 1)
}

func TestFindWinner(t *testing.T) {
	// "🗿", "✂️", "🧻",
	// game := [][]int {
	// 	{0, 2},
	// 	{1, 0},
	// 	{2, 2},
	// }
	paramsstr := "4321|3|Vitaly-Ivan|021022|1"
	game := DecodeString(paramsstr)
	game.Finished = true

	// log.Default().Println(IsPlayerWinner(0, 2))
	// log.Default().Println(IsPlayerWinner(1, 0))
	// log.Default().Println(IsPlayerWinner(2, 2))

	// log.Default().Println(IsPlayerWinner(2, 0))
	// log.Default().Println(IsPlayerWinner(0, 1))
	// log.Default().Println(IsPlayerWinner(2, 2))
}
