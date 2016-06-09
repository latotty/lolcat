package main

import "flag"

func main() {
	var wrapWidth int
	flag.IntVar(&wrapWidth, "w", 0, "wrap width.  defaults to 0 / off.")
	flag.Parse()

	// Create components, connecting the channels
	stdInReader := new(StdInReader)
	printer := new(Printer)
	wrapper := new(Wrapper)

	// Connect them
	wrapper.In = stdInReader.OutChan()
	printer.In = wrapper.OutChan()

	// Start them
	stdInReader.Init()
	wrapper.Init(wrapWidth)
	printer.Init()

	// Wait until the end
	for range printer.LineNotifChan() {
	}
}
