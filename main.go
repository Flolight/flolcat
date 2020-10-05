package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
)

func rgb(i int, c float64) (int, int, int) {
	return int(math.Sin(c*float64(i)+0)*127 + 128),
		int(math.Sin(c*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(c*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune, c float64) {
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j, c)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune
	var coef float64

	flag.Float64Var(&coef, "c", 0.1, "coefficient for rainbow coloring. Valid values are between 0 and 1")
	flag.Parse()
	if coef > 1 || coef < 0 {
		fmt.Println("Invalid value")
		os.Exit(1)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("usage: fortune | flolcat")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	print(output, coef)
}
