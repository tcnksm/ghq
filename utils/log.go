package utils

import (
	"os"

	"github.com/motemen/go-colorine"
)

var logger = &colorine.Logger{
	colorine.Prefixes{
		"git":  colorine.Verbose,
		"hg":   colorine.Verbose,
		"skip": colorine.Verbose,
		"cd":   colorine.Verbose,

		"open":   colorine.Warn,
		"exists": colorine.Warn,

		"authorized": colorine.Notice,

		"error": colorine.Error,

		"": colorine.Info,
	},
}

func Log(prefix, message string) {
	logger.Log(prefix, message)
}

func ErrorIf(err error) bool {
	if err != nil {
		Log("error", err.Error())
		return true
	}

	return false
}

func DieIf(err error) {
	if err != nil {
		Log("error", err.Error())
		os.Exit(1)
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
