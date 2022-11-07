package petal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"rosalia64/core/exe"
	"rosalia64/core/ia64"
	"strconv"
	"strings"

	"lukechampine.com/uint128"
)

func DebugUI() {

}

func PetalMain() {
	if len(os.Args) < 3 {
		fmt.Printf("Less command-line arguments than expected!\n")
		fmt.Printf("Arguments:\n\n")
		fmt.Printf("Rosalia64 <IA64 exe location> <RAM in Kilobytes>\n")
		return
	}

	exeFilepath := os.Args[1]
	vmemSizeArg := os.Args[2]

	vmemSize, parseErr := strconv.ParseInt(vmemSizeArg, 10, 64)

	if parseErr != nil {
		fmt.Printf("Failed to parse Argument 2. Not a valid integer.")
		return
	}

	fmt.Printf("Starting Execution of `%s` with %d Kilobytes of Memory.", exeFilepath, vmemSize)

	ia64.InitializeMachine(uint64(vmemSize))

	exeFile := exe.ReadExeFile("Rimukoro.exe")

	//TODO: don't do this wrong! you arent supposed to start from .text
	var instructionData *bytes.Buffer

	for _, image := range exeFile.ImageSections {
		isText := strings.HasPrefix(string(image.Name[:]), ".text")

		if isText {
			rawData := exeFile.RawFileData[image.PointerToRawData : image.PointerToRawData+image.SizeOfRawData]
			instructionData = bytes.NewBuffer(rawData)
			break
		}
	}

	for ia64.ContinueRunning {
		var bundle [16]byte

		err := binary.Read(instructionData, binary.LittleEndian, &bundle)

		if err != nil {
			break
		}

		asUint128 := uint128.FromBytes(bundle[:])

		template := asUint128.Lo & 0b11111

		unitOrder := ia64.UnitTable[template]

		slot0 := (asUint128.Lo & 0b000000000001111111111111111111111111111111111111111100000)
		slot1 := (asUint128.Lo&0b111111111110000000000000000000000000000000000000000000000)>>41 |
			(asUint128.Hi&0b000000000000000000000000000111111111111111111111111111111)<<23

		slot2 := (asUint128.Hi & 0b1111111111111111111111111111111111111111100000000000000000000000) >> 18

		//fmt.Printf("\n\n\nNEW BUNDLE: template (decimal): %d\n\n\n", template)
		//
		//fmt.Printf("high : %064b\n", asUint128.Hi)
		//fmt.Printf("low  :                                                                 %064b\n     :\n", asUint128.Lo)
		//fmt.Printf("whole: %064b%064b\n", asUint128.Hi, asUint128.Lo)
		//fmt.Printf("slot0:                                                                 %064b\n", slot0)
		//fmt.Printf("slot1:                        %064b\n", slot1)
		//fmt.Printf("slot2: %064b\n", slot2<<18)

		DecodeInstructionSlot(slot0, slot1, unitOrder.Slot0)
		DecodeInstructionSlot(slot1, slot2, unitOrder.Slot1)
		DecodeInstructionSlot(slot2, 0b000, unitOrder.Slot2)
	}

	fmt.Printf("\nIA64 Final Status Code: %d\n", ia64.RetrieveGeneralRegister(8).Value)
}

func DecodeInstructionSlot(slot uint64, nextSlot uint64, unit ia64.Unit) {
	majorOpcode := slot & (0b1111 << 42) >> 42

	table := ia64.GetInstructionTable(unit)

	instruction, exists := table[majorOpcode]

	if !exists {
		fmt.Printf("\nUNIMPLEMENTED!!!: \n")
		fmt.Printf("unit : %s\n", ia64.UnitToString(unit))
		fmt.Printf("major: %d\n", majorOpcode)

		return
	}

	instruction(slot, nextSlot)
}
