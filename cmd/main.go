package main

import (
	"bufio"
	"fmt"
	"github.com/FabianTe/spongecase"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		applyToArgs()
	} else {
		applyToStdin()
	}
}

func applyToStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(spongecase.ApplyStr(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func applyToArgs() {
	for i, arg := range os.Args[1:] {
		str := spongecase.ApplyStr(arg)
		if i < len(os.Args) - 2 {
			str += " "
		}
		fmt.Print(str)
	}
	fmt.Println()
}
