// log
package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type MyFormatter struct {
	TimestampFormat string
	LogFormat       string
}

// Format building log message.
func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	output = strings.Replace(output, "%time%", entry.Time.Format(f.TimestampFormat), 1)
	output = strings.Replace(output, "%msg%", entry.Message, 1)
	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)
	output = strings.Replace(output, "%file%", entry.Caller.File, 1)
	output = strings.Replace(output, "%file%", filepath.Base(entry.Caller.File), 1)
	output = strings.Replace(output, "%line%", strconv.Itoa(entry.Caller.Line), 1)

	for k, v := range entry.Data {
		output = fmt.Sprintf("%v %v:%v", output, k, v)
	}
	output = output + "\n"
	return []byte(output), nil
}
