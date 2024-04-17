package main

import "github.com/fatih/color"

func PrintLogo() string {
	return color.WhiteString("      ____        ") + color.YellowString("  ______  ______   _        _______ ") + color.WhiteString("        ____      ") + "\n" +
		color.WhiteString(" ____|    \\       ") + color.YellowString(" / _____)| ___  \\ | |      |  _____)") + color.WhiteString("       /    |____ ") + "\n" +
		color.WhiteString("(____|     '._____") + color.YellowString("/ /      | |__)  )| |      | |______") + color.WhiteString("_____.'     |____)") + "\n" +
		color.WhiteString(" ____|       _|___") + color.YellowString("  |      |  __  ( | |      |  ______") + color.WhiteString("___|_       |____ ") + "\n" +
		color.WhiteString("(____|     .'     ") + color.YellowString("\\ \\_____ | |__)  )| |_____ | |_____ ") + color.WhiteString("     '.     |____)") + "\n" +
		color.WhiteString("     |____/       ") + color.YellowString(" \\______)|______/ |_______)|_______)") + color.WhiteString("       \\____|     ")
}
