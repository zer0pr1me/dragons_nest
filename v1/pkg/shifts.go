package emu6502

// asl performs Arithmetic Shift Left operation on the provided value
func (cpu *CPU6502) asl(value uint8) uint8 {
	// coauthored with chatGPT
	// Set the Carry flag if the most significant bit (bit 7) is 1, otherwise clear it
	cpu.C = (value & 0x80) != 0

	// Perform the shift left operation
	shiftedValue := value << 1

	// Update the Zero flag if the result is zero
	cpu.Z = shiftedValue == 0

	// Update the Negative flag based on bit 7 of the shifted value
	cpu.N = (shiftedValue & 0x80) != 0

	// Update program counter
	cpu.PC += 1
	return shiftedValue
}

// lsr performs LSR (Logical Shift Right) operation on the provided value
func (cpu *CPU6502) lsr(value uint8) uint8 {
	cpu.C = (value & 0x01) != 0

	value >>= 1

	cpu.Z = value == 0

	// there is no way it can be 1, because higher value bit will always be
	// 0
	cpu.N = false

	cpu.PC += 1
	return value
}

// rol performs ROL operation rotating byte with a carry flag to left
// (i.e. C=1 |0|1|1|0|1|0|0|0 => C=0 |1|1|0|1|0|0|0|1)
func (cpu *CPU6502) rol(value uint8) uint8 {
	carry := cpu.C
	cpu.C = (value & 0x80) != 0
	value <<= 1
	if carry {
		value += 1
	}

	cpu.N = (value & 0x80) != 0
	cpu.Z = value == 0

	return value
}

// ror performs ROR (Rotate One bit Right) opetation rotating byte with a carry
// flag to right (i.e. C=1 |0|1|1|0|1|0|0|0 => C=0 |1|0|1|1|0|1|0|0)
func (cpu *CPU6502) ror(value uint8) uint8 {
	carry := cpu.C
	cpu.C = (value & 0x01) != 0
	value >>= 1
	if carry {
		value += 0x80
	}

	cpu.N = (value & 0x80) != 0
	cpu.Z = value == 0

	return value
}
