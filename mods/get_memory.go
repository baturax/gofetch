package mods

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func GetMem() string {
	m, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println(err.Error())
	}

	usedMem := m.Used / 1048576
	totalMem := m.Total / 1048576

	return fmt.Sprintf("  %v / %v", usedMem, totalMem)
}
