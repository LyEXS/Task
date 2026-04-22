package utils

import (
	"os"
	"runtime"
)

func GetOS() string {
	return runtime.GOOS
}

func GetFilePath() string {
	return os.Args[1]
}
