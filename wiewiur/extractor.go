package wiewiur

type WindowsInstallExtractor interface {
	AssignDiskImage(location string)
	ExtractFiles(location string) bool
}
