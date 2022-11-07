package exe

type Machine uint16
type Characteristics uint16
type ImageCharacteristics uint32
type DLLCharacteristics uint16
type Subsystem uint16
type Signature uint16

const (
	MachineIntel386         Machine = 0x14c
	Machinex64              Machine = 0x8664
	MachineR3000            Machine = 0x162
	MachineR10000           Machine = 0x168
	MachineMIPSLEv2         Machine = 0x169
	MachineAlphaAXP         Machine = 0x183
	MachineAlphaAXPOld      Machine = 0x184
	MachineSH3              Machine = 0x1a2
	MachineSH3DSP           Machine = 0x1a3
	MachineSH4              Machine = 0x1a6
	MachineSH5              Machine = 0x1a8
	MachineARMLE            Machine = 0x1c0
	MachineThumb            Machine = 0x1c2
	MachineARMv7            Machine = 0x1c4
	MachineAM33             Machine = 0x1d3
	MachinePPCLE            Machine = 0x1f0
	MachinePPCFloatingPoint Machine = 0x1f1
	MachinePPC64LE          Machine = 0x1f2
	MachineIA64             Machine = 0x200
	MachineMIPS16           Machine = 0x266
	MachineMotorolla68000   Machine = 0x268
	MachineAlphaAXP64       Machine = 0x284
	MachineMIPSFPU          Machine = 0x366
	MachineMIPS16FPU        Machine = 0x466
	MachineEFIByteCode      Machine = 0xebc
	MachineAMD64            Machine = 0x8664
	MachineM32RLE           Machine = 0x9041
	MachineARM64LE          Machine = 0xaa64
	MachineMSIL             Machine = 0xc0ee
)

func MachineToString(machine Machine) string {
	switch machine {
	case MachineIntel386:
		return "MachineIntel386"
	case MachineR3000:
		return "MachineR3000"
	case MachineR10000:
		return "MachineR10000"
	case MachineMIPSLEv2:
		return "MachineMIPSLEv2"
	case MachineAlphaAXP:
		return "MachineAlphaAXP"
	case MachineAlphaAXPOld:
		return "MachineAlphaAXPOld"
	case MachineSH3:
		return "MachineSH3"
	case MachineSH3DSP:
		return "MachineSH3DSP"
	case MachineSH4:
		return "MachineSH4"
	case MachineSH5:
		return "MachineSH5"
	case MachineARMLE:
		return "MachineARMLE"
	case MachineThumb:
		return "MachineThumb"
	case MachineARMv7:
		return "MachineARMv7"
	case MachineAM33:
		return "MachineAM33"
	case MachinePPCLE:
		return "MachinePPCLE"
	case MachinePPCFloatingPoint:
		return "MachinePPCFloatingPoint"
	case MachinePPC64LE:
		return "MachinePPC64LE"
	case MachineIA64:
		return "MachineIA64"
	case MachineMIPS16:
		return "MachineMIPS16"
	case MachineMotorolla68000:
		return "MachineMotorolla68000"
	case MachineAlphaAXP64:
		return "MachineAlphaAXP64"
	case MachineMIPSFPU:
		return "MachineMIPSFPU"
	case MachineMIPS16FPU:
		return "MachineMIPS16FPU"
	case MachineEFIByteCode:
		return "MachineEFIByteCode"
	case MachineAMD64:
		return "MachineAMD64"
	case MachineM32RLE:
		return "MachineM32RLE"
	case MachineARM64LE:
		return "MachineARM64LE"
	case MachineMSIL:
		return "MachineMSIL"
	}

	return ""
}

const (
	CharacteristicsRelocsStripped       Characteristics = 0x0001
	CharacteristicsExecutableImage      Characteristics = 0x0002
	CharacteristicsLineNumbersStripped  Characteristics = 0x0004
	CharacteristicsLocalSymbolsStripped Characteristics = 0x0008
	CharacteristicsAggressiveTrim       Characteristics = 0x0010
	CharacteristicsLargeAddressAware    Characteristics = 0x0020
	CharacteristicsBytesReservedLo      Characteristics = 0x0080
	CharacteristicsSupports32bitWord    Characteristics = 0x0100
	CharacteristicsDebugStripped        Characteristics = 0x0200
	CharacteristicsRemovableRunFromSwap Characteristics = 0x0400
	CharacteristicsNetRunFromSwap       Characteristics = 0x0800
	CharacteristicsSystemFile           Characteristics = 0x1000
	CharacteristicsDLLFile              Characteristics = 0x2000
	CharacteristicsSingleProcessorOnly  Characteristics = 0x4000
	CharacteristicsBytesReservedHi      Characteristics = 0x8000
)

