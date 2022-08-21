package main

import (
	"fmt"
	"os"
)

func main() {
	action := ""
	if len(os.Args) >= 2 {
		action = os.Args[1]
	}

	switch action {
	case "generate":
		generate()
	case "search":
		search()
	default:
		fmt.Println("./five-clique [action]")
		fmt.Println("  action = generate, search")
	}
}
