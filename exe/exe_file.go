package exe

import (
	"bytes"
	"encoding/binary"
	"os"
)

type EXEFile struct {
	RawFileData        []byte
	DOSHeader          DOSHeader
	COFFHeader         COFFHeader
	PEBuffer           *bytes.Buffer
	COFFOptionalHeader COFFOptionalHeader64
	DataDirectories    []DataDirectory
	ImageSections      []ImageSectionHeader
}

func ReadExeFile(path string) EXEFile {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(data)

	//Actual reading part
	var dosHeader DOSHeader

	binary.Read(buffer, binary.LittleEndian, &dosHeader)

	peData := data[dosHeader.PEPointer:]
	peBuffer := bytes.NewBuffer(peData)

	var coffHeader COFFHeader
	var peHeaderBytes [4]byte

	binary.Read(peBuffer, binary.LittleEndian, &peHeaderBytes)
	binary.Read(peBuffer, binary.LittleEndian, &coffHeader)

	var signature Signature

	binary.Read(peBuffer, binary.LittleEndian, &signature)

	if signature == SignatureExecutable32bit {
		panic("IA64 isnt 32bit")
	}

	var peOptHeader COFFOptionalHeader64

	binary.Read(peBuffer, binary.LittleEndian, &peOptHeader)

	var dataDirectories []DataDirectory

	for i := int32(0); i != peOptHeader.OptionalHeader.NumberOfRvaAndSizes; i++ {
		var dataDirectory DataDirectory

		binary.Read(peBuffer, binary.LittleEndian, &dataDirectory)

		dataDirectories = append(dataDirectories, dataDirectory)
	}

	var imageSections []ImageSectionHeader

	for i := int32(0); i != int32(coffHeader.NumberOfSections); i++ {
		var imageSectionHeader ImageSectionHeader

		binary.Read(peBuffer, binary.LittleEndian, &imageSectionHeader)

		imageSections = append(imageSections, imageSectionHeader)
	}

	return EXEFile{
		RawFileData:        data,
		DOSHeader:          dosHeader,
		COFFHeader:         coffHeader,
		COFFOptionalHeader: peOptHeader,
		DataDirectories:    dataDirectories,
		ImageSections:      imageSections,
	}
}
