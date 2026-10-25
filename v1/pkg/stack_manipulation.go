package emu6502

const STACK_BOTTOM = 0x100

// pushToStack pushes a given value to stack
func (cpu *CPU6502) pushToStack(value uint8) {
	cpu.Memory.Write(STACK_BOTTOM+uint16(cpu.SP), value)
	cpu.SP--
}

// pullFromStack pulls a top value from stack
func (cpu *CPU6502) pullFromStack() uint8 {
	cpu.SP++
	return cpu.Memory.Read(STACK_BOTTOM + uint16(cpu.SP))
}

// pushToStack16 pushes two bytes to stack. Higher byte will be pushed first
func (cpu *CPU6502) pushToStack16(value uint16) {
	// Push PC High Byte onto Stack
	cpu.pushToStack(uint8((value >> 8) & 0xFF))

	// Push PC Low Byte onto Stack
	cpu.pushToStack(uint8(value & 0xFF))
}

// pullFromStack16 pull two bytes from stack. Lower byte then higher byte
func (cpu *CPU6502) pullFromStack16() uint16 {
	lower := cpu.pullFromStack()
	higher := cpu.pullFromStack()

	return uint16(higher)<<8 + uint16(lower)
}

// brk executes BRK instruction
func (cpu *CPU6502) brk() {
	// coauthored by chatGPT
	// Increment PC to point to the next instruction after BRK
	cpu.PC++

	// TODO:
	// cpu.pushToStack16(cpu.PC)
	// Set Break flag (B) and push status onto Stack
	// cpu.pushToStack(cpu.cpuStatusToByte() | 0x10)

	// Set Interrupt Disable (I) flag
	cpu.I = true

	// TODO: figure out what is should be done with this one
	// Load Interrupt Vector (0xFFFE/F) into PC for Interrupt Service Routine (ISR)
	// cpu.PC = uint16(cpu.Memory.Read(0xFFFF))<<8 | uint16(cpu.Memory.Read(0xFFFE))
}

// pha executes PHA instruction pushing accumulator register to stack
func (cpu *CPU6502) pha() {
	cpu.PC += 1
	cpu.pushToStack(cpu.A)
}

// php executes PHP instruction pushing status register to stack
func (cpu *CPU6502) php() {
	cpu.PC += 1
	cpu.pushToStack(cpu.cpuStatusToByte())
}

// pla executes PLA instruction pulling value from stack to accumulator
func (cpu *CPU6502) pla() {
	cpu.PC += 1
	cpu.A = cpu.pullFromStack()
}

// plp executes PLP instruction pulling new status register value from a stack
func (cpu *CPU6502) plp() {
	cpu.PC += 1
	cpu.setCPUStatus(cpu.pullFromStack())
}

// rti executes RTI (ReTurn from Interrupt) instruction pulling status register
// and pc register from stack
func (cpu *CPU6502) rti() {
	cpu.setCPUStatus(cpu.pullFromStack())

	cpu.PC = cpu.pullFromStack16()
}

// tsx executes TSX (Transfer Stack pointer to X) instruction
func (cpu *CPU6502) tsx() {
	cpu.PC += 1
	cpu.X = cpu.SP
}

// txs executes TXS (Transfer X to Stack pointer) instruction
func (cpu *CPU6502) txs() {
	cpu.PC += 1
	cpu.SP = cpu.X
}
