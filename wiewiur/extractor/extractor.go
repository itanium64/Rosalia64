package extractor

type WindowsInstallExtractor interface {
	AssignDiskImage(location string)
	ExtractFiles(location string) bool
	ExtractRegistryData() string
}
