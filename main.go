package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("invalid args")
		os.Exit(1)
	}
	args := os.Args[1:]

	m1, m2 := ConvertStringToMove(args[0]), ConvertStringToMove(args[1])
	res, err := RockPaperScissors(Move(m1), Move(m2))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(res)
}

type Move int

const (
	_ Move = iota
	Rock
	Paper
	Scissors
)

func ConvertStringToMove(s string) Move {
	switch {
	case strings.EqualFold(s, Rock.String()):
		return Rock
	case strings.EqualFold(s, Paper.String()):
		return Paper
	case strings.EqualFold(s, Scissors.String()):
		return Scissors
	default:
		return 0
	}
}

func (m Move) String() string {
	return []string{"", "Rock", "Paper", "Scissors"}[m]
}

func (m Move) IsValid() bool {
	return m > 0 && m <= 3
}

func RockPaperScissors(move1, move2 Move) (string, error) {
	if !move1.IsValid() || !move2.IsValid() {
		return "", fmt.Errorf("invalid Move")
	}

	winningMoves := map[Move]Move{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}

	if winningMoves[move1] == move2 {
		return fmt.Sprintf("%s Wins!", move1.String()), nil
	}

	if winningMoves[move2] == move1 {
		return fmt.Sprintf("%s Wins!", move2.String()), nil
	}

	return "Its a Draw!", nil
}
