package main

import (
	"aoc2022/Day7/fileh"
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strconv"
	// "log"
	// "math"
)

func readPrompt(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	// Read first line that call cd /, we init the hierarchy here
	fscanner.Scan()
	h := fileh.CreateHierarchy()

	ok := fscanner.Scan()
	// hierarchy building
	for ok {
		line := fscanner.Text()
		fmt.Printf("Read new command line : %s\n", line)
		cmdLine := strings.Split(line, " ")
		if cmdLine[0] == "$" {
			switch cmdLine[1] {
			case "ls":
				args := []string{}
				for ok {
					ok = fscanner.Scan()
					line = fscanner.Text()
					if !ok || line[0] == '$' {
						break
					}
					args = append(args, line)
				}
				h.Ls(args)
				h.Print()
			case "cd":
				h.Cd(cmdLine[2])
				ok = fscanner.Scan()
				//fmt.Printf("New current : %s\n", h.GetCurrentName())
			}
		}

	}
	fmt.Println(h)
	//compute solution
	res := h.ComputeSumFolderSizeWithSizeBellow(100000)
	//res := 0
	return res
}

func main() {
	res := readPrompt("input.txt")

	// sp := readInputState("input_test_state.txt")
	// sp.readMoves("input_test_moves.txt")

	fmt.Printf("\nres : %v\n", res)
}
