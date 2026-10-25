package emu6502

// and executes AND instruction performing bitwise AND for A register and a
// given argument
func (cpu *CPU6502) and(arg uint8) {
	cpu.A &= arg

	cpu.Z = cpu.A == 0
	cpu.N = (cpu.A & 0x80) != 0

	cpu.PC += 1
}

// ora executes ORA instruction performing bitwise OR for A register and a
// given argument
func (cpu *CPU6502) ora(arg uint8) {
	// Performing the bitwise OR
	cpu.A |= arg

	// Setting zero flag if result is zero
	cpu.Z = cpu.A == 0

	// Setting negative flag if result can be interpreted as bitewise negative value
	// maximum bit is set to one
	cpu.N = (cpu.A & 0x80) != 0

	cpu.PC += 1
}

// eor executes EOR instruction performing bitwise exclusive OR for A register
// and a given argument
func (cpu *CPU6502) eor(arg uint8) {
	// authored by chatGPT
	cpu.A ^= arg

	// Update zero and negative flags as required
	cpu.Z = cpu.A == 0
	cpu.N = (cpu.A & 0x80) != 0

	cpu.PC += 1
}
