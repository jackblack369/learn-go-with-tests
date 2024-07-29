package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3
const countdownMax = 10

// Countdown prints a countdown from 3 to out.
func Countdown(out io.Writer) {
	//for i := countdownStart; i > 0; i-- {
	//	fmt.Fprintln(out, i)
	//}

	for i := countdownStart; i <= countdownMax; i++ {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
