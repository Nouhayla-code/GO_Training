package main

import (
	"GO_Training/exercises"
	"fmt"
)

func main() {

	//returnerer 2 værdier
	hej, _ := exercises.SayHi()


	//error handling 
if hej!=nil {
		fmt.Println(hej)
}


	//variabler

	//const firstName string = "john"
	// var lastName string = "smith"

	//var lastName = "smith";

	// redeklerer en variabel 
	//lastName:= "smith"
	//lastName = "smith2"
	//fmt.Println((lastName))

	
	// tilføj ting ind i en liste 
	numberList:= []int{2,1,4,5,6}
	numberList= append(numberList, 9)
	doubleList:= []int{}

		fmt.Println(numberList)
		


		//et looop som kørere igennem numberlist og fordobler værdierne og sætter ind i doublelist 
	for i:=0;i< len(numberList); i++ {
		doubleList= append(doubleList,numberList[i]*2 )
	}

	fmt.Println(doubleList)

}



func doubleLetters(list []int) ([]int, error) {

	// move slice and loop
	// if else
	//return multiple

	if len(list) < 3 {
		err := fmt.Errorf("list must be at least 3 items long")
		return []int{}, err
	}

	doubleList := []int{}
	for i := 0; i < len(list); i++ {
		if i < 3 {
			doubleList = append(doubleList, list[i]*2)
		} else {
			doubleList = append(doubleList, list[i])
		}
	}
	return doubleList, nil
}

