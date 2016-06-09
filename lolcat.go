package main

import (
	"flag"
	"log"
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
	printer := new(Printer)
	wrapper := new(Wrapper)

	var input interface {
		Init()
		OutChan() chan []byte
	}

	switch config.inputType {
	case inputTypeFile:
		fileReader := new(FileReader)
		fileReader.InFilePath = make(chan string, 1) // buffered so the insert wont block
		fileReader.InFilePath <- config.filePath
		input = fileReader
	default:
		input = new(StdInReader)
	}

	// Connect them
	wrapper.In = input.OutChan()
	printer.In = wrapper.OutChan()

	// Start them
	input.Init()
	wrapper.Init(config.wrapWidth)
	printer.Init()

	// Wait until the end
	for range printer.LineNotifChan() {
	}
}
