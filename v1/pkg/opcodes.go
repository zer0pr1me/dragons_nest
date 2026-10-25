package emu6502

var opcodeFunctions [0x100]func(cpu *CPU6502) = [0x100]func(cpu *CPU6502){
	// 0x00 BRK
	func(cpu *CPU6502) {
		cpu.brk()
	},
	// 0x01 ORA, (indirect, X)
	func(cpu *CPU6502) {
		cpu.ora(cpu.indexedIndirect())
	},
	// 0x02 is not defined, assign NOP function
	nop,
	// 0x03 is not defined, assign NOP function
	nop,
	// 0x04 is not defined, assign NOP function
	nop,
	// 0x05 ORA, zero page
	func(cpu *CPU6502) {
		cpu.ora(cpu.zeroPage())
	},
	// 0x06 ASL, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.asl)
	},
	// 0x07 is not defined, assign NOP function
	nop,
	// 0x08 PHP
	func(cpu *CPU6502) {
		cpu.php()
	},
	// 0x09 ORA, immediate
	func(cpu *CPU6502) {
		cpu.ora(cpu.immediate())
	},
	// 0x0a ASL, accumulator
	func(cpu *CPU6502) {
		cpu.executeWithAccumulator(cpu.asl)
	},
	// 0x0b is not defined, assign NOP function
	nop,
	// 0x0c is not defined, assign NOP function
	nop,
	// 0x0d ORA, absolute
	func(cpu *CPU6502) {
		cpu.ora(cpu.absolute())
	},
	// 0x0e ASL, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.asl)
	},
	// 0x0f is not defined, assign NOP function
	nop,

	// 0x10 BPL, relative
	func(cpu *CPU6502) {
		cpu.bpl(cpu.relative())
	},
	// 0x11 ORA, inderect indexed: (address), Y
	func(cpu *CPU6502) {
		cpu.ora(cpu.indirectIndexed())
	},
	// 0x12 is not defined, assign NOP function
	nop,
	// 0x13 is not defined, assign NOP function
	nop,
	// 0x14 is not defined, assign NOP function
	nop,
	// 0x15 ORA, zero-page, X
	func(cpu *CPU6502) {
		cpu.ora(cpu.zeroPageX())
	},
	// 0x16 ASL, zero-page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.asl)
	},
	// 0x17 is not defined, assign NOP function
	nop,
	// 0x18 CLC
	func(cpu *CPU6502) {
		cpu.clc()
	},
	// 0x19 ORA, absolute, Y
	func(cpu *CPU6502) {
		cpu.ora(cpu.absoluteY())
	},
	// 0x1a is not defined, assign NOP function
	nop,
	// 0x1b is not defined, assign NOP function
	nop,
	// 0x1c is not defined, assign NOP function
	nop,
	// 0x1d ORA, absolute, X
	func(cpu *CPU6502) {
		cpu.ora(cpu.absoluteX())
	},
	// 0x1e ASL, absolute, X
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteY(cpu.asl)
	},
	// 0x1f is not defined, assign NOP function
	nop,
	// 0x20 JSR, absolute
	func(cpu *CPU6502) {
		cpu.jsr(cpu.absoluteAddress())
	},
	// 0x21 AND, (indirect, X)
	func(cpu *CPU6502) {
		cpu.and(cpu.indexedIndirect())
	},
	// 0x22 is not defined, assign NOP function
	nop,
	// 0x23 is not defined, assign NOP function
	nop,
	// 0x24 BIT, zero page
	func(cpu *CPU6502) {
		cpu.bit(cpu.zeroPage())
	},
	// 0x25 AND, zero page
	func(cpu *CPU6502) {
		cpu.and(cpu.zeroPage())
	},
	// 0x26 ROL, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.rol)
	},
	// 0x27 is not defined, assign NOP function
	nop,
	// 0x28 PLP
	func(cpu *CPU6502) {
		cpu.plp()
	},
	// 0x29 AND, immediate
	func(cpu *CPU6502) {
		cpu.and(cpu.immediate())
	},
	// 0x2a ROL, accumulator
	func(cpu *CPU6502) {
		cpu.executeWithAccumulator(cpu.rol)
	},
	// 0x2b is not defined, assign NOP function
	nop,
	// 0x2c BIT, absolute
	func(cpu *CPU6502) {
		cpu.bit(cpu.absolute())
	},
	// 0x2d AND, absolute
	func(cpu *CPU6502) {
		cpu.and(cpu.absolute())
	},
	// 0x2e ROL, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.rol)
	},
	// 0x2f is not defined, assign NOP function
	nop,
	// 0x30 BMI, relative
	func(cpu *CPU6502) {
		cpu.bmi(cpu.relative())
	},
	// 0x31 AND, indirect indexed, (address), Y
	func(cpu *CPU6502) {
		cpu.and(cpu.indirectIndexed())
	},
	// 0x32 is not defined, assign NOP function
	nop,
	// 0x33 is not defined, assign NOP function
	nop,
	// 0x34 is not defined, assign NOP function
	nop,
	// 0x35 AND, zero page, X
	func(cpu *CPU6502) {
		cpu.and(cpu.zeroPageX())
	},
	// 0x36 ROL, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.rol)
	},
	// 0x37 is not defined, assign NOP function
	nop,
	// 0x38 SEC
	func(cpu *CPU6502) {
		cpu.sec()
	},
	// 0x39 AND, absolute, Y
	func(cpu *CPU6502) {
		cpu.and(cpu.absoluteY())
	},
	// 0x3a is not defined, assign NOP function
	nop,
	// 0x3b is not defined, assign NOP function
	nop,
	// 0x3c is not defined, assign NOP function
	nop,
	// 0x3d AND, absolute, X
	func(cpu *CPU6502) {
		cpu.and(cpu.absoluteX())
	},
	// 0x3e ROL, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteX(cpu.rol)
	},
	// 0x3f is not defined, assign NOP function
	nop,

	// 0x40 RTI
	func(cpu *CPU6502) {
		cpu.rti()
	},
	// 0x41 EOR indexed indirect, (address, X)
	func(cpu *CPU6502) {
		cpu.eor(cpu.indexedIndirect())
	},
	// 0x42 is not defined, assign NOP function
	nop,
	// 0x43 is not defined, assign NOP function
	nop,
	// 0x44 is not defined, assign NOP function
	nop,
	// 0x45 EOR, zero page
	func(cpu *CPU6502) {
		cpu.eor(cpu.zeroPage())
	},
	// 0x46 LSR, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.lsr)
	},
	// 0x47 is not defined, assign NOP function
	nop,
	// 0x48 PHA
	func(cpu *CPU6502) {
		cpu.pha()
	},
	// 0x49 EOR, immediate
	func(cpu *CPU6502) {
		cpu.eor(cpu.immediate())
	},
	// 0x4a LSR, accumulator
	func(cpu *CPU6502) {
		cpu.executeWithAccumulator(cpu.lsr)
	},
	// 0x4b is not defined, assign NOP function
	nop,
	// 0x4c JMP, absolute
	func(cpu *CPU6502) {
		cpu.jmp(cpu.absoluteAddress())
	},
	// 0x4d EOR, absolute
	func(cpu *CPU6502) {
		cpu.eor(cpu.absolute())
	},
	// 0x4e LSR, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.lsr)
	},
	// 0x4f is not defined, assign NOP function
	nop,

	// 0x50 BVC, relative
	func(cpu *CPU6502) {
		cpu.bvc(cpu.relative())
	},
	// 0x51 EOR, indirect indexed, (address), Y
	func(cpu *CPU6502) {
		cpu.eor(cpu.indirectIndexed())
	},
	// 0x52 is not defined, assign NOP function
	nop,
	// 0x53 is not defined, assign NOP function
	nop,
	// 0x54 is not defined, assign NOP function
	nop,
	// 0x55 EOR, zero page, X
	func(cpu *CPU6502) {
		cpu.eor(cpu.zeroPageX())
	},
	// 0x56 LSR, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.lsr)
	},
	// 0x57 is not defined, assign NOP function
	nop,
	// 0x58 CLI
	func(cpu *CPU6502) {
		cpu.cli()
	},
	// 0x59 EOR, absolute, Y
	func(cpu *CPU6502) {
		cpu.eor(cpu.absoluteY())
	},
	// 0x5a is not defined, assign NOP function
	nop,
	// 0x5b is not defined, assign NOP function
	nop,
	// 0x5c is not defined, assign NOP function
	nop,
	// 0x5d EOR, absolute, X
	func(cpu *CPU6502) {
		cpu.eor(cpu.absoluteX())
	},
	// 0x5e LSR, absolute, X
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteX(cpu.lsr)
	},
	// 0x5f is not defined, assign NOP function
	nop,

	// 0x60 RTS
	func(cpu *CPU6502) {
		cpu.rts()
	},
	// 0x61 ADC, indexed indirect, (address, X)
	func(cpu *CPU6502) {
		cpu.adc(cpu.indexedIndirect())
	},
	// 0x62 is not defined, assign NOP function
	nop,
	// 0x63 is not defined, assign NOP function
	nop,
	// 0x64 is not defined, assign NOP function
	nop,
	// 0x65 ADC, zero page
	func(cpu *CPU6502) {
		cpu.adc(cpu.zeroPage())
	},
	// 0x66 ROR, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.ror)
	},
	// 0x67 is not defined, assign NOP function
	nop,
	// 0x68 PLA
	func(cpu *CPU6502) {
		cpu.pla()
	},
	// 0x69 ADC, immediate
	func(cpu *CPU6502) {
		cpu.adc(cpu.immediate())
	},
	// 0x6a ROR, accumulator
	func(cpu *CPU6502) {
		cpu.executeWithAccumulator(cpu.ror)
	},
	// 0x6b is not defined, assign NOP function
	nop,
	// 0x6c JMP, indirect
	func(cpu *CPU6502) {
		cpu.jmp(cpu.indirectAddress())
	},
	// 0x6d ADC, absolute
	func(cpu *CPU6502) {
		cpu.adc(cpu.absolute())
	},
	// 0x6e ROR, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.ror)
	},
	// 0x6f is not defined, assign NOP function
	nop,

	// 0x70, BVS, relative
	func(cpu *CPU6502) {
		cpu.bvs(cpu.relative())
	},
	// 0x71, ADC, indirect indexed, (address), Y
	func(cpu *CPU6502) {
		cpu.adc(cpu.indirectIndexed())
	},
	// 0x72 is not defined, assign NOP function
	nop,
	// 0x73 is not defined, assign NOP function
	nop,
	// 0x74 is not defined, assign NOP function
	nop,
	// 0x75 ADC, zero page, X
	func(cpu *CPU6502) {
		cpu.adc(cpu.zeroPageX())
	},
	// 0x76 ROR, zero page, X
	func(cpu *CPU6502) {
		cpu.ror(cpu.zeroPageX())
	},
	// 0x77 is not defined, assign NOP function
	nop,
	// 0x78 SEI
	func(cpu *CPU6502) {
		cpu.sei()
	},
	// 0x79 ADC, absolute, Y
	func(cpu *CPU6502) {
		cpu.adc(cpu.absoluteY())
	},
	// 0x7a is not defined, assign NOP function
	nop,
	// 0x7b is not defined, assign NOP function
	nop,
	// 0x7c is not defined, assign NOP function
	nop,
	// 0x7d ADC, absolute, X
	func(cpu *CPU6502) {
		cpu.adc(cpu.absoluteX())
	},
	// 0x7e ROR, absolute, X
	func(cpu *CPU6502) {
		cpu.ror(cpu.absoluteX())
	},
	// 0x7f is not defined, assign NOP function
	nop,

	// 0x80 is not defined, assign NOP function
	nop,
	// 0x81 STA, indexed indirect, (address, X)
	func(cpu *CPU6502) {
		cpu.executeWithIndexedIndirect(cpu.sta)
	},
	// 0x82 is not defined, assign NOP function
	nop,
	// 0x83 is not defined, assign NOP function
	nop,
	// 0x84 STY, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.sty)
	},
	// 0x85 STA, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.sta)
	},
	// 0x86 STX, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.stx)
	},
	// 0x87 is not defined, assign NOP function
	nop,
	// 0x88 DEY
	func(cpu *CPU6502) {
		cpu.dey()
	},
	// 0x89 is not defined, assign NOP function
	nop,
	// 0x8a TXA
	func(cpu *CPU6502) {
		cpu.txa()
	},
	// 0x8b is not defined, assign NOP function
	nop,
	// 0x8c STY, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.sty)
	},
	// 0x8d STA, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.sta)
	},
	// 0x8e STX, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.stx)
	},
	// 0x8f is not defined, assign NOP function
	nop,

	// 0x90 BCC, relative
	func(cpu *CPU6502) {
		cpu.bcc(cpu.relative())
	},
	// 0x91 STA, indirect indexed (address, Y)
	func(cpu *CPU6502) {
		cpu.executeWithIndirectIndexed(cpu.sta)
	},
	// 0x92 is not defined, assign NOP function
	nop,
	// 0x93 is not defined, assign NOP function
	nop,
	// 0x94 STY, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.sty)
	},
	// 0x95 STA, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.sta)
	},
	// 0x96 STX, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.stx)
	},
	// 0x97 is not defined, assign NOP function
	nop,
	// 0x98 TYA
	func(cpu *CPU6502) {
		cpu.tya()
	},
	// 0x99 STA, absolute, Y
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteY(cpu.sta)
	},
	// 0x9a TXS
	func(cpu *CPU6502) {
		cpu.txs()
	},
	// 0x9b is not defined, assign NOP function
	nop,
	// 0x9c is not defined, assign NOP function
	nop,
	// 0x9d STA, absolute, X
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteX(cpu.sta)
	},
	// 0x9e is not defined, assign NOP function
	nop,
	// 0x9f is not defined, assign NOP function
	nop,

	// 0xa0 LDY, immediate
	func(cpu *CPU6502) {
		cpu.ldy(cpu.immediate())
	},
	// 0xa1 LDA, indexed indirect (address, X)
	func(cpu *CPU6502) {
		cpu.lda(cpu.indexedIndirect())
	},
	// 0xa2 LDX, immediate
	func(cpu *CPU6502) {
		cpu.ldx(cpu.immediate())
	},
	// 0xa3 is not defined, assign NOP function
	nop,
	// 0xa4 LDY, zero page
	func(cpu *CPU6502) {
		cpu.ldy(cpu.zeroPage())
	},
	// 0xa5 LDA, zero page
	func(cpu *CPU6502) {
		cpu.lda(cpu.zeroPage())
	},
	// 0xa6 LDX, zero page
	func(cpu *CPU6502) {
		cpu.ldx(cpu.zeroPage())
	},
	// 0xa7 is not defined, assign NOP function
	nop,
	// 0xa8 TAY
	func(cpu *CPU6502) {
		cpu.tay()
	},
	// 0xa9 LDA, immediate
	func(cpu *CPU6502) {
		cpu.lda(cpu.immediate())
	},
	// 0xaa TAX
	func(cpu *CPU6502) {
		cpu.tax()
	},
	// 0xab is not defined, assign NOP function
	nop,
	// 0xac LDY, absolute
	func(cpu *CPU6502) {
		cpu.ldy(cpu.absolute())
	},
	// 0xad LDA, absolute
	func(cpu *CPU6502) {
		cpu.lda(cpu.absolute())
	},
	// 0xae LDX, absolute
	func(cpu *CPU6502) {
		cpu.ldx(cpu.absolute())
	},
	// 0xaf is not defined, assign NOP function
	nop,

	// 0xb0 BCS, relative
	func(cpu *CPU6502) {
		cpu.bcs(cpu.relative())
	},
	// 0xb1 LDA, indirect indexed, (address), Y
	func(cpu *CPU6502) {
		cpu.lda(cpu.indirectIndexed())
	},
	// 0xb2 is not defined, assign NOP function
	nop,
	// 0xb3 is not defined, assign NOP function
	nop,
	// 0xb4 LDY, zero page, X
	func(cpu *CPU6502) {
		cpu.ldy(cpu.zeroPageX())
	},
	// 0xb5 LDA, zero page, X
	func(cpu *CPU6502) {
		cpu.lda(cpu.zeroPageX())
	},
	// 0xb6 LDX, zero page, Y
	func(cpu *CPU6502) {
		cpu.ldx(cpu.zeroPageY())
	},
	// 0xb7 is not defined, assign NOP function
	nop,
	// 0xb8 CLV
	func(cpu *CPU6502) {
		cpu.clv()
	},
	// 0xb9 LDA, absolute, Y
	func(cpu *CPU6502) {
		cpu.lda(cpu.absoluteY())
	},
	// 0xba TSX
	func(cpu *CPU6502) {
		cpu.tsx()
	},
	// 0xbb is not defined, assign NOP function
	nop,
	// 0xbc LDY, absolute, X
	func(cpu *CPU6502) {
		cpu.ldy(cpu.absoluteX())
	},
	// 0xbd LDA, absolute, X
	func(cpu *CPU6502) {
		cpu.lda(cpu.absoluteX())
	},
	// 0xbe LDX, absolute, Y
	func(cpu *CPU6502) {
		cpu.ldx(cpu.absoluteY())
	},
	// 0xbf is not defined, assign NOP function
	nop,

	// 0xc0 CPY, immediate
	func(cpu *CPU6502) {
		cpu.cpy(cpu.immediate())
	},
	// 0xc1 CMP, indexed indirect, (X, address)
	func(cpu *CPU6502) {
		cpu.cmp(cpu.indexedIndirect())
	},
	// 0xc2 is not defined, assign NOP function
	nop,
	// 0xc3 is not defined, assign NOP function
	nop,
	// 0xc4 CPY, zero page
	func(cpu *CPU6502) {
		cpu.cpy(cpu.zeroPage())
	},
	// 0xc5 CMP, zero page
	func(cpu *CPU6502) {
		cpu.cmp(cpu.zeroPage())
	},
	// 0xc6 DEC, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.dec)
	},
	// 0xc7 is not defined, assign NOP function
	nop,
	// 0xc8 INY
	func(cpu *CPU6502) {
		cpu.iny()
	},
	// 0xc9 CMP, immediate
	func(cpu *CPU6502) {
		cpu.cmp(cpu.immediate())
	},
	// 0xca DEX
	func(cpu *CPU6502) {
		cpu.dex()
	},
	// 0xcb is not defined, assign NOP function
	nop,
	// 0xcc CPY, absolute
	func(cpu *CPU6502) {
		cpu.cpy(cpu.absolute())
	},
	// 0xcd CMP, absolute
	func(cpu *CPU6502) {
		cpu.cmp(cpu.absolute())
	},
	// 0xce DEC, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.dec)
	},
	// 0xcf is not defined, assign NOP function
	nop,

	// 0xd0 BNE, relative
	func(cpu *CPU6502) {
		cpu.bne(cpu.relative())
	},
	// 0xd1 CMP, indirect indexed (address), Y
	func(cpu *CPU6502) {
		cpu.cmp(cpu.indirectIndexed())
	},
	// 0xd2 is not defined, assign NOP function
	nop,
	// 0xd3 is not defined, assign NOP function
	nop,
	// 0xd4 is not defined, assign NOP function
	nop,
	// 0xd5 CMP, zero page, X
	func(cpu *CPU6502) {
		cpu.cmp(cpu.zeroPageX())
	},
	// 0xd6 DEC, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.dec)
	},
	// 0xd7 is not defined, assign NOP function
	nop,
	// 0xd8 CLD
	func(cpu *CPU6502) {
		cpu.cld()
	},
	// 0xd9 CMP, absolute, Y
	func(cpu *CPU6502) {
		cpu.cmp(cpu.absoluteY())
	},
	// 0xda is not defined, assign NOP function
	nop,
	// 0xdb is not defined, assign NOP function
	nop,
	// 0xdc is not defined, assign NOP function
	nop,
	// 0xdd CMP, absolute, X
	func(cpu *CPU6502) {
		cpu.cmp(cpu.absoluteX())
	},
	// 0xde DEC, absolute, X
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteX(cpu.dec)
	},
	// 0xdf is not defined, assign NOP function
	nop,

	// 0xe0 CPX, immediate
	func(cpu *CPU6502) {
		cpu.cpx(cpu.immediate())
	},
	// 0xe1 SBC, indexed indirect, (address, X)
	func(cpu *CPU6502) {
		cpu.sbc(cpu.indexedIndirect())
	},
	// 0xe2 is not defined, assign NOP function
	nop,
	// 0xe3 is not defined, assign NOP function
	nop,
	// 0xe4 CPX, zero page
	func(cpu *CPU6502) {
		cpu.cpx(cpu.zeroPage())
	},
	// 0xe5 SBC, zero page
	func(cpu *CPU6502) {
		cpu.sbc(cpu.zeroPage())
	},
	// 0xe6 INC, zero page
	func(cpu *CPU6502) {
		cpu.inc(cpu.zeroPage())
	},
	// 0xe7 is not defined, assign NOP function
	nop,
	// 0xe8 INX
	func(cpu *CPU6502) {
		cpu.inx()
	},
	// 0xe9 SBC, immediate
	func(cpu *CPU6502) {
		cpu.sbc(cpu.immediate())
	},
	// 0xea NOP, this time it is real NOP function
	nop,
	// 0xeb is not defined, assign NOP function
	nop,
	// 0xec CPX, absolute
	func(cpu *CPU6502) {
		cpu.cpx(cpu.absolute())
	},
	// 0xed SBC, absolute
	func(cpu *CPU6502) {
		cpu.sbc(cpu.absolute())
	},
	// 0xee INC, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.inc)
	},
	// 0xef is not defined, assign NOP function
	nop,

	// 0xf0 BEQ
	func(cpu *CPU6502) {
		cpu.beq(cpu.relative())
	},
	// 0xf1 SBC, indirect indexed, (address), Y
	func(cpu *CPU6502) {
		cpu.sbc(cpu.indirectIndexed())
	},
	// 0xf2 is not defined, assign NOP function
	nop,
	// 0xf3 is not defined, assign NOP function
	nop,
	// 0xf4 is not defined, assign NOP function
	nop,
	// 0xf5 SBC, zero page, X
	func(cpu *CPU6502) {
		cpu.sbc(cpu.zeroPageX())
	},
	// 0xf6 INC, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.inc)
	},
	// 0xf7 is not defined, assign NOP function
	nop,
	// 0xf8 SED
	func(cpu *CPU6502) {
		cpu.sed()
	},
	// 0xf9 SBC, absolute, Y
	func(cpu *CPU6502) {
		cpu.sbc(cpu.absoluteY())
	},
	// 0xfa is not defined, assign NOP function
	nop,
	// 0xfb is not defined, assign NOP function
	nop,
	// 0xfc is not defined, assign NOP function
	nop,
	// 0xfd SBC, absolute, X
	func(cpu *CPU6502) {
		cpu.sbc(cpu.absoluteX())
	},
	// 0xfe INC, absolute,X
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteX(cpu.inc)
	},
	// 0xff is not defined, assign NOP function
	nop,
}
