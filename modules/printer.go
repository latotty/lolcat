/*
	Forked from https://github.com/samuell/glow
*/

package modules

import (
	"bufio"
	"fmt"
	"os"
)

// Printer prints
type Printer struct {
	In        chan []byte
	LineNotif chan int
}

// LineNotifChan notifies for every line it prints, and closed on finish
func (printer *Printer) LineNotifChan() chan int {
	printer.LineNotif = make(chan int)
	return printer.LineNotif
}

func (printer *Printer) Start() {
	go func() {
		w := bufio.NewWriter(os.Stdout)
		for line := range printer.In {
			if len(line) > 0 {
				fmt.Fprintln(w, string(line))
			}
			printer.LineNotif <- 1
		}
		w.Flush()
		close(printer.LineNotif)
	}()
}
