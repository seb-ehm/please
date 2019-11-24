package util

import "os"

//IsNoPipe returns whether a file is no pipe
func IsNoPipe(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	mode := info.Mode()
	isPipe := mode&os.ModeNamedPipe != 0
	isCharDevice := mode&os.ModeDevice != 0 && mode&os.ModeCharDevice != 0

	return !isPipe || isCharDevice
}
