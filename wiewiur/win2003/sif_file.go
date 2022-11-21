package win2003

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type SIFSection uint32

const (
	NotImportant         SIFSection = 0
	SourceDisksNames     SIFSection = 1
	SourceDisksNamesIA64 SIFSection = 2
	WinntDirectories     SIFSection = 3
	SourceDisksFiles     SIFSection = 4
	SourceDisksFilesIA64 SIFSection = 5
	Strings              SIFSection = 6
)

type SIFFile struct {
	DirectoryTable  map[string]string
	CdTable         map[string]string
	Strings         map[string]string
	SourceDiskFiles []SourceDiskFile
	currentSection  SIFSection
}

type SourceDiskFile struct {
	IsoFilename        string
	Source             string
	SourceSubdirectory string
	TargetSubdirectory string
	TargetFilename     string
	NewInstallCode     int64
}

func (sourceDiskFile *SourceDiskFile) ExpandString(key string, value string) {
	sourceDiskFile.IsoFilename = strings.ReplaceAll(sourceDiskFile.IsoFilename, key, value)
	sourceDiskFile.Source = strings.ReplaceAll(sourceDiskFile.Source, key, value)
	sourceDiskFile.SourceSubdirectory = strings.ReplaceAll(sourceDiskFile.SourceSubdirectory, key, value)
	sourceDiskFile.TargetSubdirectory = strings.ReplaceAll(sourceDiskFile.TargetSubdirectory, key, value)
	sourceDiskFile.TargetFilename = strings.ReplaceAll(sourceDiskFile.TargetFilename, key, value)
}

func ParseSIFFile(contents []byte) SIFFile {
	sifFile := SIFFile{
		currentSection: NotImportant,
		DirectoryTable: make(map[string]string),
		CdTable:        make(map[string]string),
		Strings:        make(map[string]string),
	}

	lines := bytes.Split(contents, []byte{'\r', '\n'})

	for _, line := range lines {
		strLine := string(line)

		sifFile.NextLine(strLine)
	}

	for key, value := range sifFile.CdTable {
		for str, replaceWith := range sifFile.Strings {
			if !strings.Contains(value, "%") {
				break
			}

			value = strings.ReplaceAll(value, "%"+str+"%", replaceWith)
		}

		sifFile.CdTable[key] = value
	}

	for key, value := range sifFile.DirectoryTable {
		for str, replaceWith := range sifFile.Strings {
			if !strings.Contains(value, "%") {
				break
			}

			value = strings.ReplaceAll(value, "%"+str+"%", replaceWith)
		}

		sifFile.DirectoryTable[key] = value
	}

	for index := range sifFile.SourceDiskFiles {
		for str, replaceWith := range sifFile.Strings {
			sifFile.SourceDiskFiles[index].ExpandString("%"+str+"%", replaceWith)
		}
	}

	return sifFile
}

func (sifFile *SIFFile) NextLine(line string) {
	if line == "" {
		return
	}

	if strings.HasPrefix(line, "[") {
		switch line {
		case "[WinntDirectories]":
			sifFile.currentSection = WinntDirectories
			return
		case "[SourceDisksNames]":
			sifFile.currentSection = SourceDisksNames
			return
		case "[SourceDisksNames.x86]":
			sifFile.currentSection = NotImportant
			return
		case "[SourceDisksNames.amd64]":
			sifFile.currentSection = NotImportant
			return
		case "[sourcedisksfiles.ia64]":
			sifFile.currentSection = SourceDisksFiles
			return
		case "[SourceDisksNames.ia64]":
			sifFile.currentSection = SourceDisksNamesIA64
			return
		case "[SourceDisksFiles]":
			sifFile.currentSection = SourceDisksFiles
			return
		case "[SourceDisksFiles.ia64]":
			sifFile.currentSection = SourceDisksFiles
			return
		case "[Strings]":
			sifFile.currentSection = Strings
			return
		default:
			sifFile.currentSection = NotImportant
			fmt.Printf("Unimplemented Section: %s\n", line)
			return
		}
	}

	switch sifFile.currentSection {
	case SourceDisksNamesIA64:
		sifFile.NextSourceDisksNamesLine(line)
	case SourceDisksNames:
		sifFile.NextSourceDisksNamesLine(line)
	case WinntDirectories:
		sifFile.NextWinntDirectoriesLine(line)
	case Strings:
		sifFile.NextStringsLine(line)
	case SourceDisksFiles:
		sifFile.NextSourceDisksFilesLine(line)
	}
}

