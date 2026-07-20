package main

import (
	"fmt"
	"practic/internal"
)

func main() {
	if err := internal.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}

}
