package main

import (
	"fmt"

	"github.com/Microsoft/go-winio"
)

// dial mpv pipe
func dial() {
	fmt.Println("Attempting to connect to pipe!")
	_, err := winio.DialPipe("\\\\.\\pipe\\mpvsocket", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func main() {
    dial()
}

