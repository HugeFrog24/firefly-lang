// magefile.go
//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// Variables
var (
	BuildDir     = "build"
	BinaryName   = "firefly"
	Source       = "main.go"
	VersionMajor = "1"
	VersionMinor = "0"
	VersionPatch = "0"
)

// Build compiles the binary.
func Build() error {
	// Create build directory if it doesn't exist
	if _, err := os.Stat(BuildDir); os.IsNotExist(err) {
		fmt.Println("Creating build directory...")
		if err := os.Mkdir(BuildDir, 0755); err != nil {
			return err
		}
	}

	binaryPath := fmt.Sprintf("%s%s%s", BuildDir, getSeparator(), BinaryName)
	if os.PathSeparator == '\\' {
		binaryPath += ".exe"
	}

	// Construct ldflags
	ldflags := fmt.Sprintf("-X main.versionMajor=%s -X main.versionMinor=%s -X main.versionPatch=%s",
		VersionMajor, VersionMinor, VersionPatch)

	// Build command
	cmd := exec.Command("go", "build", "-ldflags", ldflags, "-o", binaryPath, Source)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building the project...")
	return cmd.Run()
}

// Clean removes the build directory.
func Clean() error {
	if _, err := os.Stat(BuildDir); os.IsNotExist(err) {
		fmt.Println("Build directory does not exist. Nothing to clean.")
		return nil
	}

	fmt.Println("Removing build directory...")
	return os.RemoveAll(BuildDir)
}

// Rebuild cleans and then builds the project.
func Rebuild() {
	mg.SerialDeps(Clean, Build)
}

// Helper function to get OS-specific path separator
func getSeparator() string {
	if os.PathSeparator == '\\' {
		return "\\"
	}
	return "/"
}
