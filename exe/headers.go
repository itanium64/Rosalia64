package exe

type DOSHeader struct {
	Signature          [2]byte
	LastSize           uint16
	NumBlocks          uint16
	NumRelocations     uint16
	HdrSize            uint16
	MinAlloc           uint16
	MaxAlloc           uint16
	SS                 uint16
	SP                 uint16
	Checksum           uint16
	IP                 uint16
	CS                 uint16
	RelocationPosition uint16
	NumOverlay         uint16

	Reserved1 [4]uint16

	OemId   uint16
	OemInfo uint16

	Reserved2 [10]uint16

	PEPointer uint32
}

type DataDirectory struct {
	VirtualAddress int32
	Size           int32
}

type COFFHeader struct {
	Machine              Machine
	NumberOfSections     uint16
	Timestamp            int32
	PointerToSymbolTable int32
	NumberOfSymbols      int32
	SizeOfOptionalHeader uint16
	Characteristics      Characteristics
}

type COFFOptionalHeader struct {
	ImageBase             int64
	SectionAlignment      int32
	FileAlignment         int32
	MajorOSVersion        uint16
	MinorOSVersion        uint16
	MajorImageVersion     uint16
	MinorImageVersion     uint16
	MajorSubsystemVersion uint16
	MinorSubsystemVersion uint16
	Win32VersionValue     int32
	SizeOfImage           int32
	SizeOfHeaders         int32
	Checksum              int32
	Subsystem             Subsystem
	DLLCharacteristics    DLLCharacteristics
	SizeOfStackReserve    int64
	SizeOfStackCommit     int64
	SizeOfHeapReserve     int64
	SizeOfHeapCommit      int64
	LoaderFlags           int32
	NumberOfRvaAndSizes   int32
	//DataDirectory * NumberOfRvaAndSizes
}

type COFFOptionalHeader32 struct {
	//No signature
	MajorLinkerVersion      byte
	MinorLinkerVersion      byte
	SizeOfCode              int32
	SizeOfInitializedData   int32
	SizeOfUninitializedData int32
	AddressOfEntryPoint     int32
	BaseOfCode              int32
	BaseOfData              int32

	OptionalHeader COFFOptionalHeader
}

type COFFOptionalHeader64 struct {
	//No signature
	MajorLinkerVersion      byte
	MinorLinkerVersion      byte
	SizeOfCode              int32
	SizeOfInitializedData   int32
	SizeOfUninitializedData int32
	AddressOfEntryPoint     int32
	BaseOfCode              int32

	OptionalHeader COFFOptionalHeader
}

type ImageSectionHeader struct {
	Name                         [8]byte
	PhysicalAddressOrVirtualSize int32

	VirtualAddress       int32
	SizeOfRawData        int32
	PointerToRawData     int32
	PointerToRelocations int32
	PointerToLinenumbers int32
	NumberOfRelocations  int16
	NumberOfLineNumbers  int16
	ImageCharacteristics ImageCharacteristics
}
