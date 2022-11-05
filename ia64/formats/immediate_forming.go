package formats

func SignExt(i uint64, n uint32) int64 {
	return (((int64)(i) << (64 - (n))) >> (64 - (n)))
}

func Imm22(sign uint64, imm5c uint64, imm9d uint64, imm7b uint64) int64 {
	return SignExt(sign<<21|imm5c<<16|imm9d<<7|imm7b, 22)
}

func Imm14(sign uint64, imm6d uint64, imm7b uint64) int64 {
	return SignExt(sign<<13|imm6d<<7|imm7b, 14)
}
