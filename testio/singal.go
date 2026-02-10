package testio

import (
	"fmt"
	"os"
)

var signal = make(chan struct{})

func Notify() {
	signal <- struct{}{}
}

func queueInput(in *os.File, input []any) {
	for _, line := range input {
		<-signal
		fmt.Fprintln(in, line)
	}
}
