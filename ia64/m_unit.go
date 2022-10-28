package ia64

import "fmt"

var M_UnitInstructionTable map[uint64]func(instructionBits uint64) = map[uint64]func(instructionBits uint64){
	9: AddImmediate22,
}

func SignExt(i uint64, n uint32) int64 {
	return (((int64)(i) << (64 - (n))) >> (64 - (n)))
}

func AddImmediate22(instructionBits uint64) {
	s := instructionBits & (0b00001 << 41) >> 42

	imm9d := (instructionBits & (0b00000111111111 << 32)) >> 32
	imm5c := (instructionBits & (0b0000000000000011111 << 27)) >> 27
	imm7b := (instructionBits & (0b00000000000000000000011111111 << 18)) >> 18

	fmt.Printf("sign: %d\n", s)
	fmt.Printf("imm9d: %064b\n", imm9d)
	fmt.Printf("imm5c: %064b\n", imm5c)
	fmt.Printf("imm7b: %064b\n", imm7b)

	immediate := SignExt(s<<21|imm5c<<16|imm9d<<7|imm7b, 22)

	fmt.Printf("imm22: %d\n", immediate)
	fmt.Printf("imm22: %064b\n", immediate)
}
