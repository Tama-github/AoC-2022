package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// read from file "input.txt" and fill an array with the data

/** Elfs **/
type Elf struct {
	foods    []int
	calories int
}

func (elf *Elf) FoodSum() int {
	sum := 0
	for _, calories := range elf.foods {
		sum += calories
	}
	return sum
}

/** Expedition **/
type Expedition struct {
	elfs []*Elf
}

func (exp *Expedition) AddElf(e *Elf) {
	exp.elfs = append(exp.elfs, e)
}

func (exp *Expedition) SearchForFood() (elfId int, caloriesCarried int) {
	resId := -1
	resFood := 0
	for i, elf := range exp.elfs {
		if elf.calories > resFood {
			resFood = elf.calories
			resId = i
		}
	}
	return resId, resFood
}

func (exp *Expedition) SearchTopForFood(topSize int) int {
	// Make the top
	top := InitTopFromExpedition(exp, topSize)

	for _, elf := range exp.elfs {
		top.Update(elf)
	}

	return top.sum
}

type TopFood struct {
	sizeOfTop int
	top       []*Elf
	sum       int
	poorest   *Elf
	poorestId int
}

func (tf *TopFood) Print() {
	fmt.Printf("\nTop %v\n\tcontent : [", tf.sizeOfTop)
	for _, elf := range tf.top {
		fmt.Printf("%v, ", elf.calories)
	}
	fmt.Printf("]\n\tsum : %v | poorest : %v | id : %v\n\n", tf.sum, tf.poorest.calories, tf.poorestId)
}

func InitTopFromExpedition(exp *Expedition, topSize int) *TopFood {
	resTop := &TopFood{}
	if len(exp.elfs) >= topSize {
		resTop.top = exp.elfs[:topSize]
	} else {
		resTop.top = exp.elfs[:]
	}
	resTop.GetTopSum()
	resTop.GetPoorest()
	return resTop
}

func (tf *TopFood) Update(elf *Elf) {

	//fmt.Printf("Try to add %v in the top where the poorest is %v\n", elf.calories, tf.poorest.calories)
	if elf.calories > tf.poorest.calories {
		fmt.Printf("_____________________________________________\nstate of the top before insertion :\n")
		tf.Print()
		fmt.Printf("Inserting %v, to replace %v\n", elf.calories, tf.top[tf.poorestId].calories)
		// tf.sum -= tf.top[tf.poorestId].calories
		// tf.sum += elf.calories

		tf.top[tf.poorestId] = elf
		tf.GetTopSum()
		tf.GetPoorest()
		//fmt.Printf("The new poorest is %v\n", tf.poorest.calories)

		fmt.Printf("state of the top efter insertion :\n")
		tf.Print()
	}
}

func (tf *TopFood) GetTopSum() int {
	sum := 0
	for _, elf := range tf.top {
		sum += elf.calories
	}
	tf.sum = sum
	return sum
}

func (tf *TopFood) GetPoorest() (poorest *Elf, poorestId int) {
	minValue := math.MaxInt
	minId := 0
	for i, elf := range tf.top {
		if elf.calories < minValue {
			minValue = elf.calories
			minId = i
		}
	}
	tf.poorestId = minId
	tf.poorest = tf.top[minId]
	return tf.top[minId], minId
}

func convertStringToInt(toConvert string) int64 {
	i, err := strconv.ParseInt(toConvert, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func (exp *Expedition) populateExpedition(filepath string) {
	// read the whole file at once
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	fscanner := bufio.NewScanner(file)

	foods := make([]int, 0)
	calories := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		if line != "" {
			food := int(convertStringToInt(line))
			foods = append(foods, food)
			calories += food
			fmt.Printf("Food : %v | Food(int) : %v | sum : %v\n", line, food, calories)
		} else {
			exp.AddElf(&Elf{foods: foods, calories: calories})
			foods = make([]int, 0)
			calories = 0
		}

		// fmt.Printf(line1)
		// if strings.Index(line, "\n") == 0 {

		// }

	}

}

func main() {
	expedition := &Expedition{make([]*Elf, 0)}

	// Read and populate data
	expedition.populateExpedition("input.txt")

	// ex 1
	// elfId, caloriesCarried := expedition.SearchForFood()

	// ex 2
	topSize := 3
	topSumCaloriesCarried := expedition.SearchTopForFood(topSize)

	// output
	fmt.Printf("The top %v elfs carries %v calories.\n", topSize, topSumCaloriesCarried)
}
