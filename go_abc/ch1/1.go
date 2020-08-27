package main

import (
	"fmt"
	"os"
)

func main()  {
	var msg []string
	if len(os.Args) > 1 {
		msg = os.Args
	}
	fmt.Println("Hello", msg)
	os.Exit(0)
}