const (
	SignatureExecutable32bit Signature = 0x10b
	SignatureExecutable64bit Signature = 0x20b
	SignatureROM             Signature = 0x107
)

const (
	SubsystemUnknown                Subsystem = 0
	SubsystemNative                 Subsystem = 1
	SubsystemWindowsGui             Subsystem = 2
	SubsystemWindowsCui             Subsystem = 3
	SubsystemOs2Cui                 Subsystem = 5
	SubsystemPosixCui               Subsystem = 7
	SubsystemWindowsCeGui           Subsystem = 9
	SubsystemEfiApplication         Subsystem = 10
	SubsystemEfiBootServiceDriver   Subsystem = 11
	SubsystemEfiRuntimeDriver       Subsystem = 12
	SubsystemEfiRom                 Subsystem = 13
	SubsystemXbox                   Subsystem = 14
	SubsystemWindowsBootApplication Subsystem = 16
)

const (
	CharacteristicsDynamicBase         DLLCharacteristics = 0x0040
	CharacteristicsForceIntegrity      DLLCharacteristics = 0x0080
	CharacteristicsNXCompat            DLLCharacteristics = 0x0100
	CharacteristicsNoIsolation         DLLCharacteristics = 0x0200
	CharacteristicsNoSeh               DLLCharacteristics = 0x0400
	CharacteristicsNoBind              DLLCharacteristics = 0x0800
	CharacteristicsAppContainer        DLLCharacteristics = 0x1000
	CharacteristicsWdmDriver           DLLCharacteristics = 0x2000
	CharacteristicsTerminalServerAware DLLCharacteristics = 0x8000
)

const (
	ImageCharacteristicsNoPad                     ImageCharacteristics = 0x00000008
	ImageCharacteristicsCode                      ImageCharacteristics = 0x00000020
	ImageCharacteristicsInitializedData           ImageCharacteristics = 0x00000040
	ImageCharacteristicsUninitializedData         ImageCharacteristics = 0x00000080
	ImageCharacteristicsOther                     ImageCharacteristics = 0x00000100
	ImageCharacteristicsInfo                      ImageCharacteristics = 0x00000200
	ImageCharacteristicsRemove                    ImageCharacteristics = 0x00000800
	ImageCharacteristicsComdat                    ImageCharacteristics = 0x00001000
	ImageCharacteristicsDeferSpeculativeExecution ImageCharacteristics = 0x00004000
	ImageCharacteristicsGlobalPointer             ImageCharacteristics = 0x00008000
	ImageCharacteristicsPurgeable                 ImageCharacteristics = 0x00020000
	ImageCharacteristicsLocked                    ImageCharacteristics = 0x00040000
	ImageCharacteristicsPreload                   ImageCharacteristics = 0x00080000
	ImageCharacteristicsAlign1Bytes               ImageCharacteristics = 0x00100000
	ImageCharacteristicsAlign2Bytes               ImageCharacteristics = 0x00200000
	ImageCharacteristicsAlign4Bytes               ImageCharacteristics = 0x00300000
	ImageCharacteristicsAlign8Bytes               ImageCharacteristics = 0x00400000
	ImageCharacteristicsAlign16Bytes              ImageCharacteristics = 0x00500000
	ImageCharacteristicsAlign32Bytes              ImageCharacteristics = 0x00600000
	ImageCharacteristicsAlign64Bytes              ImageCharacteristics = 0x00700000
	ImageCharacteristicsAlign128Bytes             ImageCharacteristics = 0x00800000
	ImageCharacteristicsAlign256Bytes             ImageCharacteristics = 0x00900000
	ImageCharacteristicsAlign512Bytes             ImageCharacteristics = 0x00A00000
	ImageCharacteristicsAlign1024Bytes            ImageCharacteristics = 0x00B00000
	ImageCharacteristicsAlign2048Bytes            ImageCharacteristics = 0x00C00000
	ImageCharacteristicsAlign4096Bytes            ImageCharacteristics = 0x00D00000
	ImageCharacteristicsAlign8192Bytes            ImageCharacteristics = 0x00E00000
	ImageCharacteristicsExtraRelocations          ImageCharacteristics = 0x01000000
	ImageCharacteristicsDiscardable               ImageCharacteristics = 0x02000000
	ImageCharacteristicsNotCached                 ImageCharacteristics = 0x04000000
	ImageCharacteristicsNotPaged                  ImageCharacteristics = 0x08000000
	ImageCharacteristicsShared                    ImageCharacteristics = 0x10000000
	ImageCharacteristicsExecute                   ImageCharacteristics = 0x20000000
	ImageCharacteristicsRead                      ImageCharacteristics = 0x40000000
	ImageCharacteristicsWrite                     ImageCharacteristics = 0x80000000
)
