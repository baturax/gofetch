package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetBattery() string {
	batteryFiles, err := filepath.Glob("/sys/class/power_supply/BAT*/capacity")
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}

	if len(batteryFiles) == 0 {
		return "No Battery Found"
	}

	var builder strings.Builder

	for _, capacityFile := range batteryFiles {
		dir := filepath.Dir(capacityFile)
		statusFile := filepath.Join(dir, "status")

		capacityData, err := os.ReadFile(capacityFile)
		if err != nil {
			builder.WriteString(fmt.Sprintf("Unable to read capacity: %s\n", capacityFile))
			continue
		}
		capacity := strings.TrimSpace(string(capacityData))

		status := "Unknown"
		if data, err := os.ReadFile(statusFile); err == nil {
			status = strings.TrimSpace(string(data))
		}

		builder.WriteString(fmt.Sprintf("  %s: %s%% (%s)\n", filepath.Base(dir), capacity, status))
	}

	return builder.String()
}
