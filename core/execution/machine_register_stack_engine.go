package execution

import "rosalia64/core/structures"

type RegisterStackWindow struct {
}

type RegisterStackEngine struct {
	StackWindows structures.Stack[RegisterStackWindow]
}
