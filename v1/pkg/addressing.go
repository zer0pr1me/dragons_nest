package emu6502

// absoluteAddress returns next operation argument address using absolute
// addressing mode. Updates PC
func (cpu *CPU6502) absoluteAddress() uint16 {
	cpu.PC += 1
	lower := uint16(cpu.Memory.Read(cpu.PC))

	cpu.PC += 1
	higher := uint16(cpu.Memory.Read(cpu.PC))

	return higher<<8 | lower
}

// absoluteAddressX returns next operation argument address using absolute, X
// addressing mode, Updates PC
func (cpu *CPU6502) absoluteAddressX() uint16 {
	return cpu.absoluteAddress() + uint16(cpu.X)
}

// absoluteAddressY returns next operation argument address using absolute, Y
// addressing mode, Updates PC
func (cpu *CPU6502) absoluteAddressY() uint16 {
	return cpu.absoluteAddress() + uint16(cpu.Y)
}

// indirectAddress returns address for next operation using indirect addressing
// mode. Updates PC
func (cpu *CPU6502) indirectAddress() uint16 {
	initialAddress := cpu.absoluteAddress()
	lower := uint16(cpu.Memory.Read(initialAddress))
	higher := uint16(cpu.Memory.Read(initialAddress + 1))

	return higher<<8 | lower
}

// indexedIndirectAddress returns address for next operation using indexed
// indirect addressing mode (address), X. Updates PC
func (cpu *CPU6502) indexedIndirectAddress() uint16 {
	// Getting initial address using zeropage, X adrressing
	address := cpu.zeroPageXAddress()

	// Return indirect address
	return uint16(cpu.Memory.Read(address+1))<<8 + uint16(cpu.Memory.Read(address))
}

// indirectIndexedAddress returns address for next operation using indirect
// indexed addressing mode (address, Y). Updates PC
func (cpu *CPU6502) indirectIndexedAddress() uint16 {
	address := cpu.zeroPageAddress()
	address = uint16(cpu.Memory.Read(address+1))<<8 + uint16(cpu.Memory.Read(address))
	return address + uint16(cpu.Y)
}

// zeroPageAddress returns address for next operation using zero page addressing
// mode. Updates PC
func (cpu *CPU6502) zeroPageAddress() uint16 {
	cpu.PC += 1
	return uint16(cpu.Memory.Read(cpu.PC))
}

// zeroPageXAddress returns address for next operation using zero page + X addressing
// mode. Updates PC
func (cpu *CPU6502) zeroPageXAddress() uint16 {
	cpu.PC += 1
	return uint16(cpu.Memory.Read(cpu.PC) + cpu.X)
}

// immediate is used to get argument for operation using immediate addressing mode
// just getting value of the byte next to the opcode. Updates PC
func (cpu *CPU6502) immediate() uint8 {
	cpu.PC += 1

	return cpu.Memory.Read(cpu.PC)
}

// absolute is used to get argument for operation using absolute addressing mode
// getting value from 2-byte address. Updates PC
func (cpu *CPU6502) absolute() uint8 {
	return cpu.Memory.Read(cpu.absoluteAddress())
}

// absoluteX is used to get argument for operation using absolute addressing mode
// getting value from 2-byte address plus X register. Updates PC
func (cpu *CPU6502) absoluteX() uint8 {
	return cpu.Memory.Read(cpu.absoluteAddressX())
}

// absoluteY is used to get argument for operation using absolute addressing mode
// getting value from 2-byte address plus value in Y register. Updates PC
func (cpu *CPU6502) absoluteY() uint8 {
	return cpu.Memory.Read(cpu.absoluteAddressY())
}

// indexedIndirect is used to get argument for operation using indexedIndirect ($address, X)
// addressing. Updates PC
func (cpu *CPU6502) indexedIndirect() uint8 {
	return cpu.Memory.Read(cpu.indexedIndirectAddress())
}

// indirectIndexed is used to get argument for operation using indirect indexed ($address), Y
// addressing. Updates PC
func (cpu *CPU6502) indirectIndexed() uint8 {
	return cpu.Memory.Read(cpu.indirectIndexedAddress())
}

