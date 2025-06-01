package main

import (
	"baturax/mods"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	mL := len(mods.GetIcon())
	ri := rand.Intn(len(colors()))
	re := colors()[ri]

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

		fmt.Println(re+icon, info)
	}

}

func colors() []string {
	return []string{
		"\033[0m",
		"\033[31m",
		"\033[32m",
		"\033[33m",
		"\033[34m",
		"\033[35m",
		"\033[36m",
		"\033[97m",
	}
}

func icons() []string {

	lol := []string{
		mods.GetUser(),
		mods.GetSystem(),
		mods.GetKernel(),
		mods.GetMem(),
		mods.GetUptime(),
		mods.GetShell(),
		mods.GetDesktop(),
	}

	if os.Getenv("SHOW_PACKAGES") == "yes" {
		lol = append(lol, mods.GetPackages())
	}

	return lol

}
