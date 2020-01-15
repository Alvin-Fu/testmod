package testmod

import "fmt"

func Printf(args ...interface{}) {
	fmt.Println("welcome")
	fmt.Println(args...)
	fmt.Println("See you!")
}
