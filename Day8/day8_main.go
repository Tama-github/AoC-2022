package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "log"
	// "math"
)

type Forest [][]int

func (f Forest) Print() {
	fmt.Println()
	for _, vx := range f {
		for _, v := range vx {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
}

func (f Forest) scenicScore(x, y int) int {
	return f.scenicScoreFromBottom(x, y) * f.scenicScoreFromLeft(x, y) * f.scenicScoreFromRight(x, y) * f.scenicScoreFromTop(x, y)
}

func (f Forest) isVisible(x, y int) bool {
	return f.isVisibleFromBottom(x, y) || f.isVisibleFromLeft(x, y) || f.isVisibleFromRight(x, y) || f.isVisibleFromTop(x, y)
}

func (f Forest) isVisibleFromTop(x, y int) bool {
	if x == 0 {
		return true
	}
	selfHeight := f[x][y]
	for i := x - 1; i >= 0; i-- {
		if selfHeight <= f[i][y] {
			return false
		}
	}
	return true
}

func (f Forest) isVisibleFromBottom(x, y int) bool {
	if x == len(f)-1 {
		return true
	}
	selfHeight := f[x][y]
	for i := x + 1; i < len(f); i++ {
		if selfHeight <= f[i][y] {
			return false
		}
	}
	return true
}

func (f Forest) isVisibleFromLeft(x, y int) bool {
	if y == 0 {
		return true
	}
	selfHeight := f[x][y]
	for i := y - 1; i >= 0; i-- {
		if selfHeight <= f[x][i] {
			return false
		}
	}
	return true
}

func (f Forest) isVisibleFromRight(x, y int) bool {
	if y == len(f[x]) {
		return true
	}
	selfHeight := f[x][y]
	for i := y + 1; i < len(f[x]); i++ {
		if selfHeight <= f[x][i] {
			return false
		}
	}
	return true
}

func (f Forest) scenicScoreFromTop(x, y int) int {
	if x == 0 {
		return 0
	}
	selfHeight := f[x][y]
	res := 0
	for i := x - 1; i >= 0; i-- {
		res++
		if selfHeight <= f[i][y] {
			break
		}
	}
	return res
}

func (f Forest) scenicScoreFromBottom(x, y int) int {
	if x == len(f)-1 {
		return 0
	}
	selfHeight := f[x][y]
	res := 0
	for i := x + 1; i < len(f); i++ {
		res++
		if selfHeight <= f[i][y] {
			break
		}
	}
	return res
}

func (f Forest) scenicScoreFromLeft(x, y int) int {
	if y == 0 {
		return 0
	}
	selfHeight := f[x][y]
	res := 0
	for i := y - 1; i >= 0; i-- {
		res++
		if selfHeight <= f[x][i] {
			break
		}
	}
	return res
}

func (f Forest) scenicScoreFromRight(x, y int) int {
	if y == len(f[x]) {
		return 0
	}
	selfHeight := f[x][y]
	res := 0
	for i := y + 1; i < len(f[x]); i++ {
		res++
		if selfHeight <= f[x][i] {
			break
		}
	}
	return res
}

func (f Forest) getBestScenicScore() int {
	best := 0
	for i, vx := range f {
		for j := range vx {
			tmp := f.scenicScore(i, j)
			if tmp > best {
				best = tmp
			}
		}
	}
	return best
}

func readPrompt(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	// Read first line that call cd /, we init the hierarchy here
	res := 0
	f := Forest{}
	for fscanner.Scan() {
		line := fscanner.Text()
		tmp := make([]int, 0)
		for _, c := range line {
			i, _ := strconv.Atoi(string(c))
			tmp = append(tmp, i)
		}
		f = append(f, tmp)
		fmt.Print(res)
	}
	f.Print()
	// ex 1
	// for i, vx := range f {
	// 	for j := range vx {

	// 		if f.isVisible(i, j) {
	// 			res++
	// 		}
	// 	}
	// }

	// ex 2
	res = f.getBestScenicScore()

	return res
}

func main() {
	res := readPrompt("input.txt")

	// sp := readInputState("input_test_state.txt")
	// sp.readMoves("input_test_moves.txt")

	fmt.Printf("\nres : %v\n", res)
}
