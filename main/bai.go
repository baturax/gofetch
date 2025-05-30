package main

import (
	"baturax/mods"
	"fmt"
)

func main() {
	mL := len(mods.GetIcon())

	if l := len(icons()); l > mL {
		mL = l
	}

	for i := range mL {
		icon := ""
		info := ""

		if i < len(mods.GetIcon()) {
			icon = mods.GetIcon()[i]
		} else {
			icon = "                  "
		}

		if i < len(icons()) {
			info = icons()[i]
		}

		fmt.Println(icon, info)

	}
}

func icons() []string {
	return []string{
		mods.GetUser(),
		mods.GetSystem(),
		mods.GetKernel(),
		mods.GetMem(),
		mods.GetUptime(),
		mods.GetShell(),
		mods.GetDesktop(),
		mods.GetPackages(),
	}
}
