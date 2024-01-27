package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - learncodeOnline.in"

	file , err := os.Create("./mylcogofiles.txt")

	if err !=nil{
		panic(err)
	}

	length , err:= io.WriteString(file, content)
	if err !=nil{
		panic(err)
	}
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filname string){
	databytes, err := ioutil.ReadFile(filname)
	if err != nil{
		panic(err)
	}
	fmt.Println("Text data inside the files is \n", databytes)

}