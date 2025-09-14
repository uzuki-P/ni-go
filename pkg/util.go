package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type PackageManager string

const (
	Npm     PackageManager = "npm"
	Yarn    PackageManager = "yarn"
	Pnpm    PackageManager = "pnpm"
	Unknown PackageManager = ""
)

func CallCliCommand(packageManager PackageManager, command string) error {
	switch packageManager {
	case Npm:
		return exec.Command("npm", command).Run()
	case Yarn:
		return exec.Command("yarn", command).Run()
	case Pnpm:
		return exec.Command("pnpm", command).Run()
	default:
		return fmt.Errorf("unsupported package manager")
	}
}

func CheckLockFileExists() (PackageManager, error) {
	if _, err := os.Stat("package-lock.json"); err == nil {
		return Npm, nil
	}
	if _, err := os.Stat("yarn.lock"); err == nil {
		return Yarn, nil
	}
	if _, err := os.Stat("pnpm-lock.yaml"); err == nil {
		return Pnpm, nil
	}
	return "", fmt.Errorf("no lock file found")
}

func CheckPackageManager() (PackageManager, error) {
	// Read package.json file
	data, err := os.ReadFile("package.json")
	if err != nil {
		return "", fmt.Errorf("failed to read package.json: %w", err)
	}

	// Parse package.json file
	var pkg struct {
		PackageManager string `json:"packageManager"`
	}
	err = json.Unmarshal(data, &pkg)
	if err != nil {
		return "", fmt.Errorf("failed to parse package.json: %w", err)
	}

	// Return the first word that split with "@"
	return PackageManager(strings.Split(pkg.PackageManager, "@")[0]), nil
}

func GetPackageManager() (PackageManager, error) {
	packageManager, err := CheckLockFileExists()
	if err != nil {
		fmt.Println(err)
		return CheckPackageManager()
	}

	return packageManager, err
}
