package cansim
import "fmt"

type broadcast struct {
	msg Message
	sender *Device
}

type Channel struct {
	Id uint
	in chan broadcast
	addDevice chan *Device
	devices []*Device
}

func NewChannel(id uint) *Channel {
	var ch Channel
	ch.Id = id
	ch.in = make(chan broadcast)
	ch.addDevice = make (chan *Device)

	go ch.dispatcher()
	return &ch
}

func (ch Channel) CreateDevice() *Device {
	new_device := Device{&ch, make(chan Message, 20)}
	ch.addDevice <- &new_device
	return &new_device
}

func (ch Channel) post(sender *Device, msg Message) {
	ch.in <- broadcast{msg, sender}
}

func (ch Channel) dispatcher() {
	for {
		select {
		case device := <- ch.addDevice:
			ch.devices = append(ch.devices, device)
		case bc := <- ch.in:
			ch.broadcast(bc)
		}
	}
}

func (ch Channel) broadcast(bc broadcast) {
	for _,device := range ch.devices {
		if device != bc.sender  {
			select {
			case device.C <- bc.msg:
			default:
			}
		}
	}
}

func (ch *Channel) ConnectTo(ch2 *Channel) {
	dev1 := ch.CreateDevice()
	dev2 := ch2.CreateDevice()

	go func() {
		for {
			select {
			case msg := <-dev1.C:
				dev2.SendMessage(msg)
			case msg := <-dev2.C:
				dev1.SendMessage(msg)
			}
		}
	}()
}

func (ch *Channel) DumpToConsole(prefix string) {
	dev := ch.CreateDevice()

	go func() {
		for msg := range(dev.C) {
			fmt.Println(prefix, msg)
		}
	}()

}
