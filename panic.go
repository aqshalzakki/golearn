package main

import (
	"fmt"
)

func log(channel string) {
	fmt.Println("logging to", channel, "channel")
}

func startApp(error bool) {
	// dijalankan ketika function ini selesai dijalankan,
	// meskipun ada error / panic, akan tetap dijalankan
	defer stopApp()

	fmt.Println("Starting apps...")

	if error {
		panic("Error! Data Leak has detected")
	}

	fmt.Println("Apps has started...")
}

func stopApp() {
	fmt.Println("Stopping apps")

	recoverMsg := recover()

	if recoverMsg != nil {
		fmt.Println("Recovered Message :", recoverMsg)
	}
}

func main() {

	startApp(true)
}
