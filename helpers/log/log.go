package log

import "fmt"

// Err - Print the error
func Err(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
