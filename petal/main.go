package petal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"rosalia64/core/decoding"
	"rosalia64/core/exe"
	"rosalia64/core/execution"
	"strconv"
	"strings"
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
		fmt.Printf("Failed to parse Argument 2. Not a valid integer.\n")
		return
	}

	fmt.Printf("Starting Execution of `%s` with %d Kilobytes of Memory.\n", exeFilepath, vmemSize)

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

	execution.InitializeFunctionDeclarations()
	decoding.InitializeDecoderAndTables()

	for {
		var bundle [16]byte

		err := binary.Read(instructionData, binary.LittleEndian, &bundle)

		if err != nil {
			break
		}

		decoding.DecodingContext.NextBundle(bundle)
	}

	//for execution.ContinueRunning {

	//}

	execution.InitializeMachine(uint64(vmemSize))
	execution.NewExecutionContext(decoding.DecodingContext.ExecutableInstructions, decoding.DecodingContext.InstructionStructs)

	execution.CurrentExecutionContext.Run()

	fmt.Printf("\nIA64 Final Status Code: %d\n", execution.RetrieveGeneralRegister(8).Value)
}
