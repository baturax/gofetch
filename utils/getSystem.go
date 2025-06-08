package utils

import (
	"fmt"
)

func GetSystem() string {
	return fmt.Sprintf("  %v %v %v", getHost().Platform, getHost().OS, getHost().KernelArch)
}
