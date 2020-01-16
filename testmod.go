package testmod

import "fmt"

func Printf(args ...interface{}) {
	fmt.Println("Welcome !")
	fmt.Println(args...)
	fmt.Println("See you!")
}
