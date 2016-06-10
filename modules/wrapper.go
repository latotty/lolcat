package modules

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
func (wrapper *Wrapper) Start() {
	go func() {
		for line := range wrapper.In {
			if wrapper.Width != 0 {
				for len(line) > wrapper.Width {
					wrapper.Out <- append([]byte(nil), line[:wrapper.Width]...)
					line = line[wrapper.Width:]
				}
			}
			wrapper.Out <- append([]byte(nil), line...)
		}
		close(wrapper.Out)
	}()
}
