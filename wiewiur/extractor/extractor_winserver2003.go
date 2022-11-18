package extractor

import (
	"os"
)

type WindowsServer2003Extractor struct {
	isoFile string
}

func (extractor *WindowsServer2003Extractor) AssignDiskImage(location string) {
	extractor.isoFile = location

	//isoFile, err := os.Open(location)

	//if err != nil {
	//	panic("Failed to open .iso file!")
	//}

	//if err = isoUtil.ExtractImageToDirectory(isoFile, "_ext_temp"); err != nil {
	//panic("failed to extract to a temporary directory")
	//}
}

func (extractor *WindowsServer2003Extractor) ExtractFiles(location string) bool {
	txtSetup, err1 := os.ReadFile("_ext_temp/IA64/TXTSETUP.SIF")

	if err1 != nil {
		return false
	}

	txtSetupSif := ParseSIFFile(txtSetup)
	txtSetupSif = txtSetupSif

	//os.RemoveAll("_ext_temp")

	return true
}

func (extractor *WindowsServer2003Extractor) ExtractRegistryData() string {
	return ""
}

func CreateWindows2003Extractor() WindowsInstallExtractor {
	return &WindowsServer2003Extractor{}
}
