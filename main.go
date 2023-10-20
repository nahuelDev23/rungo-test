package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userProgram := r.FormValue("code")
		fmt.Printf("%s", userProgram)
		response, err := runGoProgram([]byte(userProgram))

		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(response))
	})

	http.ListenAndServe(":8080", nil)

}

func runGoProgram(program []byte) ([]byte, error) {

	programFileUrl, err := filepath.Abs("./user-program/main.go")
	if err != nil {
		return nil, err
	}

	programFile, err := os.Create(programFileUrl)
	if err != nil {
		return nil, err
	}

	defer programFile.Close()
	if _, err := io.Copy(programFile, bytes.NewReader(program)); err != nil {
		return nil, err
	}

	// go.mod file creation
	goModUrl, err := filepath.Abs("./user-program/go.mod")
	if err != nil {
		return nil, err
	}

	goModFile, err := os.Create(goModUrl)
	if err != nil {
		return nil, err
	}
	defer goModFile.Close()

	goModContent := `module user-program

go 1.21.2`

	_, err = io.Copy(goModFile, strings.NewReader(goModContent))
	if err != nil {
		return nil, err
	}
	return exec.Command("go", "run", "./user-program/main.go").Output()
}
