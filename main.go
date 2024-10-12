package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func copyToClipboard(content string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		// macOS
		cmd = exec.Command("pbcopy")
	case "linux":
		// Linux
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		// Windows
		cmd = exec.Command("clip")
	default:
		return fmt.Errorf("unsupported platform")
	}

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	_, err = in.Write([]byte(content))
	if err != nil {
		return err
	}

	if err = in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}

func readFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: clip example.txt")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := readFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	err = copyToClipboard(content)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
		os.Exit(1)
	}

	fmt.Println("File content copied to clipboard successfully!")
}
