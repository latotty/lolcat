package main

import (
	"flag"
	"log"

	"github.com/latotty/lolcat/modules"
)

type inputTypeEnum int

const (
	inputTypeStdin inputTypeEnum = iota
	inputTypeFile
)

type configStruct struct {
	wrapWidth int
	inputType inputTypeEnum
	filePath  string
}

func main() {
	config := configStruct{}

	parseArgs(&config)

	start(&config)
}

func parseArgs(config *configStruct) {
	flag.IntVar(&config.wrapWidth, "w", 0, "wrap width.  defaults to 0 / off.")
	flag.Parse()

	args := flag.Args()

	switch len(args) {
	case 0:
		config.inputType = inputTypeStdin
	case 1:
		config.inputType = inputTypeFile
		config.filePath = args[0]
	default:
		log.Fatal("Invalid arguments length")
	}
}

func start(config *configStruct) {
	// Create components, connecting the channels
	printer := modules.Printer{}
	wrapper := modules.Wrapper{Width: config.wrapWidth}

	var input interface {
		Start()
		OutChan() chan []byte
	}

	switch config.inputType {
	case inputTypeFile:
		fileReader := modules.FileReader{}
		fileReader.InFilePath = make(chan string, 1) // buffered so the next insert won't block
		fileReader.InFilePath <- config.filePath
		input = &fileReader
	default:
		input = &modules.StdInReader{}
	}

	// Connect them
	wrapper.In = input.OutChan()
	printer.In = wrapper.OutChan()

	// Start them
	input.Start()
	wrapper.Start()
	printer.Start()

	// Wait until the end
	for range printer.LineNotifChan() {
	}
}
