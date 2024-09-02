package main

import (
	"fmt"
	"strings"
)

func isEven(num int) bool {
	return num%2 == 0
}

func sayHello(name string) bool {
	fmt.Println("Hello", name)

	return true
}

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, firstName + " " + lastName
}

func namedReturnValue(firstName, lastName string) (lowerFirstName, upperLastName string) {
	lowerFirstName = strings.ToLower(firstName)
	upperLastName = strings.ToUpper(lastName)

	return lowerFirstName, upperLastName
}

func varags(name string, names ...string) string {
	namesSliced := append([]string{name}, names...)
	return strings.Join(namesSliced, ", ")
}

func sumAll(nums ...int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

// function as parameter
func sayHelloFiltered(name string, filterWord FilterWord) {
	fmt.Println(filterWord(name))
}

func filterWord(text string) string {
	if strings.ToLower(text) == "asu" {
		return "Asu is not allowed lil bro 💀"
	}
	return "Hello " + text
}

func registerUser(email string, isBlackListed IsBlackListed) {
	if isBlackListed(email) {
		fmt.Println("Gagal login! anda sudah ter-blacklist lil nig💀")
	} else {
		fmt.Println("selamat datang", email)
	}

}

func factorialLoopRecursive(num int) int {
	if num == 1 {
		return num
	} else {
		return num * factorialLoopRecursive(num-1)
	}
}

// Type Declarations
type FilterWord func(string) string
type IsBlackListed func(email string) bool

func main() {
	// _, fullName := getFullName("Aqshal", "Zakki")
	// fmt.Println(fullName)

	// fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// summarizeAll := sumAll

	// nums := []int{1, 2, 3, 4, 5, 40, 42, 12, 63}

	// fmt.Println(summarizeAll(nums...))

	// sayHelloFiltered("ASU", filterWord)

	// * Anonymous Function
	// registerUser("devsomething3@gmail.com", func(email string) bool {
	// 	blackListEmails := []string{
	// 		"kecoatempur2@gmail.com",
	// 		"devsomething3@gmail.com",
	// 		"guest@gmail.com",
	// 	}

	// 	result := false

	// 	for _, blackListEmail := range blackListEmails {
	// 		if blackListEmail == email {
	// 			result = true
	// 		}
	// 	}

	// 	return result
	// })

	fmt.Println(factorialLoopRecursive(5))
}
