package cansim

import "time"

type Device struct {
	ch *Channel
	C chan Message
}

func (d *Device) SendMessage(msg Message) {
	d.ch.post(d, msg)
}

func (d *Device) ReadMessage() Message {
	return <- d.C
}

func (d *Device) SendCyclic(msg Message, interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for t := range ticker.C {
			_ = t
			d.SendMessage(msg)
		}
	}()
	return ticker
}
