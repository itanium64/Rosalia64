package win2003

import (
	"bytes"
	"strings"
)

type RegistryEntry struct {
	RegRoot                  string
	Subkey                   string
	ValueEntryName           string
	Flags                    uint32
	Value                    string
	SecurityDescriptorString string
}

type HiveInfFile struct {
	Strings                map[string]string
	Sections               map[string][]RegistryEntry
	currentSectionName     string
	currentSectionElements []RegistryEntry
}

func ExtractRegistryInfFile(location string, contents []byte) {
	hiveInfFile := HiveInfFile{
		Strings:                make(map[string]string),
		Sections:               make(map[string][]RegistryEntry),
		currentSectionName:     "",
		currentSectionElements: []RegistryEntry{},
	}

	lines := bytes.Split(contents, []byte{'\r', '\n'})

	for i := 0; i != len(lines); i++ {
		strLine := string(lines[i])

		for strings.HasSuffix(strLine, "\\") {
			strLine += string(lines[i+1])
			i += 1
		}

		if strings.HasPrefix(strLine, ";") || hiveInfFile.currentSectionName == "Version" {
			continue
		}

		if strings.HasPrefix(strLine, "[") {
			if hiveInfFile.currentSectionName != "" {
				hiveInfFile.Sections[hiveInfFile.currentSectionName] = hiveInfFile.currentSectionElements
			}

			hiveInfFile.currentSectionName = strings.ReplaceAll(strings.ReplaceAll(strLine, "[", ""), "]", "")
			hiveInfFile.currentSectionElements = []RegistryEntry{}

			continue
		}

	}

}
