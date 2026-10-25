package emu6502

// bcc executes BCC (Branch on Carry Clear) instruction to a given address
func (cpu *CPU6502) bcc(address uint16) {
	cpu.PC += 1

	if !cpu.C {
		cpu.PC = address
	}
}

// bcs executes BCS (Branch on Carry Set) instruction to a given address
func (cpu *CPU6502) bcs(address uint16) {
	cpu.PC += 1

	if cpu.C {
		cpu.PC = address
	}
}

// bmi executes BMI (Branch on Minus) instructions on a given address setting
// program counter if negative flag is set
func (cpu *CPU6502) bmi(address uint16) {
	cpu.PC += 1

	if cpu.N {
		cpu.PC = address
	}
}

// beq executes BEQ (Branch on EQual/Branch on zEro) instruction on a given address
// setting program counter if zero flag is set
func (cpu *CPU6502) beq(address uint16) {
	cpu.PC += 1

	if cpu.Z {
		cpu.PC = address
	}
}

// bne executes BNE (Branch on Not Equal/Branch on Not zEro) instruction on a given
// address setting program counter if zero flag is not set
func (cpu *CPU6502) bne(address uint16) {
	cpu.PC += 1
	if !cpu.Z {
		cpu.PC = address
	}
}

// bpl executes BPL instruction on a given address setting program counter to it
// if negative flag is not set.
func (cpu *CPU6502) bpl(address uint16) {
	cpu.PC += 1

	if !cpu.N {
		cpu.PC = address
	}
}

// bvc executes BVC (Branch on overflow(V) Clear) instructions on a given address
// setting program counter if overflow flag is clear
func (cpu *CPU6502) bvc(address uint16) {
	cpu.PC += 1

	if !cpu.V {
		cpu.PC = address
	}
}

// bvs executes BVS (Branch on overflow(V) Set) instruction on a given address
// setting program counter if overflow flag is set
func (cpu *CPU6502) bvs(address uint16) {
	cpu.PC += 1

	if cpu.V {
		cpu.PC = address
	}
}

// jmp executes JMP instruction jumping on a given address
func (cpu *CPU6502) jmp(address uint16) {
	cpu.PC += 1

	cpu.PC = address
}

// jsr executes JSR instruction pushing return address to stack and jumping on a
// given address
func (cpu *CPU6502) jsr(address uint16) {
	cpu.pushToStack16(cpu.PC)
	cpu.PC = address
}

// rts executes RTS (ReTurn from Subroutine) instruction pulling address from
// a stack and setting it to program counter
func (cpu *CPU6502) rts() {
	cpu.PC = cpu.pullFromStack16()
	cpu.PC += 1
}
