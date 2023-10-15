package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("node", "script.js")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Resultado:")
	fmt.Println(string(output))
}
