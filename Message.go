package cansim

import "fmt"

type Message struct {
	Id uint32
	Data []byte
}

func (msg Message) String() string {
	s := fmt.Sprintf("0x%08x [", msg.Id)
	for _,data_byte := range msg.Data {
		s += fmt.Sprintf("%02x ", data_byte)
	}
	return s + "]"
}
