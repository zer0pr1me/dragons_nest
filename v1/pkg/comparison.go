package emu6502

// cmp performs CMP (CoMPare memore with accumulator) operation
func (cpu *CPU6502) cmp(value uint8) {
	cpu.PC += 1

	// Subtract the value from A
	result := cpu.A - value

	// Set or clear the Carry flag based on the result
	cpu.C = cpu.A >= value

	// Update zero and negative flags
	cpu.Z = result == 0
	cpu.N = (result & 0x80) != 0
}

// cpx performs CPX (ComPare memory and index X) operation
func (cpu *CPU6502) cpx(value uint8) {
	cpu.PC += 1

	// Subtract the value from X
	result := cpu.X - value

	// Set or clear the Carry flag based on the result
	cpu.C = cpu.X >= value

	// Update zero and negative flags
	cpu.Z = result == 0
	cpu.N = (result & 0x80) != 0
}

// cpy performs CPY (ComPare memory and index Y) operation
func (cpu *CPU6502) cpy(value uint8) {
	// authored by chatGPT
	cpu.PC += 1

	// Subtract the value from Y
	result := cpu.Y - value

	// Set or clear the Carry flag based on the result
	cpu.C = cpu.Y >= value

	// Update zero and negative flags
	cpu.Z = result == 0
	cpu.N = (result & 0x80) != 0
}
