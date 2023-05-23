package utils_log

import (
	"io"
	"log"
	"strings"
)

func NewShortDebug(msgOut io.WriteCloser, prefixes ...string) *log.Logger {
	return log.New(msgOut, makePrefixes("DEBUG", prefixes...), log.Ltime|log.Lshortfile)
}

func makePrefixes(s string, prefixes ...string) string {
	return "[" + "[" + s + "]" + "[" + strings.Join(prefixes, "]-[") + "]]"
}
