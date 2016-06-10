/*
	Forked from https://github.com/samuell/glow
*/

package modules

import (
	"bufio"
	"log"
	"os"
)

// FileReader is a file reader stream
type FileReader struct {
	InFilePath chan string
	Out        chan []byte
}

func (fileReader *FileReader) OutChan() chan []byte {
	fileReader.Out = make(chan []byte, 16)
	return fileReader.Out
}

func (fileReader *FileReader) Start() {
	go func() {
		file, err := os.Open(<-fileReader.InFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			fileReader.Out <- append([]byte(nil), scan.Bytes()...)
		}
		if scan.Err() != nil {
			log.Fatal(scan.Err())
		}

		close(fileReader.Out)
	}()
}
