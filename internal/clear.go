package internal

import (
	"os"
	"os/exec"
	"runtime"
)

// clear terminal screen
func Screen() {
	var cmd *exec.Cmd
	// defer fmt.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
