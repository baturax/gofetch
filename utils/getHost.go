package utils

import (
	"github.com/shirou/gopsutil/v4/host"
)

func getHost() *host.InfoStat {
	h, err := host.Info()
	showError(err)
	return h
}
