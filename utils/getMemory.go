package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func GetMemory() string {
	memory, err := mem.VirtualMemory()
	showError(err)

	usedMem := memory.Used / 1048576
	totalMem := memory.Total / 1048576

	return fmt.Sprintf("  %v / %v", usedMem, totalMem)
}
