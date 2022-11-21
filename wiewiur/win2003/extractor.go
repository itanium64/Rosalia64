package win2003

import (
	"os"
	//isoUtil "github.com/kdomanski/iso9660/util"
)

type WindowsServer2003Extractor struct {
	isoFile string
}

func (extractor *WindowsServer2003Extractor) AssignDiskImage(location string) {
	extractor.isoFile = location

	//isoFile, err := os.Open(location)
	//
	//if err != nil {
	//	panic("Failed to open .iso file!")
	//}
	//
	//if err = isoUtil.ExtractImageToDirectory(isoFile, "_ext_temp"); err != nil {
	//	panic("failed to extract to a temporary directory")
	//}
}

func (extractor *WindowsServer2003Extractor) ExtractFiles(location string) bool {
	txtSetup, err1 := os.ReadFile("_ext_temp/IA64/TXTSETUP.SIF")

	if err1 != nil {
		return false
	}

	os.Mkdir(location, os.ModePerm)
	os.Mkdir(location+"/drive_c", os.ModePerm)
	os.Mkdir(location+"/drive_c/Windows", os.ModePerm)
	os.Mkdir(location+"/drive_c/Windows/Driver Cache", os.ModePerm)
	os.Mkdir(location+"/drive_c/Windows/Driver Cache/ia64", os.ModePerm)
	os.Mkdir(location+"/drive_c/Windows/system32", os.ModePerm)
	os.Mkdir(location+"/drive_c/Windows/system32/drivers", os.ModePerm)
	os.Mkdir(location+"/drive_c/Windows/system32/drivers/etc", os.ModePerm)

	txtSetupSif := ParseSIFFile(txtSetup)
	txtSetupSif.Extract(location)

	//os.RemoveAll("_ext_temp")

	return true
}

func (extractor *WindowsServer2003Extractor) ExtractRegistryData() string {
	return ""
}

func CreateWindows2003Extractor() *WindowsServer2003Extractor {
	return &WindowsServer2003Extractor{}
}
