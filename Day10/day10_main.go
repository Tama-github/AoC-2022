package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	// "log"
	// "math"
)

type CRT struct {
	screen      []string
	currentLine int
	m           sync.Mutex
}

func (s *CRT) Print() {
	for _, str := range s.screen {
		fmt.Println(str)
	}
}

func InitCRT() *CRT {
	return &CRT{screen: []string{"", "", "", "", "", ""},
		currentLine: 0}
}

func (s *CRT) updateScreen() {
	s.m.Lock()
	idx := pmod(int(elapsed()-1), 40)
	lineidx := int(float64(elapsed()-1) / float64(40))
	fmt.Printf("line : %v, char : %v\n", lineidx, idx)
	if lineidx > 5 {
		return
	}
	line := s.screen[lineidx]
	if idx >= Counter-1 && idx <= Counter+1 {
		line = fmt.Sprintf("%s%c", line, '#')
	} else {
		line = fmt.Sprintf("%s%c", line, '.')
	}
	s.screen[lineidx] = line
	s.m.Unlock()
}

var tickTime = 100 * time.Millisecond
var start = time.Now()
var Counter int = 1
var nextToAdd = 0
var nextAction <-chan time.Time
var actionToDo = func() {
	//fmt.Printf("     X = %v + %v = %v\n", Counter, nextToAdd, Counter+nextToAdd)
	Counter += nextToAdd
}
var signalStrength = 0

func computeSignalStength(tickNum int) {
	//fmt.Printf("    reached %v: sig strength %v * %v = %v\n", tickNum, tickNum, Counter, tickNum*Counter)
	signalStrength += tickNum * Counter
}

func processNextCmd(fscanner *bufio.Scanner) bool {
	if !fscanner.Scan() {
		return false
	}

	cmd := strings.Split(fscanner.Text(), " ")
	fmt.Printf("     initiate next action : %s", cmd[0])
	switch cmd[0] {
	case "noop":
		fmt.Println()
		nextAction = time.After(tickTime)
		nextToAdd = 0
	case "addx":
		nextAction = time.After(tickTime * 2)
		nextToAdd, _ = strconv.Atoi(cmd[1])
		fmt.Printf(" %v\n", nextToAdd)
	}
	return true
}

func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

func elapsed() time.Duration {
	return time.Since(start).Round(time.Millisecond) / tickTime
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

	//g := initGround()
	//res := 0
	screen := InitCRT()
	tick := time.Tick(tickTime)
	action40 := time.After(tickTime * 40)
	action80 := time.After(tickTime * 80)
	action120 := time.After(tickTime * 120)
	action160 := time.After(tickTime * 160)
	action200 := time.After(tickTime * 200)
	action240 := time.After(tickTime * 240)

	// elapsed := func() time.Duration {
	// 	return time.Since(start).Round(time.Millisecond) / tickTime
	// }
	//fmt.Printf("[%v] start.\n", elapsed())
	stop := !processNextCmd(fscanner)

	for !stop {
		select {
		case <-tick:
			screen.updateScreen()
		//fmt.Printf("[%v] tick.\n", elapsed())

		default:
			select {
			case <-action40:
				fmt.Printf("[%v] compute sig 40\n", elapsed())
				//computeSignalStength(20)
				screen.currentLine++
			case <-action80:
				fmt.Printf("[%v] compute sig 80\n", elapsed())
				//computeSignalStength(60)
				screen.currentLine++
			case <-action120:
				fmt.Printf("[%v] compute sig 120\n", elapsed())
				//computeSignalStength(100)
				screen.currentLine++
			case <-action160:
				fmt.Printf("[%v] compute sig 160\n", elapsed())
				//computeSignalStength(140)
				screen.currentLine++
			case <-action200:
				fmt.Printf("[%v] compute sig 200\n", elapsed())
				//computeSignalStength(180)
				screen.currentLine++
			case <-action240:
				fmt.Printf("[%v] compute sig 240\n", elapsed())
				//computeSignalStength(220)
				screen.currentLine++

			default:
				select {
				case <-nextAction:
					//fmt.Printf("[%v] reach next action.\n", elapsed())
					actionToDo()
					stop = !processNextCmd(fscanner)
				default:
				}

			}
		}

		//time.Sleep(tickTime)
	}
	screen.Print()
	return signalStrength
}

func main() {
	readPrompt("input.txt")

	// sp := readInputState("input_test_state.txt")
	// sp.readMoves("input_test_moves.txt")

	//fmt.Printf("\nres : %v\n", res)
}
