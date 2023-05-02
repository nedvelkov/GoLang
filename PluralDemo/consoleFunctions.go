package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConsoleCliApp() {
	fmt.Println("What to scream?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s)
}
