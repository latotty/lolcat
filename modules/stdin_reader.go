/*
	Forked from https://github.com/samuell/glow
*/

package modules

import (
	"bufio"
	"os"
)

type StdInReader struct {
	Out chan []byte
}

func (self *StdInReader) OutChan() chan []byte {
	self.Out = make(chan []byte, 16)
	return self.Out
}

func (self *StdInReader) Start() {
	go func() {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			self.Out <- append([]byte(nil), scan.Bytes()...)
		}
		close(self.Out)
	}()
}
