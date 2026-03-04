package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
	// "log"
	// "math"
)

type Element struct {
	name  string
	point int
	// store match result 0 lose | 3 draw | 6 win
	states  map[*Element]int
	states2 map[int]*Element
}

var rock = &Element{
	name:    "Rock",
	point:   1,
	states:  map[*Element]int{},
	states2: map[int]*Element{},
}
var paper = &Element{
	name:    "Paper",
	point:   2,
	states:  map[*Element]int{},
	states2: map[int]*Element{},
}
var scissors = &Element{
	name:    "Scissors",
	point:   3,
	states:  map[*Element]int{},
	states2: map[int]*Element{},
}

func initGameState() {
	rock.states = map[*Element]int{
		rock:     3,
		paper:    0,
		scissors: 6,
	}
	rock.states2 = map[int]*Element{
		3: rock,
		6: paper,
		0: scissors,
	}
	paper.states = map[*Element]int{
		rock:     6,
		paper:    3,
		scissors: 0,
	}
	paper.states2 = map[int]*Element{
		0: rock,
		3: paper,
		6: scissors,
	}
	scissors.states = map[*Element]int{
		rock:     0,
		paper:    6,
		scissors: 3,
	}
	scissors.states2 = map[int]*Element{
		6: rock,
		0: paper,
		3: scissors,
	}
}

func convertStringToGameElem(s string) *Element {
	switch s {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissors
	default:
		fmt.Printf("This is not a valid move %c\n", s)
	}
	return rock
}

func (e *Element) convertStringToGameResult(s string) *Element {
	switch s {
	case "X":
		return e.states2[0]
	case "Y":
		return e.states2[3]
	case "Z":
		return e.states2[6]
	default:
		fmt.Printf("This is not a valid move %c\n", s)
	}
	return e
}

func readStrategyAndComputExpectedResult(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	totalScore := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		play := strings.Split(line, " ")
		e := convertStringToGameElem(play[1])
		totalScore += e.states[convertStringToGameElem(play[0])] + e.point
	}

	return totalScore
}

func readStrategyAndComputExpectedResult2(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	totalScore := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		play := strings.Split(line, " ")
		advPlay := convertStringToGameElem(play[0])
		mePlay := advPlay.convertStringToGameResult(play[1])
		score := mePlay.point + mePlay.states[advPlay]
		totalScore += score
		fmt.Printf("Adversary play %s(%s), I should play %s(%s) (it would give %v + %v = %v points\n", advPlay.name, play[0], mePlay.name, play[1], mePlay.point, mePlay.states[advPlay], score)
	}

	return totalScore
}

func main() {
	initGameState()
	sum := readStrategyAndComputExpectedResult2("input.txt")
	fmt.Printf("sum : %v\n", sum)
}
