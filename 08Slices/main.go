package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to silces tut")

	var fruitlist = []string{"apple","tomoatos","banna"}
	fmt.Printf("Types of fruitlist is %T",fruitlist)

	fruitlist = append(fruitlist, "Mango","oranges")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0]= 234
	highScores[1]= 2355
	highScores[2]= 236
	highScores[3]= 236
	// highScores[4]= 240
	// append will relocate the memory 

	highScores = append(highScores, 554,666,562)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"react js","javascript","java","go-lang"}
	fmt.Println(courses)
	var index int = 2;
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}