// relative returns relative to current PC address that can be used by controlling
// instructions for jumps. Updates PC
func (cpu *CPU6502) relative() uint16 {
	// 0xff = -1 = 0b1111_1111
	// 0xfe = -2 = 0b1111_1110
	// 0xfd = -3 = 0b1111_1101
	// 0xfc = -4 = 0b1111_1100
	cpu.PC += 1

	// +2 because we start to count after the instruction ([Instruction] [Relative argument])
	// we assume all controling instructions will take 2 byte. And one byte we already counted
	offset := int16(int8(cpu.Memory.Read(cpu.PC))) + 1

	value := uint16(int16(cpu.PC) + offset)

	return value
}

// zeroPage is used to get argument for operation using zero page addressing mode.
// Updates PC
func (cpu *CPU6502) zeroPage() uint8 {
	return cpu.Memory.Read(cpu.zeroPageAddress())
}

// zeroPageX is used to get argument for operation using (zero page + X) addressing mode.
// Updates PC
func (cpu *CPU6502) zeroPageX() uint8 {
	return cpu.Memory.Read(cpu.zeroPageXAddress())
}

// zeroPageY is used to get argument for operation using (zero page + Y) addressing mode.
// Updates PC
func (cpu *CPU6502) zeroPageY() uint8 {
	cpu.PC += 1

	// Getting  zero-page address from argument
	return cpu.Memory.Read(cpu.zeroPageAddress() + uint16(cpu.Y))
}

// executeWithAccumulator is used for operations for performing operations and than
// storing result using accumulator register
func (cpu *CPU6502) executeWithAccumulator(operation func(value uint8) uint8) {
	cpu.A = operation(cpu.A)
}

// executeWithAbsolute is used for performing operations and then storing result
// using zeropage address. Updates PC
func (cpu *CPU6502) executeWithAbsolute(operation func(value uint8) uint8) {
	address := cpu.absoluteAddress()
	cpu.Memory.Write(address, operation(cpu.Memory.Read(address)))
}

// executeWithAbsoluteX is used for performing operations and then storing result
// using zeropage address plus X register. Updates PC
func (cpu *CPU6502) executeWithAbsoluteX(operation func(value uint8) uint8) {
	address := cpu.absoluteAddressX()
	cpu.Memory.Write(address, operation(cpu.Memory.Read(address)))
}

// executeWithAbsoluteY is used for performing operations and then storing result
// using zeropage address plus Y register. Updates PC
func (cpu *CPU6502) executeWithAbsoluteY(operation func(value uint8) uint8) {
	address := cpu.absoluteAddressY()
	cpu.Memory.Write(address, operation(cpu.Memory.Read(address)))
}

// executeWithIndexedIndirect is used for performing operations and then storing result
// using indexed indirect addressing mode (address, X). Update PC
func (cpu *CPU6502) executeWithIndexedIndirect(operation func(uint8) uint8) {
	address := cpu.indexedIndirectAddress()
	cpu.Memory.Write(address, operation(cpu.Memory.Read(address)))
}

// exectureWithIndirectIndexed is used for performing operations and then storing result
// using indirect indexed addressing mode (address), Y. Update PC
func (cpu *CPU6502) executeWithIndirectIndexed(operation func(uint8) uint8) {
	address := cpu.indirectIndexedAddress()
	cpu.Memory.Write(address, operation(cpu.Memory.Read(address)))
}

// executeWithZeroPage is used for performing operations and than storing result
// using zeropage address. Updates PC
func (cpu *CPU6502) executeWithZeroPage(operation func(value uint8) uint8) {
	address := cpu.zeroPageAddress()
	cpu.Memory.Write(address, operation(cpu.Memory.Read(address)))
}

// executeWithZeroPageX is used for performing operations and than storing result
// using zeropage, X address. Updates PC
func (cpu *CPU6502) executeWithZeroPageX(operation func(value uint8) uint8) {
	address := cpu.zeroPageXAddress()
	cpu.Memory.Write(address, operation(cpu.Memory.Read(address)))
}
