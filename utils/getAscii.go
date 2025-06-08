package utils

import "strings"

func GetAscii() []string {
	distro := GetDistro()
	switch {
	case strings.Contains(distro, "arch"):
		return []string{
			"      /\\          ",
			"     /  \\         ",
			"    /    \\        ",
			"   /      \\       ",
			"  /   ,,   \\      ",
			" /   |  |   \\     ",
			"/_-''    ''-_\\    ",
		}
	case strings.Contains(distro, "void"):
		return []string{
			"    _______        ",
			" _ \\______ -      ",
			"| \\  ___  \\ |     ",
			"| | /   \\ | |     ",
			"| | \\___/ | |     ",
			"| \\______ \\_|     ",
			" -_______\\        ",
		}
	case strings.Contains(distro, "alpine"):
		return []string{
			"      /\\          ",
			"     /  \\         ",
			"    / /\\ \\  /\\    ",
			"   / /  \\ \\/  \\   ",
			"  / /    \\ \\ /\\   ",
			" / / /|   \\ \\ \\   ",
			"/_/ /_|    \\_\\_\\  ",
		}
	case strings.Contains(distro, "ubuntu"):
		return []string{
			"         _        ",
			"     ---(_)       ",
			" _/  ---  \\       ",
			"(_) |   |         ",
			"  \\  --- _/       ",
			"     ---(_)       ",
		}

	case strings.Contains(distro, "cachy"):
		return []string{
			"   /''''''''''''/      ",
			"  /''''''''''''/       ",
			" /''''''/              ",
			"/''''''/               ",
			"\\......\\               ",
			" \\......\\              ",
			"  \\.............../    ",
			"   \\............./     ",
		}

	default:
		return []string{
			"    ___           ",
			"   (.. \\          ",
			"   (<> |          ",
			"  //  \\ \\         ",
			" ( |  | /|        ",
			"_/\\ __)/_)        ",
			"\\/-____\\/         ",
		}
	}
}
