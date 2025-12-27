package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func identifyFdType(path string) string {
	if strings.HasPrefix(path, "socket:") {
		return "SOCKET"
	} else if strings.HasPrefix(path, "pipe:") {
		return "PIPE"
	} else if strings.HasPrefix(path, "/dev") {
		return "DEVICE"
	} else {
		return "FILE"
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fd-spy <PID>")
		os.Exit(1)
	}
	userProcess := os.Args[1]
	procSelfFdPath := fmt.Sprintf("/proc/%s/fd", userProcess)
	files, err := os.ReadDir(procSelfFdPath)
	if err != nil {
		fmt.Printf("Failed to open %s: %s", procSelfFdPath, err)
		return
	}

	for _, f := range files {
		fullPath := filepath.Join(procSelfFdPath, f.Name())
		targetPath, err := os.Readlink(fullPath)
		if err != nil {
			fmt.Printf("%s: %s\n", fullPath, err)
			continue
		}
		if f.Name() == "0" {
			fmt.Printf("%-8s %s (STDIN)\n", identifyFdType(targetPath), targetPath)
		} else if f.Name() == "1" {
			fmt.Printf("%-8s %s (STDOUT)\n", identifyFdType(targetPath), targetPath)
		} else if f.Name() == "2" {
			fmt.Printf("%-8s %s (STDERR)\n", identifyFdType(targetPath), targetPath)
		} else {
			fmt.Printf("%-8s %s\n", identifyFdType(targetPath), targetPath)
		}
	}
}
