package utils

import (
	"fmt"
)

func GetUptime() string {
	uptime := getHost().Uptime

	day := uptime / 86400
	hour := (uptime % 86400) / 3600
	minute := (uptime % 3600) / 60

	return fmt.Sprintf(" %v days, %v hours, %v minutes", day, hour, minute)
}
