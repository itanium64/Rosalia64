package extractor

import (
	"bytes"
	"strings"
)

type SIFSection uint32

const (
	NotImportant         SIFSection = 0
	SourceDisksNamesIA64 SIFSection = 1
	WinntDirectories     SIFSection = 2
	//Version              SIFSection = 0
	//Version              SIFSection = 0
	//Version              SIFSection = 0
	//Version              SIFSection = 0
	//Version              SIFSection = 0
	//Version              SIFSection = 0
)

type SIFFile struct {
	DirectoryTable map[string]string
	CdTable        map[string]string
	currentSection SIFSection
}

func ParseSIFFile(contents []byte) SIFFile {
	sifFile := SIFFile{
		currentSection: NotImportant,
		DirectoryTable: make(map[string]string),
		CdTable:        make(map[string]string),
	}

	lines := bytes.Split(contents, []byte{'\r', '\n'})

	for _, line := range lines {
		strLine := string(line)

		sifFile.NextLine(strLine)
	}

	return sifFile
}

func (sifFile *SIFFile) NextLine(line string) {
	if sifFile.currentSection == NotImportant {
		if strings.Contains(line, "[SourceDisksNames.ia64]") {
			sifFile.currentSection = SourceDisksNamesIA64
			return
		} else {
			return
		}
	}

	if strings.HasPrefix(line, "[") {
		switch line {
		case "[WinntDirectories]":
			sifFile.currentSection = WinntDirectories
			return
		}
	}

	switch sifFile.currentSection {
	case SourceDisksNamesIA64:
		sifFile.NextSourceDisksNamesIALine(line)
	case WinntDirectories:
		sifFile.NextWinntDirectoriesLine(line)
	}
}

func TrimAll(arr []string) {
	for index, _ := range arr {
		arr[index] = strings.TrimSpace(arr[index])
	}
}

func (sifFile *SIFFile) NextSourceDisksNamesIALine(line string) {
	if line == "" {
		return
	}

	equalsSplit := strings.Split(line, "=")
	TrimAll(equalsSplit)

	infoSplit := strings.Split(equalsSplit[1], ",")
	TrimAll(infoSplit)

	// equalsSplit[0]
	// |                      infoSplit[3]
	// |                           |
	// v                           v
	// 1  = %cdname%,%cdtagfilem%,,\ia64
	diskId := equalsSplit[0]
	path := infoSplit[3]

	sifFile.CdTable[diskId] = path
}

func (sifFile *SIFFile) NextWinntDirectoriesLine(line string) {
	if line == "" {
		return
	}

	equalsSplit := strings.Split(line, "=")
	TrimAll(equalsSplit)

	// equalsSplit[0]
	// |       equalsSplit[1]
	// |          |
	// v          v
	//39 = "Driver Cache\ia64"
	dirId := equalsSplit[0]
	name := strings.ReplaceAll(equalsSplit[1], "\"", "")

	sifFile.DirectoryTable[dirId] = name
}
