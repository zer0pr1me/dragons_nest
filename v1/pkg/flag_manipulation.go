package emu6502

// bit executes BIT instruction
func (cpu *CPU6502) bit(value uint8) {
	cpu.PC += 1
	cpu.N = value&(0x01<<7) == 1
	cpu.V = value&(0x01<<6) == 1

	cpu.Z = (value & cpu.A) == 0
}

// clc executes CLC instruction clearing carry flag
func (cpu *CPU6502) clc() {
	cpu.PC += 1
	cpu.C = false
}

// cli executes CLI (CLear Interrupt disable bit) instruction
func (cpu *CPU6502) cli() {
	cpu.PC += 1
	cpu.I = false
}

// cld executes CLD (Clear Decimal) instruction
func (cpu *CPU6502) cld() {
	cpu.PC += 1
	cpu.D = true
}

// cli executes CLV (CLear oVerflow flag)
func (cpu *CPU6502) clv() {
	cpu.PC += 1
	cpu.V = false
}

// sec executes SEC instruction setting carry flag
func (cpu *CPU6502) sec() {
	cpu.PC += 1
	cpu.C = true
}

// sei executes SEI (SEt Interrupt disable) instruction
func (cpu *CPU6502) sei() {
	cpu.PC += 1
	cpu.I = true
}

// sed executes SED (SEt Desimal flag) instruction
func (cpu *CPU6502) sed() {
	cpu.PC += 1
	cpu.D = true
}

// cpuStatusToByte converts CPU status to byte format
func (cpu *CPU6502) cpuStatusToByte() uint8 {
	status := uint8(0x20) // Bit 5 is always set for the 6502
	if cpu.N {
		status |= 0x80
	}
	if cpu.V {
		status |= 0x40
	}
	if cpu.U {
		status |= 0x20
	}
	if cpu.B {
		status |= 0x10
	}
	if cpu.D {
		status |= 0x08
	}
	if cpu.I {
		status |= 0x04
	}
	if cpu.Z {
		status |= 0x02
	}
	if cpu.C {
		status |= 0x01
	}
	return status
}

// setCPUStatus sets CPU status according to bit values in a given byte
func (cpu *CPU6502) setCPUStatus(status uint8) {
	cpu.N = (status & 0x80) != 0
	cpu.V = (status & 0x40) != 0
	cpu.U = (status & 0x20) != 0
	cpu.B = (status & 0x10) != 0
	cpu.D = (status & 0x08) != 0
	cpu.I = (status & 0x04) != 0
	cpu.Z = (status & 0x02) != 0
	cpu.C = (status & 0x01) != 0
}
