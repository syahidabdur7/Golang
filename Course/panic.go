package main

import "fmt"

func panicApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(b bool) {
	defer panicApp()
	if b {
		panic("erorr")
	}
	fmt.Println("App Running")
}

func main() {
	runApp(false)
}
