package producer

import "time"

// -- Producer simulates an external library that invokes the
// registered callback when it has new data for us once per 100ms.
type Producer struct {
	CallbackFunc func(event int)
}

func (p Producer) Start() {

	for {
		p.CallbackFunc(eventIndex)
		eventIndex++
		time.Sleep(time.Millisecond * 100)
	}
}
