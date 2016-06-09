package main

// Wrapper wraps the stream if the line is too long
type Wrapper struct {
	In    chan []byte
	Out   chan []byte
	Width int
}

// OutChan inits and returns the Wrapper's Out channel
func (wrapper *Wrapper) OutChan() chan []byte {
	wrapper.Out = make(chan []byte, 16)
	return wrapper.Out
}

// Init inits the wrapper
func (wrapper *Wrapper) Init(width int) {
	wrapper.Width = width
	go func() {
		for line := range wrapper.In {
			if width != 0 {
				for len(line) > width {
					wrapper.Out <- append([]byte(nil), line[:width]...)
					line = line[width:]
				}
			}
			wrapper.Out <- append([]byte(nil), line...)
		}
		close(wrapper.Out)
	}()
}
