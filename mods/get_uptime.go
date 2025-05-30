package mods

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func GetUptime() string {
	h, err := host.Info()
	if err != nil {
		fmt.Println(err.Error())
	}

	uptime := h.Uptime

	day := uptime / 86400
	hour := (uptime % 86400) / 3600
	minute := (uptime % 3600) / 60

	return fmt.Sprintf("  %v days, %v hours, %v minutes", day, hour, minute)
}
