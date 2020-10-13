package main

import (
	"bufio"
	"fmt"
	"os"
	sponge_case "sponge-case/sponge-case"
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
		fmt.Println(sponge_case.ApplyStr(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func applyToArgs() {
	fmt.Println(sponge_case.ApplyStr(os.Args[1]))
}