func TrimAll(arr []string) {
	for index := range arr {
		arr[index] = strings.TrimSpace(arr[index])
	}
}

func (sifFile *SIFFile) NextSourceDisksNamesLine(line string) {
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

func (sifFile *SIFFile) NextSourceDisksFilesLine(line string) {
	equalsSplit := strings.Split(line, "=")
	TrimAll(equalsSplit)

	infoSplit := strings.Split(equalsSplit[1], ",")
	TrimAll(infoSplit)

	//filename to copy, is still .dll even if compressed and saved as .dl_ on ISO
	//equalsSplit[0]
	//    |     status, this is apperantly the source of the file, should match one entry in [SourceDisksNames.*]
	//    |     infoSplit[0]
	//    |          |
	//    |          | specifies the source subdirectory to copy the file from
	//    |          | infoSplit[1]
	//    |          | |
	//    |          | | specifies the uncompressed size of the file
	//    |          | | infoSplit[2]
	//    |          | | |
	//    |          | | | unknown
	//    |          | | | infoSplit[3]
	//    |          | | | |
	//    |          | | | | unknown
	//    |          | | | | infoSplit[4]
	//    |          | | | | |
	//    |          | | | | | unknown
	//    |          | | | | | infoSplit[5]
	//    |          | | | | | |
	//    |          | | | | | |  diskid is used when installing from floppies, should match one entry in [SourceDisksNames.*]
	//    |          | | | | | |  infoSplit[6]
	//    |          | | | | | |  |
	//    |          | | | | | |  |  specifies the target subdirectory to copy to
	//    |          | | | | | |  |  infoSplit[7]
	//    |          | | | | | |  |   |
	//    |          | | | | | |  |   |  upgradecode, used when upgrading windows (0 always copy, 1 copy if exists in install dir, 2 dont copy if exists in install dir, 3 dont copy)
	//    |          | | | | | |  |   |  infoSplit[8]
	//    |          | | | | | |  |   |  |
	//    |          | | | | | |  |   |  |  newinstallcode1, used when installing fresh (0 always copy, 1 copy if exists in install dir, 2 dont copy if exists in install dir, 3 dont copy)
	//    |          | | | | | |  |   |  |  infoSplit[9]
	//    |          | | | | | |  |   |  |  |
	//    |          | | | | | |  |   |  |  | newfilename specifies the target filename
	//    |          | | | | | |  |   |  |  | infoSplit[10]
	//    |          | | | | | |  |   |  |  | |
	//    |          | | | | | |  |   |  |  | |  newinstallcode2, unknown
	//    |          | | | | | |  |   |  |  | |  infoSplit[11]
	//    |          | | | | | |  |   |  |  | |  |
	//    |          | | | | | |  |   |  |  | |  |  unknown
	//    |          | | | | | |  |   |  |  | |  |  infoSplit[12]
	//    |          | | | | | |  |   |  |  | |  |  |
	//    |          | | | | | |  |   |  |  | |  |  |
	//    |          | | | | | |  |   |  |  | |  |  |
	//    v          v v v v v v  v   v  v  v v  v  v
	//bootvid.dll  = 1, , , , , , 3_, 2, 0, 0, , 1, 2

	var source, sourceSubdir, targetSubdir, newInstallCode, newFilename string

	isoFilename := equalsSplit[0]
	source = infoSplit[0]
	sourceSubdir = infoSplit[1]

	if len(infoSplit) > 7 {
		targetSubdir = infoSplit[7]
	}

	if len(infoSplit) > 9 {
		newInstallCode = infoSplit[9]
	}

	if len(infoSplit) > 10 {
		newFilename = infoSplit[10]
	}

	var newInstallCodeInt int64 = -1

	if newInstallCode != "" {
		newInstallCodeInt, _ = strconv.ParseInt(newInstallCode, 10, 64)
	}

	sifFile.SourceDiskFiles = append(sifFile.SourceDiskFiles, SourceDiskFile{
		IsoFilename:        isoFilename,
		Source:             source,
		SourceSubdirectory: sourceSubdir,
		TargetSubdirectory: targetSubdir,
		TargetFilename:     newFilename,
		NewInstallCode:     newInstallCodeInt,
	})
}

func (sifFile *SIFFile) NextStringsLine(line string) {
	equalsSplit := strings.Split(line, "=")
	TrimAll(equalsSplit)

	sifFile.Strings[equalsSplit[0]] = strings.ReplaceAll(equalsSplit[1], "\"", "")
}
