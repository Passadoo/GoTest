package main

import "fmt"

func greeting(name string) string  {
	return "Hello " + name

}

func getSum(int1, int2 int) int {
	return int1 + int2	
}

func main()  {
	// Arrays
	var fruitArr [2]string

	//Assigning values
	fruitArr[0] = "Äpple"
	fruitArr[1] = "Päron"

	fruitArrInit := [3]string{"Apelsin", "Jordgubbe","Fikus"}
	fruitSlice := []string{"Apelsin", "Jordgubbe","Fikus","Sädmästare"}

	fmt.Println(fruitArrInit)
	fmt.Println(len(fruitSlice))

	fmt.Println(fruitArr[0]+fruitArr[1])

	fmt.Println(greeting("Simon"))
	fmt.Println(getSum(3, 6))
}
