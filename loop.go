package main

import "fmt"

func main() {
	// loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Loop sekarang ", i)
	// }

	// names := [...]string{"Aqshal", "Ganteng", "Zakki"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Println("Index : ", index, " Value : ", value)
	// }

	// for _, value := range names {
	// 	fmt.Println(" Value : ", value)
	// }

	// Break Continue

	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// print only even number
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 15 {
			fmt.Println("Reached max value of 15")
			break
		}

		if numbers[i]%2 == 1 {
			continue
		}

		fmt.Println(numbers[i])
	}
}
