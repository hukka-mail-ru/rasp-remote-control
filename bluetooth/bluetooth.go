package bluetooth

import (
	"strconv"
	"strings"
)

func ParseMsg(msg string) (int, error) {
	//	log.Info("parsing: ", string(msg))

	msg = strings.TrimFunc(msg, func(r rune) bool {
		return r == '\n' || r == '\x00'
	})

	//	log.Info("trimmed string: ", string(msg))

	var i int
	var err error
	i, err = strconv.Atoi(string(msg))

	return i, err
}
