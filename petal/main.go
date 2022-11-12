package petal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"rosalia64/core/decoding"
	"rosalia64/core/exe"
	"rosalia64/core/execution"
	"rosalia64/core/structures"
	"strconv"
	"strings"
)

func DebugUI() {

}

func PetalMain() {
	//Stack test
	stack := structures.Stack[int]{}

	stack.Push(123)
	stack.Push(89)
	stack.Push(1111)

	popped := stack.Pop()
	popped = popped

	top := stack.Top()
	top = top

	stack.Push(9231)

	top2 := stack.Top()
	top2 = top2

	popped2 := stack.Pop()
	popped2 = popped2

	popped3 := stack.Pop()
	popped3 = popped3

	stack.Push(69420)

	popped4 := stack.Pop()
	popped4 = popped4

	popped5 := stack.Pop()
	popped5 = popped5

	//End of stack test

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
		fmt.Printf("Failed to parse Argument 2. Not a valid integer.\n")
		return
	}

	fmt.Printf("Starting Execution of `%s` with %d Kilobytes of Memory.\n", exeFilepath, vmemSize)

	exeFile := exe.ReadExeFile(exeFilepath)

	var instructionData *bytes.Buffer
	var rdata *bytes.Buffer

	var textAddress int32
	//var rdataAddress int32

	for _, image := range exeFile.ImageSections {
		isText := strings.HasPrefix(string(image.Name[:]), ".text")
		isRdata := strings.HasPrefix(string(image.Name[:]), ".rdata")

		if isText {
			rawData := exeFile.RawFileData[image.PointerToRawData : image.PointerToRawData+image.SizeOfRawData]
			textAddress = image.VirtualAddress
			instructionData = bytes.NewBuffer(rawData)
		}

		if isRdata {
			rawData := exeFile.RawFileData[image.PointerToRawData : image.PointerToRawData+image.SizeOfRawData]
			rdata = bytes.NewBuffer(rawData)
			//rdataAddress = image.VirtualAddress
		}
	}

	var entryPoint uint64

	binary.Read(rdata, binary.LittleEndian, &entryPoint)

	execution.InitializeFunctionDeclarations()
	decoding.InitializeDecoderAndTables()

	currentAddress := exeFile.COFFOptionalHeader.OptionalHeader.ImageBase + int64(textAddress)

	for {
		var bundle [16]byte

		err := binary.Read(instructionData, binary.LittleEndian, &bundle)

		if err != nil {
			break
		}

		decoding.DecodingContext.NextBundle(bundle, uint64(currentAddress))

		currentAddress += 16
	}

	execution.InitializeMachine(uint64(vmemSize))
	execution.NewExecutionContext(decoding.DecodingContext.ExecutableInstructions, decoding.DecodingContext.InstructionStructs, decoding.DecodingContext.AddressToInstructionIndex[entryPoint])

	execution.CurrentExecutionContext.Run()

	fmt.Printf("\nIA64 Final Status Code: %d\n", execution.RetrieveGeneralRegister(8).Value)
}
