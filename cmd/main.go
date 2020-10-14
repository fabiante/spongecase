package main

import (
	"bufio"
	"fmt"
	"github.com/FabianTe/spongecase"
	"os"
)

func main() {
	if len(os.Args) == 2 {
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
	fmt.Println(spongecase.ApplyStr(os.Args[1]))
}
