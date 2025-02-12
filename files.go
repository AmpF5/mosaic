package main

import (
	"os/exec"
	"runtime"
)

func OpenFileExplorer() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer")
	case "darwin":
		cmd = exec.Command("open", ".")
	case "linux":
		cmd = exec.Command("xdg-open", ".")
	default:
		return
	}

	// Run the command in the background
	err := cmd.Start()
	if err != nil {
		println("Error:", err.Error())
	}
}
