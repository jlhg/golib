package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Tee logs message on both file and standard output.
func Tee(logger *logrus.Logger, filePath string) (err error) {
	var f *os.File
	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	mw := io.MultiWriter(os.Stdout, f)
	logger.SetOutput(mw)

	return
}

// SetJSONFormat sets JSON formatter to logger.
func SetJSONFormat(logger *logrus.Logger) {
	logger.Formatter = new(logrus.JSONFormatter)
}
