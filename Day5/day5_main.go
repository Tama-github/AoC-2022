package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "log"
	// "math"
)

const QUANT_IDX = 1
const FROM_IDX = 3
const TO_IDX = 5

type Raw struct {
	storage []byte
}

func (r *Raw) Print() {
	fmt.Printf("[")
	for _, v := range r.storage {
		fmt.Printf(" %c ", v)
	}
	fmt.Printf("]")
}

type Stockpile struct {
	pile map[int]*Raw
}

func (sp *Stockpile) Print() {
	fmt.Printf("\nStockpile: \n")
	for i := 1; i <= len(sp.pile); i++ {
		sp.pile[i].Print()
		fmt.Printf(" | %v\n", i)
	}
	fmt.Printf("\n")
}

func getValueFromIndex(i int) int {
	return 1 + (i * 4)
}

func InitStockpile(state []string) *Stockpile {
	resPile := &Stockpile{}
	resPile.pile = make(map[int]*Raw)
	numberOfRawNeeded := strings.Fields(state[len(state)-1])
	fmt.Printf("the idxes %v? ", len(numberOfRawNeeded))
	fmt.Print(numberOfRawNeeded)
	fmt.Print("\n")
	for i, idxString := range numberOfRawNeeded {
		newRaw := &Raw{}
		fmt.Printf("Make Raw with :  [ ")
		for _, r := range state[:len(state)-1] {
			toAdd := r[getValueFromIndex(i)]

			if toAdd != ' ' {
				fmt.Printf("%c ", toAdd)
				newRaw.storage = append(newRaw.storage, toAdd)
			}
		}
		fmt.Printf("]\nAdd this raw to the pile : ")
		newRaw.Print()
		id, _ := strconv.Atoi(string(idxString[0]))
		fmt.Printf("\nAdd it to %v", id)
		resPile.pile[id] = newRaw
		fmt.Printf("\n\n")
	}
	resPile.Print()
	return resPile
}

func (sp *Stockpile) Move(quant int, fromId int, toId int) {

	fmt.Printf("\n\nHave to move %v items from %v to %v.", quant, fromId, toId)
	from := sp.pile[fromId]
	to := sp.pile[toId]
	fmt.Printf("\nBefore move :\nfrom ")
	from.Print()
	fmt.Printf("\nTo ")
	to.Print()
	for i := 0; i < quant; i++ {
		if len(from.storage) <= 0 {
			break
		}
		tmp := make([]byte, len(from.storage))
		copy(tmp, from.storage)
		to.storage = append(tmp[:1], to.storage...)
		from.storage = from.storage[1:]

	}

	fmt.Printf("\nAfter move :\nfrom ")
	from.Print()
	fmt.Printf("\nTo ")
	to.Print()
	sp.Print()

}

func (sp *Stockpile) Move2(quant int, fromId int, toId int) {

	fmt.Printf("\n\nHave to move %v items from %v to %v.", quant, fromId, toId)
	from := sp.pile[fromId]
	to := sp.pile[toId]
	fmt.Printf("\nBefore move :\nfrom ")
	from.Print()
	fmt.Printf("\nTo ")
	to.Print()
	tmp := make([]byte, len(from.storage))
	copy(tmp, from.storage)
	to.storage = append(tmp[:quant], to.storage...)
	from.storage = from.storage[quant:]

	fmt.Printf("\nAfter move :\nfrom ")
	from.Print()
	fmt.Printf("\nTo ")
	to.Print()
	sp.Print()

}

func (sp *Stockpile) GetTopMessage() string {
	res := make([]byte, len(sp.pile))

	for i := 1; i <= len(sp.pile); i++ {
		res[i-1] = sp.pile[i].storage[0]
	}
	return string(res)
}

func readInputState(filePath string) *Stockpile {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	stateArray := make([]string, 0)
	for fscanner.Scan() {
		stateArray = append(stateArray, fscanner.Text())
	}

	return InitStockpile(stateArray)
}

func (sp *Stockpile) readMoves(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	totalScore := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		w := strings.Split(line, " ")
		qid, _ := strconv.Atoi(w[QUANT_IDX])
		fid, _ := strconv.Atoi(w[FROM_IDX])
		tid, _ := strconv.Atoi(w[TO_IDX])
		//sp.Move(qid, fid, tid)
		sp.Move2(qid, fid, tid)
	}

	return totalScore
}

func main() {
	sp := readInputState("input_state.txt")
	sp.readMoves("input_moves.txt")

	// sp := readInputState("input_test_state.txt")
	// sp.readMoves("input_test_moves.txt")

	fmt.Printf("\nres : %v\n", sp.GetTopMessage())
}
