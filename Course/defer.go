package main

import (
	"fmt"
)

func logApp() {
	fmt.Println("Ini Deffer")
}

func defferApp(a int) {
	defer logApp()
	fmt.Println("Function")
	b := 10 / a
	fmt.Println("a = ", b)

}

func main() {
	defferApp(0)
}
