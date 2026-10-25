package emu6502

import "errors"

// Memory represents memory with a [0x00000-0x10000] mapping used by 6502 CPU
type Memory interface {
	// Read gets the byte from memory using a given index
	Read(index uint16) uint8
	// Write sets the byte with a given index in memory with a given value
	Write(index uint16, value uint8)
}

type rawArrayMemory struct {
	array [0x10000]uint8
}

// NewRawArrayMemory creates new memory instance that is just 65536 byte array
// under the hood
func NewRawArrayMemory() Memory {
	return &rawArrayMemory{}
}

func (ram *rawArrayMemory) Read(index uint16) uint8 {
	return ram.array[index]
}

func (ram *rawArrayMemory) Write(index uint16, value uint8) {
	ram.array[index] = value
}

// BulkWrite is convenience function that allows to load in memory multiple
// bytes at once
func BulkWrite(memory Memory, start uint16, bytes ...uint8) error {
	if int(start)+len(bytes) > 0x10000 {
		return errors.New("Exceeding the memory maximum possible address of 0xffff")
	}

	for i, b := range bytes {
		memory.Write(start+uint16(i), b)
	}

	return nil
}
