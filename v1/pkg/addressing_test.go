package emu6502

import "testing"

func TestRelative(t *testing.T) {
	cpu := setupCPU()
	cpu.PC = 0x608
	BulkWrite(cpu.Memory, cpu.PC, 0xd0, 0xf8)
	address := cpu.relative()

	assertAddressEquals(t, address, 0x602)
	cpu.assertPCEquals(t, 0x609)
}
