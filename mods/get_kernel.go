package mods

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func GetKernel() string {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	kernel := h.KernelVersion

	return fmt.Sprintf("  %v %v", platform, kernel)
}
