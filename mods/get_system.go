package mods

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func GetSystem() string {
	h, err := host.Info()

	if err != nil {
		fmt.Println(err.Error())
	}

	platform := h.Platform
	os := h.OS
	arch := h.KernelArch

	return fmt.Sprintf("  %v %v %v", platform, os, arch)
}
