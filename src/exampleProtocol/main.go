package exampleProtocol

import (
	"encoding/binary"
)

const (
	MessageTypeJSON = iota
	MessageTypeText = 2
	MessageTypeXML = 3
)

/**
* 4 byte message
* 4 byte message length
* message body - dynamic
*/
func CreateMessage(mtype int, data string) []byte {
	buf := make([]byte, 4+4+len(data))

	//See: https://www.techtarget.com/searchnetworking/definition/big-endian-and-little-endian
	// Binary sort algorithm
	binary.LittleEndian.PutUint32(buf[0:], uint32(mtype))
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data)))
	copy(buf[8:], []byte(data))

	return buf
}

/**
* 4 byte message
*/
func ReadMessage(data []byte) (mtype, mlen uint32, msg string) {
	mtype = binary.LittleEndian.Uint32(data[0:])
	mlen = binary.LittleEndian.Uint32(data[4:])
	msg = string(data[8:])

	return
}