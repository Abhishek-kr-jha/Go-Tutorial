package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["Java"] = "java"
	languages["Py"] = "Python"

	fmt.Println("list of all the languagaes:", languages)
	fmt.Println("JS shorts for:", languages["js"])

	delete(languages, "Py")
	fmt.Println("list of all the languagaes:", languages)


	// loop are  intresting in golang

	for key , value := range languages {
		fmt.Printf("for key %v, values is %v\n", key, value)
	}
}