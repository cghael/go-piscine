package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	USAGE          = "USAGE:\t-mn\n\t-md\n\t-mo\n\t-sd\n\tor no flags"
	INVVALID_INPUT = "Invalid input"
	EXIT_ERROR     = 1
)

func ftCheckArgs(args []string) {
	flags := [4]string{"-mn", "-md", "-mo", "-sd"}
	for _, v := range args {
		res := 0
		for _, v_ := range flags {
			if v == v_ {
				res = 1
			}
		}
		if res == 0 {
			fmt.Println(USAGE)
			os.Exit(0)
		}
	}
}

func ftParsing() (numbers []int) {
	numbers = make([]int, 0)
	var input string

	for length, err1 := fmt.Scanln(&input); length > 0; length, err1 = fmt.Scanln(&input) {
		if err1 != nil {
			fmt.Println(INVVALID_INPUT)
			os.Exit(EXIT_ERROR)
		}
		if num, err2 := strconv.Atoi(input); err2 != nil || num > 100000 || num < -100000 {
			fmt.Println(INVVALID_INPUT)
			os.Exit(EXIT_ERROR)
		} else {
			numbers = append(numbers, num)
		}
	}
	return
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		args = append(args, "-mn", "-md", "-mo", "-sd")
	} else {
		ftCheckArgs(args)
	}
	numbers := ftParsing()
	ftCountMetrics(args, numbers)
}
