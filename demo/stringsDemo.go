package demo

import (
	"strings"
)

func stringDemo() interface{} {
	command := "read"
	return strings.Contains(command, "read")
}
