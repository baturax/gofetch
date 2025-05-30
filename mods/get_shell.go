package mods

import (
	"fmt"
	"os"
)

func GetShell() string {
	s := os.Getenv("SHELL")
	if s == "" {
		return fmt.Sprintf("󰲒 sorry %v", s)
	}

	return fmt.Sprintf("󰲒 %v", s)

}
