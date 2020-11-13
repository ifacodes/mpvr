package main

import (
	"fmt"

	"github.com/Microsoft/go-winio"
)

// dial mpv pipe
func dial() {
	conn, err := winio.DialPipe("\\\\.\\pipe\\mpvsocket", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	conn2, err := winio.DialPipe("\\\\.\\pipe\\mpvsocket", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	num, err := conn.Write([]byte("echo show-text ${playback-time}"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("num: ", num)

	var buf = make([]byte, 4096)
	num, err = conn2.Read(buf)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("num: {0}\nbuf: {1}", num, buf)
}

func main() {
    dial()
}

