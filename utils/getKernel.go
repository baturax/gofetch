package utils

import (
	"fmt"
)

func GetKernel() string {
	return fmt.Sprintf("  %v %v", getHost().Platform, getHost().KernelVersion)
}
