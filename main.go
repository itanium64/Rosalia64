package main

import (
	"Rosalia64/exe"
	"Rosalia64/ia64"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"lukechampine.com/uint128"
)

func main() {
	data, err := os.ReadFile("Rimukoro.exe")

	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(data)

	//Actual reading part
	var dosHeader exe.DOSHeader

	binary.Read(buffer, binary.LittleEndian, &dosHeader)

	peData := data[dosHeader.PEPointer:]
	peBuffer := bytes.NewBuffer(peData)

	var coffHeader exe.COFFHeader
	var peHeaderBytes [4]byte

	binary.Read(peBuffer, binary.LittleEndian, &peHeaderBytes)
	binary.Read(peBuffer, binary.LittleEndian, &coffHeader)

	var signature exe.Signature

	binary.Read(peBuffer, binary.LittleEndian, &signature)

	if signature == exe.SignatureExecutable32bit {
		panic("IA64 isnt 32bit")
	}

	var peOptHeader exe.COFFOptionalHeader64

	binary.Read(peBuffer, binary.LittleEndian, &peOptHeader)

	var dataDirectories []exe.DataDirectory

	for i := int32(0); i != peOptHeader.OptionalHeader.NumberOfRvaAndSizes; i++ {
		var dataDirectory exe.DataDirectory

		binary.Read(peBuffer, binary.LittleEndian, &dataDirectory)

		dataDirectories = append(dataDirectories, dataDirectory)
	}

	var imageSections []exe.ImageSectionHeader

	for i := int32(0); i != int32(coffHeader.NumberOfSections); i++ {
		var imageSectionHeader exe.ImageSectionHeader

		binary.Read(peBuffer, binary.LittleEndian, &imageSectionHeader)

		imageSections = append(imageSections, imageSectionHeader)
	}

	//TODO: don't do this wrong! you arent supposed to start from .text
	var instructionData *bytes.Buffer

	for _, image := range imageSections {
		isText := strings.HasPrefix(string(image.Name[:]), ".text")

		if isText {
			rawData := data[image.PointerToRawData : image.PointerToRawData+image.SizeOfRawData]
			instructionData = bytes.NewBuffer(rawData)
			break
		}
	}

	for {
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

		fmt.Printf("\n\n\nNEW BUNDLE: template (decimal): %d\n\n\n", template)

		fmt.Printf("high : %064b\n", asUint128.Hi)
		fmt.Printf("low  :                                                                 %064b\n     :\n", asUint128.Lo)
		fmt.Printf("whole: %064b%064b\n", asUint128.Hi, asUint128.Lo)
		fmt.Printf("slot0:                                                                 %064b\n", slot0)
		fmt.Printf("slot1:                        %064b\n", slot1)
		fmt.Printf("slot2: %064b\n", slot2<<18)

		DecodeInstructionSlot(slot0, slot1, unitOrder.Slot0)
		DecodeInstructionSlot(slot1, slot2, unitOrder.Slot1)
		DecodeInstructionSlot(slot2, 0b000, unitOrder.Slot2)

	}
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
