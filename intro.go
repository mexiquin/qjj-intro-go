package qjj_intro

import (
	"fmt"
	"time"
)

// function to clear console
func clear() {
	fmt.Print("\033[H\033[2J")
}

// function to print string in typewriter style
func typewriter(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Print("\n")
}

// sho loading bar for 3 seconds on same line
func loading(sec int) {
	fmt.Print("Loading")
	for i := 0; i < sec; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
}

func Show_intro() {
	qjj_string := `
⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⠛⠛⠛⠛⠛⢻⡇
⠀⣿⣿⣿⡿⠟⠋⠉⠁⠀⠈⠉⠙⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⢸⡇
⠀⣿⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿⣿⣿⣿⣿⣿⣇⠀⠀⠀⢸⡇
⠀⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⢸⡇
⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀⢸⡇
⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣦⣼⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⢸⡇
⠀⣿⣄⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠀⠀⠀⢸⡇
⠀⣿⣿⣦⣀⠀⠀⠀⠀⠀⠀⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⢸⡇
⠀⣿⣿⣿⣿⣷⣦⣤⣤⣤⣤⣤⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣀⠀⢸⡇
⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⢸⡇
⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠛⢿⣿⣿⣿⠏⠀⠀⢸⡇
⠀⡟⠻⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠟⠋⠀⠀⠀⠀⠈⠛⠁⠀⠀⠀⢸⡇
⠀⡇⠀⠀⠀⠈⠉⠉⠉⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇
⠀⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠃
⠀⡇⢹⢸⠈⡆⡏⢹⠐⡇⡇⢸⢰⡇⡎⢹⠀⢽⠀⡇⢸⢩⡆⢻⠈⡇⣾⢹⡂⠀
⠀⡇⢸⢸⠀⡇⡇⢸⠀⡇⡇⢸⢸⡇⡇⠸⠀⢸⠀⡇⢸⢸⡇⢸⠀⡇⣿⢸⠇⠀
⠀⡇⢸⢸⠀⡇⡇⢸⠀⡇⡇⢸⢸⡇⡇⠀⠀⢸⠀⡇⢸⢸⡇⣾⡷⠇⠹⡀⠀⠀
⠀⡏⠁⢸⠙⡆⡇⢸⢰⡇⡇⢸⢸⡇⡇⠀⠀⢸⠀⡇⢸⢸⡇⡇⢸⠀⣀⢹⡀⠀
⠀⡇⠀⢸⠀⡅⡇⢸⢘⡇⡇⢸⢸⡇⡇⢸⠀⢸⠀⡇⢸⢸⡇⣯⢸⡇⣿⢸⠇⠀
⠀⠇⠀⠸⠀⠇⠣⠜⠐⠦⠇⠘⠴⠃⠳⠜⠀⠚⠀⠇⠘⠴⠃⠛⠘⠇⠻⠼⠁
`
	clear()
	typewriter(qjj_string)
	loading(4)
	clear()
}
