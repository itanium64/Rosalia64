package win2003

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CopyFile(src string, dest string) error {
	fin, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fin.Close()

	fout, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		return err
	}

	return nil
}

func FileExists(location string) bool {
	if _, err := os.Stat(location); err == nil {
		return true
	}

	return false
}

func (sifFile *SIFFile) Extract(location string) {
	for _, file := range sifFile.SourceDiskFiles {
		if file.NewInstallCode == 3 {
			continue
		}

		newLocation := location + "/drive_c/Windows"
		newLocation += "/" + sifFile.DirectoryTable[file.TargetSubdirectory]

		newLocation = strings.ReplaceAll(newLocation, "\\", "/")

		var newFilename string

		if file.TargetFilename == "" {
			newFilename += "/" + file.IsoFilename
		} else {
			newFilename += "/" + file.TargetFilename
		}

		sourceLocation := "_ext_temp"

		sourceLocationExtension := strings.ReplaceAll(sifFile.CdTable[file.Source], "\\", "/") + "/"

		cabbedLocation := sourceLocationExtension + file.IsoFilename
		cabbedLocation = cabbedLocation[:len(cabbedLocation)-1] + "_"

		normalLocation := sourceLocationExtension + file.IsoFilename

		//Check for lowercase
		cabbedExists := FileExists(sourceLocation + cabbedLocation)
		normalExists := FileExists(sourceLocation + normalLocation)

		cabbedFilename := sourceLocation + cabbedLocation
		normalFilename := sourceLocation + normalLocation

		if !cabbedExists && !normalExists {
			cabbedFilename = sourceLocation + strings.ToUpper(cabbedLocation)
			normalFilename = sourceLocation + strings.ToUpper(normalLocation)

			//check for uppercase
			cabbedExists = FileExists(cabbedFilename)
			normalExists = FileExists(normalFilename)

			if !cabbedExists && !normalExists {
				fmt.Printf("Neither %s nor %s exist!\n", normalFilename, cabbedFilename)
			}
		}

		if cabbedExists {
			switch runtime.GOOS {
			case "windows":
				expand := exec.Command("C:\\Windows\\System32\\expand.exe", cabbedFilename, newLocation+newFilename)
				expand.Output()
			case "linux":
				cabextract := exec.Command("cabextract", "-d", newLocation, "./"+cabbedFilename)
				cabextract.Output()
			}
		}

		if normalExists {
			CopyFile(normalFilename, newLocation+newFilename)
		}

		//fmt.Printf("%d : Copied file %s...\n", index, newLocation+newFilename)
	}
}
