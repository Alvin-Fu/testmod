package testmod

import "fmt"

func Printf(args ...interface{}) error {
	fmt.Println("Welcome !")
	if len(args) <= 0 {
		return fmt.Errorf("args err!")
	}
	fmt.Println(args...)
	fmt.Println("See you!")
	return nil
}
