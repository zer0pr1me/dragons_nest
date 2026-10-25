package emu6502

// lda executes LDA (LoaD Accumulator with memory)
func (cpu *CPU6502) lda(arg uint8) {
	cpu.PC += 1
	cpu.A = arg
}

// ldx executes LDX (LoaD index X with memory)
func (cpu *CPU6502) ldx(arg uint8) {
	cpu.PC += 1
	cpu.X = arg
}

// ldy executes LDY (LoaD index Y with memory)
func (cpu *CPU6502) ldy(arg uint8) {
	cpu.PC += 1
	cpu.Y = arg
}

// sta executes STA (STore Accumulator in memory). Argument is presumed to
// be value stored already in the address, so that it is ignored
func (cpu *CPU6502) sta(_ uint8) uint8 {
	cpu.PC += 1
	return cpu.A
}

// stx executes STX (STore X in memory). Argument is presumed to be value
// stored already in the address, so that it is ignored
func (cpu *CPU6502) stx(_ uint8) uint8 {
	cpu.PC += 1
	return cpu.X
}

// sty executes STY (STore Y in memory). Argument is presumed to be value
// stored already in the address, so that it is ignored
func (cpu *CPU6502) sty(_ uint8) uint8 {
	cpu.PC += 1
	return cpu.Y
}

// tax executes TAX (Transfer Accumulator to X) operation
func (cpu *CPU6502) tax() {
	cpu.PC += 1
	cpu.X = cpu.A
}

// tay executes TAY (Transfer Accumulator to Y) operation
func (cpu *CPU6502) tay() {
	cpu.PC += 1
	cpu.Y = cpu.A
}

// txa executes TXA (Transfer X to Accumulator) operation
func (cpu *CPU6502) txa() {
	cpu.PC += 1
	cpu.A = cpu.X
}

// tya executes TYA (Transfer Y to Accumulator) operation
func (cpu *CPU6502) tya() {
	cpu.PC += 1
	cpu.A = cpu.Y
}
