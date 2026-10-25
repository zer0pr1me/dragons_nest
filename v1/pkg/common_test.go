package emu6502

import "testing"

func setupCPU() *CPU6502 {
	cpu := NewCPU6502()
	cpu.PC = 0x600
	return cpu
}

func assertAddressEquals(t *testing.T, actual uint16, expected uint16) {
	if actual != expected {
		t.Errorf("Actual address %x != Expected address %x", actual, expected)
	}
}

func (cpu *CPU6502) assertAEquals(t *testing.T, value uint8) {
	if cpu.A != value {
		t.Errorf("Value in accumulator is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.A)
	}
}

func (cpu *CPU6502) assertXEquals(t *testing.T, value uint8) {
	if cpu.X != value {
		t.Errorf("Value in X register is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.X)
	}
}

func (cpu *CPU6502) assertYEquals(t *testing.T, value uint8) {
	if cpu.Y != value {
		t.Errorf("Value in Y register is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.Y)
	}
}

func (cpu *CPU6502) assertSPEquals(t *testing.T, value uint8) {
	if cpu.SP != value {
		t.Errorf("Value in SP register is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.SP)
	}
}

func (cpu *CPU6502) assertPCEquals(t *testing.T, value uint16) {
	if cpu.PC != value {
		t.Errorf("Value in program counter is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.PC)
	}
}

func (cpu *CPU6502) assertCarryEquals(t *testing.T, value bool) {
	if cpu.C != value {
		t.Errorf("Carry flag is expected to be equal to %t. Actual: %t", value, cpu.C)
	}
}

func (cpu *CPU6502) assertNegativeEquals(t *testing.T, value bool) {
	if cpu.N != value {
		t.Errorf("Negative flag is expected to be equal to %t. Actual: %t", value, cpu.N)
	}
}

func (cpu *CPU6502) assertZeroEquals(t *testing.T, value bool) {
	if cpu.Z != value {
		t.Errorf("Zero flag is expected to be equal to %t. Actual: %t", value, cpu.Z)
	}
}

func (cpu *CPU6502) assertMemoryEquals(t *testing.T, start uint16, bytes ...byte) {
	for i, expected := range bytes {
		actual := cpu.Memory.Read(start + uint16(i))
		if actual != expected {
			t.Errorf("Value at 0x%X is expected to be equal to 0x%x. Actual: 0x%x", start+uint16(i), expected, actual)
		}
	}
}
