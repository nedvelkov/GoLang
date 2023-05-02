package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

loop:
	for {

		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
		fmt.Println(strings.Repeat("#", 100))
	}
}
