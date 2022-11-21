package petal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"rosalia64/core/decoding"
	"rosalia64/core/exe"
	"rosalia64/core/execution"
	"rosalia64/wiewiur"
	"rosalia64/wiewiur/win2003"
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

	var vmemSize int64
	var exeFilepath string
	var launchWiewiur bool
	var wiewiurExt bool
	var wiewiurIso string
	var wiewiurVersion wiewiur.SystemVersion

	for i := 1; i != len(os.Args); i++ {
		split := strings.Split(os.Args[i], "=")

		if len(split) == 0 {
			continue
		}

		var key, val string

		key = split[0]

		if len(split) == 2 {
			val = split[1]
		}

		switch key {
		case "-vmemsize":
			memSize, parseErr := strconv.ParseInt(val, 10, 64)

			if parseErr != nil {
				fmt.Printf("Failed to parse Argument 2. Not a valid integer.\n")
				return
			}

			vmemSize = memSize
		case "-exe":
			exeFilepath = val
		case "-wiewiur":
			launchWiewiur = true
		case "-wiewiur-extract":
			wiewiurExt = true
		case "-wiewiur-iso":
			wiewiurIso = val
		case "-wiewiur-sys":
			switch val {
			case "win2003":
				wiewiurVersion = wiewiur.WindowsServer2003
			}
		}
	}

	if launchWiewiur {
		if wiewiurExt {
			switch wiewiurVersion {
			case wiewiur.WindowsServer2003:
				extractor := win2003.CreateWindows2003Extractor()
				extractor.AssignDiskImage(wiewiurIso)
				extractor.ExtractFiles("iawin")

				return
			}
		}
	}

	fmt.Printf("Starting Execution of `%s` with %d Kilobytes of Memory.\n", exeFilepath, vmemSize)

	//Read exe file
	exeFile := exe.ReadExeFile(exeFilepath)

	var instructionData *bytes.Buffer
	var rdata *bytes.Buffer

	var textAddress int32

	//find .text for code and .rdata for the entrypoint
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
		}
	}

	var entryPoint uint64

	//read entrypoint
	binary.Read(rdata, binary.LittleEndian, &entryPoint)

	//initialize decoder
	execution.InitializeFunctionDeclarations()
	decoding.InitializeDecoderAndTables()

	//Store current address, this is so the decoder can store which instructions live where
	//this is used to know exactly to which instruction to jump to when branching
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

	//Initialize VM
	execution.InitializeMachine(uint64(vmemSize))
	execution.NewExecutionContext(decoding.DecodingContext.ExecutableInstructions, decoding.DecodingContext.InstructionStructs, decoding.DecodingContext.AddressToInstructionIndex[entryPoint])

	//Let it run free
	execution.CurrentExecutionContext.Run()

	fmt.Printf("\nIA64 Final Status Code: %d\n", execution.RetrieveGeneralRegister(8).Value)
}
