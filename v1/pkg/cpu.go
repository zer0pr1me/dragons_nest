package emu6502

type CPU6502 struct {
	// A is accumulator register
	A uint8
	// X register
	X uint8
	// Y register
	Y uint8

	// PC is program counter. It stores an address of the next instruction to be
	// executed
	PC uint16

	// SP is stack pointer. It stores an address of current top of the stack
	SP uint8

	// C is carry flag. It is set whenever there are overflow happening in
	// arithmetic operations
	C bool

	// Z is zero flag. It is set after result of any operation is zero and unset
	// when it is not
	Z bool

	// D is decimal flag. Is is used to set processor to BCD (Binary-Coder Decimal)
	// mode for arithmetic operations. It is not used in current implementation and
	// added just FYI
	D bool

	// I is interrupt disable flag. It enables and disables CPU's ability to respond
	// to maskable interrupts.
	I bool

	// B is break flag. Indicates that a software interrupt was executed
	B bool

	// U is an extra unused bit in 6502 status register. Can be used for some things,
	// but added here just FYI
	U bool

	// V is signed overflow flag. Set whenever a signed arithmetic overlow happens.
	V bool

	// N is negative flag. Set whenever result of operation can be interpreted as
	// negative in context of signed arithmetic operations (most significant bit is set)
	N bool

	// Memory represents a memory used for storing
	Memory Memory
}

func nop(cpu *CPU6502) {
	// do nothing, just increment program counter
	cpu.PC++
}

// NewCPU6502 creates and initializes new 6502 CPU emulator instance
func NewCPU6502() *CPU6502 {
	return &CPU6502{
		A:      0,
		X:      0,
		Y:      0,
		PC:     0,
		SP:     0xFF, // Initialize stack pointer to 0xFF (top of stack)
		C:      false,
		Z:      false,
		I:      false,
		D:      false,
		V:      false,
		N:      false,
		Memory: NewRawArrayMemory(),
	}
}

// ExecuteNext runs next instruction at which PC currently points
func (cpu *CPU6502) ExecuteNext() {
	opcode := cpu.Memory.Read(cpu.PC)

	opcodeFunctions[opcode](cpu)
}

// ExecuteN is convenience function that runs next N instructions starting with
// instruction at [PC]
func (cpu *CPU6502) ExecuteN(n int) {
	for i := 0; i <= n; i++ {
		cpu.ExecuteNext()
	}
}

// LoadAndExecute is convenience function that loads bytes to memory at address
// at which PC points and executes until PC exceeds the last of them (be careful
// with jumps)
func LoadAndExecute(cpu *CPU6502, bytecode []byte) error {
	err := BulkWrite(cpu.Memory, cpu.PC, bytecode...)
	if err != nil {
		return err
	}

	end := cpu.PC + uint16(len(bytecode))
	for cpu.PC < end {
		cpu.ExecuteNext()
	}

	return nil
}
