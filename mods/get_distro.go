package mods

import (
	"github.com/shirou/gopsutil/v4/host"
)

func getDistro() string {
	h, _ := host.Info()
	d := h.Platform
	return d
}
