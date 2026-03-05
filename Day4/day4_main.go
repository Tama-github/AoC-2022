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

func getInterval(interval []string) []int {
	min, _ := strconv.Atoi(interval[0])
	max, _ := strconv.Atoi(interval[1])
	return []int{min, max}
}

func checkIfIntervalAreIncluded(A []int, B []int) bool {
	if A[0] <= B[0] && A[1] >= B[1] {
		fmt.Printf("[%v,%v] is included in [%v,%v]\n\n", B[0], B[1], A[0], A[1])
		return true
	}
	if A[0] >= B[0] && A[1] <= B[1] {
		fmt.Printf("[%v,%v] is included in [%v,%v]\n\n", A[0], A[1], B[0], B[1])
		return true
	}

	fmt.Printf("[%v,%v] and [%v,%v] are not included in each others\n\n", A[0], A[1], B[0], B[1])
	return false
}

func checkIfIntervalAreOverlaped(A []int, B []int) bool {
	if A[0] <= B[0] && A[1] >= B[1] {
		fmt.Printf("[%v,%v] is included in [%v,%v]\n\n", B[0], B[1], A[0], A[1])
		return true
	}
	if A[0] >= B[0] && A[1] <= B[1] {
		fmt.Printf("[%v,%v] is included in [%v,%v]\n\n", A[0], A[1], B[0], B[1])
		return true
	}

	if A[1] >= B[0] && A[1] <= B[1] {
		fmt.Printf("[%v,%v] is included in [%v,%v]\n\n", A[0], A[1], B[0], B[1])
		return true
	}
	if A[0] >= B[0] && A[0] <= B[1] {
		fmt.Printf("[%v,%v] is included in [%v,%v]\n\n", A[0], A[1], B[0], B[1])
		return true
	}

	fmt.Printf("[%v,%v] and [%v,%v] are not included in each others\n\n", A[0], A[1], B[0], B[1])
	return false
}

func readAllPaires(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	totalScore := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		pairs := strings.Split(line, ",")
		fmt.Printf("Check %s", line)
		if checkIfIntervalAreIncluded(getInterval(strings.Split(pairs[0], "-")), getInterval(strings.Split(pairs[1], "-"))) {
			totalScore++
		}
	}

	return totalScore
}

func readAllPaires2(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	totalScore := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		pairs := strings.Split(line, ",")
		fmt.Printf("Check %s", line)
		if checkIfIntervalAreOverlaped(getInterval(strings.Split(pairs[0], "-")), getInterval(strings.Split(pairs[1], "-"))) {
			totalScore++
		}
	}

	return totalScore
}

func main() {
	res := readAllPaires2("input_test.txt")
	fmt.Printf("res : %v\n", res)
}
