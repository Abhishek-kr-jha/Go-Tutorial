package main

import "fmt"

func main() {
	fmt.Println("Loop is fun...!!")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday"}
	fmt.Println(days)

	// for loop in go

	// for d:= 0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day:= range days{
	// 	fmt.Printf("index is  and value is %v\n",  day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			// break
			continue
		}

		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumming  at lco...!")

}
