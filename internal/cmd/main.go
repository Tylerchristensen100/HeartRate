package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"os/exec"
)

const (
	pythonScriptPath        = "scripts/graph.py"
	requirementsPath        = "scripts/requirements.txt"
	writtenScriptPath       = "graph.py"
	writtenRequirementsPath = "requirements.txt"
)

//go:embed scripts/graph.py
var pythonScript embed.FS

//go:embed scripts/requirements.txt
var requirements embed.FS

func RunPythonScript(args ...string) error {
	err := writePythonScript()
	if err != nil {
		return fmt.Errorf("failed to write python script: %v", err)
	}

	cmd := exec.Command("python3", writtenScriptPath)
	cmd.Args = append(cmd.Args, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running python script: %v\nOutput: %s", err, out.String())
	}

	deletePythonScript()
	return nil
}

func writePythonScript() error {

	_, err := exec.LookPath("python3")
	if err != nil {
		return fmt.Errorf("python 3 is not installed or not in PATH")
	}

	scriptData, err := pythonScript.ReadFile(pythonScriptPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded python script: %v", err)
	}

	requirementsData, err := requirements.ReadFile(requirementsPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded requirements file: %v", err)
	}

	err = os.WriteFile(writtenScriptPath, scriptData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write python script: %v", err)
	}
	err = os.WriteFile(writtenRequirementsPath, requirementsData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write requirements file: %v", err)
	}

	cmd := exec.Command("pip3", "install", "-r", writtenRequirementsPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install python packages: %v\nOutput: %s", err, out.String())
	}

	return nil
}

func deletePythonScript() error {
	err := os.Remove(writtenScriptPath)
	if err != nil {
		return fmt.Errorf("failed to delete python script: %v", err)
	}
	err = os.Remove(writtenRequirementsPath)
	if err != nil {
		return fmt.Errorf("failed to delete requirements file: %v", err)
	}
	return nil
}
