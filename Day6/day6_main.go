package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	// "log"
	// "math"
)

type Buffer struct {
	size   int
	buffer map[byte]int
}

func (b *Buffer) Print() {
	fmt.Printf("[")
	for k, v := range b.buffer {
		fmt.Printf(" (%c, %v) ", k, v)
	}
	fmt.Printf("]\n")
}

func (b *Buffer) Init(size int, s string) bool {
	b.size = size
	b.buffer = map[byte]int{}
	for _, c := range s {
		if b.Add(byte(c)) {
			return true
		}
	}
	return false
}

func (b *Buffer) Remove(c byte) {
	b.buffer[c]--
	if b.buffer[c] == 0 {
		delete(b.buffer, c)
	}
}

func (b *Buffer) Add(c byte) bool {
	//fmt.Print(b.buffer[byte(c)])
	b.buffer[byte(c)]++
	if len(b.buffer) >= b.size {
		return true
	}
	return false
}

func searchForFirstIndicator(sig string, size int) int {
	i := size
	bu := &Buffer{}
	fmt.Printf("\nsearch for marker in : %s\n", sig)

	if bu.Init(i, sig[:i]) {
		return i
	}
	fmt.Printf("init : \n")
	bu.Print()

	fmt.Printf("search:\n")
	for _, c := range sig[bu.size:] {
		rem := sig[i-bu.size]
		add := byte(c)
		fmt.Printf("remove %c and add %c : ", rem, add)
		bu.Remove(rem)
		i++
		if bu.Add(add) {
			fmt.Printf("Found marker at %v\n", i)
			return i
		}

		bu.Print()
	}
	return i
}

func readMoves(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	// ex1
	//size1 := 4

	//ex2
	size2 := 14

	res := size2
	for fscanner.Scan() {
		line := fscanner.Text()
		res = searchForFirstIndicator(line, size2)
		fmt.Print(res)
	}

	return res
}

func main() {
	res := readMoves("input.txt")

	// sp := readInputState("input_test_state.txt")
	// sp.readMoves("input_test_moves.txt")

	fmt.Printf("\nres : %v\n", res)
}
