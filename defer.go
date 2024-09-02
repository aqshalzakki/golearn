package main

import (
	"fmt"
)

func logging(channel string) {
	fmt.Println("logging to", channel, "channel")
}

func runApp() {
	fmt.Println("Running apps...")

	fmt.Println("Done Running apps...")
}

func main() {
	// defer digunakan untuk run function di akhir function
	defer logging("main")

	runApp()
}
