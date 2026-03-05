package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "log"
	// "math"
)

func convertCharToWeight(c byte) int {
	// Convert lower case char from [97, 122] to [1,26]
	if c >= 97 && c <= 122 {
		return int(c) - 96
	}
	// Convert upper case char from [65, 90] to [27, 90]
	if c >= 65 && c <= 90 {
		return int(c) - 38
	}

	return 0
}

func intersectionOfTwoArrays(arr, arr2 string) int {
	//num map to keep track of al element
	numMap := map[byte]bool{}
	//adding all element of first array in hash table to checck the intersection in other array
	for _, value := range arr {
		numMap[byte(value)] = true
	}
	//result of intersect numbers
	result := []byte{}
	//keep track of result that no reapeated number added
	resultMap := map[byte]bool{}
	//checking the intersected element in second array
	for i := 0; i < len(arr2); i++ {
		//if the num is present in nummap means intersected and not in result then add in result
		if numMap[arr2[i]] && !resultMap[arr2[i]] {
			result = append(result, arr2[i])
			resultMap[arr2[i]] = true
			break
		}
	}
	res := convertCharToWeight(result[0])
	fmt.Printf("\nresult : %c(%v)\n", result[0], res)

	//return the result
	return res
}

func readAllBags(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	totalScore := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		comp1 := line[:len(line)/2]
		comp2 := line[len(line)/2:]
		res := intersectionOfTwoArrays(comp1, comp2)
		totalScore += res
		fmt.Printf("bag : %s\n", line)
		fmt.Printf("\t%s   |   %s\n", comp1, comp2)

	}

	return totalScore
}

func inter(arr, arr2 []byte) []byte {
	//num map to keep track of al element
	numMap := map[byte]bool{}
	//adding all element of first array in hash table to checck the intersection in other array
	for _, value := range arr {
		numMap[value] = true
	}
	//result of intersect numbers
	result := []byte{}
	//keep track of result that no reapeated number added
	resultMap := map[byte]bool{}
	//checking the intersected element in second array
	for i := 0; i < len(arr2); i++ {
		//if the num is present in nummap means intersected and not in result then add in result
		if numMap[arr2[i]] && !resultMap[arr2[i]] {
			result = append(result, arr2[i])
			resultMap[arr2[i]] = true
		}
	}
	//return the result
	return result
}

func intersectionOfThreeArrays(arr, arr2, arr3 string) int {

	arrtmp := inter([]byte(arr), []byte(arr2))
	res := inter(arrtmp, []byte(arr3))
	itemScore := convertCharToWeight(res[0])
	fmt.Printf("the common item is '%c'(%v)\n\n", res[0], itemScore)

	return itemScore
}

func read3Bags(fscanner *bufio.Scanner) (r int, c bool) {
	bags := make([]string, 3)

	for i := 0; i < 3; i++ {
		if !fscanner.Scan() {
			return 0, false
		}

		bags[i] = fscanner.Text()
	}
	fmt.Printf("3 bags :\n")
	fmt.Print("\t", bags[0], "\n")
	fmt.Print("\t", bags[1], "\n")
	fmt.Print("\t", bags[2], "\n")

	res := intersectionOfThreeArrays(bags[0], bags[1], bags[2])

	return res, true
}

func readAllBags2(filePath string) int {
	// read the whole file at once
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	totalScore := 0
	c := true
	r := 0
	for c {
		r, c = read3Bags(fscanner)
		totalScore += r
	}

	return totalScore
}

func main() {
	res := readAllBags2("input.txt")
	fmt.Printf("res : %v\n", res)
}
