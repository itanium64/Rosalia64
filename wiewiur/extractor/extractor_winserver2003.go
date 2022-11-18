package extractor

type WindowsServer2003Extractor struct {
	isoFile string
}

func (extractor *WindowsServer2003Extractor) AssignDiskImage(location string) {
	extractor.isoFile = location
}

func (extractor *WindowsServer2003Extractor) ExtractFiles(location string) bool {

}
