package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "log"
	// "math"
)

type Pos struct {
	x int
	y int
}

func (p Pos) toString() string {
	return fmt.Sprintf("%v%v", p.x, p.y)
}

type Ground struct {
	tiles map[string]bool
	//ex 1
	h *Pos
	t *Pos
	//ex 2
	knots []*Pos
}

func initGround() *Ground {
	g := &Ground{
		tiles: make(map[string]bool),
		h:     &Pos{0, 0},
		t:     &Pos{0, 0},
		knots: make([]*Pos, 10)}
	for i := 0; i < 10; i++ {
		g.knots[i] = &Pos{0, 0}
	}
	return g
}

func (g *Ground) exec(cmd string, repeat int) {
	fmt.Printf("$ %s %v\n", cmd, repeat)
	for i := 0; i < repeat; i++ {
		switch cmd {
		case "R":
			g.Right()
		case "L":
			g.Left()
		case "U":
			g.Up()
		case "D":
			g.Down()
		}
		g.tiles[g.t.toString()] = true
	}
}

func (g *Ground) Right() {
	fmt.Printf("    Move Right\n")
	g.h.x++
	dx := g.h.x - g.t.x
	if dx == 2 {
		g.t.x++
		if g.t.y != g.h.y {
			g.t.y = g.h.y
		}
	}
}

func (g *Ground) Left() {
	fmt.Printf("    Move Left\n")
	g.h.x--
	dx := g.h.x - g.t.x
	if dx == -2 {
		g.t.x--
		if g.t.y != g.h.y {
			g.t.y = g.h.y
		}
	}
}

func (g *Ground) Up() {
	fmt.Printf("    Move Up\n")
	g.h.y++
	dy := g.h.y - g.t.y
	if dy == 2 {
		g.t.y++
		if g.t.x != g.h.x {
			g.t.x = g.h.x
		}
	}
}

func (g *Ground) Down() {
	fmt.Printf("    Move Down\n")
	g.h.y--
	dy := g.h.y - g.t.y
	if dy == -2 {
		g.t.y--
		if g.t.x != g.h.x {
			g.t.x = g.h.x
		}
	}
}

func (g *Ground) exec2(cmd string, repeat int) {
	fmt.Printf("$ %s %v\n", cmd, repeat)
	for i := 0; i < repeat; i++ {
		g.MoveKnot(cmd, 0)
	}
}

func (g *Ground) MoveKnot(cmd string, i int) {
	switch cmd {
	case "R":
		g.knots[0].x++
	case "L":
		g.knots[0].x--
	case "U":
		g.knots[0].y++
	case "D":
		g.knots[0].y--
	}
	g.SolveKnot(1)

	// Write in map for the solution
	g.tiles[g.knots[len(g.knots)-1].toString()] = true
}

func signInt(i int) int {
	if math.Signbit(float64(i)) {
		return -1
	}
	return 1
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (g *Ground) SolveKnot(i int) {
	me := g.knots[i]
	ahead := g.knots[i-1]
	dx := ahead.x - me.x
	dy := ahead.y - me.y

	if absInt(dx)*absInt(dy) == 4 { // Diagonal move
		fmt.Printf("Diagonal move ahead[%v]=(%v, %v) / me[%v]=(%v, %v)\n", i-1, ahead.x, ahead.y, i, me.x, me.y)
		me.x += 1 * signInt(dx)
		me.y += 1 * signInt(dy)
	} else if absInt(dx)*absInt(dy) == 2 { // knight move
		fmt.Printf("Knight move ahead(%v, %v) / me(%v, %v)\n", ahead.x, ahead.y, me.x, me.y)
		me.x += 1 * signInt(dx)
		me.y += 1 * signInt(dy)
	} else if absInt(dx) == 2 { // Horizontal move
		fmt.Printf("Horizontal move ahead(%v, %v) / me(%v, %v)\n", ahead.x, ahead.y, me.x, me.y)
		me.x += 1 * signInt(dx)
	} else if absInt(dy) == 2 { // Vertical move
		fmt.Printf("Vertical move ahead(%v, %v) / me(%v, %v)\n", ahead.x, ahead.y, me.x, me.y)
		me.y += 1 * signInt(dy)
	}

	if i == len(g.knots)-1 {
		return
	}
	g.SolveKnot(i + 1)
}

func readPrompt(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	// Read first line that call cd /, we init the hierarchy here
	//res := 0

	g := initGround()

	for fscanner.Scan() {
		line := fscanner.Text()
		cmd := strings.Split(line, " ")
		n, _ := strconv.Atoi(cmd[1])
		g.exec2(cmd[0], n)
		//fmt.Print(res)
	}

	return len(g.tiles)
}

func main() {
	res := readPrompt("input.txt")

	// sp := readInputState("input_test_state.txt")
	// sp.readMoves("input_test_moves.txt")

	fmt.Printf("\nres : %v\n", res)
}
