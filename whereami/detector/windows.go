package detector

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// WindowsDetect determines if the system is running Windows within a VM and checks if WSL is available.
func WindowsDetect() {
	if runtime.GOOS != "windows" {
		fmt.Println("Not running on Windows.")
		return
	}

	// Print the hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
	}

	// Check if the system is running in a VM
	isVM := isRunningInVM()
	if isVM {
		fmt.Println("The system is running on a virtual machine.")
	} else {
		fmt.Println("The system is not running on a virtual machine.")
	}

	checkWSL()
}

// isRunningInVM performs a low-level check to see if the system is running in a VM
func isRunningInVM() bool {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	getSystemFirmwareTable := kernel32.NewProc("GetSystemFirmwareTable")

	// Querying SMBIOS for information
	const smbios = "RSMB"     // SMBIOS firmware table identifier
	buf := make([]byte, 4096) // Buffer to store the result
	smbiosPtr := syscall.UTF16FromString(smbios)

	ret, _, err := getSystemFirmwareTable.Call(
		uintptr(smbiosPtr),
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if ret == 0 {
		fmt.Printf("Error querying SMBIOS: %v\n", err)
		return false
	}

	// Check the buffer content for common VM indicators
	smbiosString := string(buf)
	if containsVMIdentifier(smbiosString) {
		return true
	}

	return false
}

// containsVMIdentifier checks the firmware table string for VM-specific identifiers
func containsVMIdentifier(s string) bool {
	vmIdentifiers := []string{
		"VMware", "VirtualBox", "Hyper-V", "QEMU", "KVM", "Parallels",
	}

	for _, identifier := range vmIdentifiers {
		if contains(s, identifier) {
			return true
		}
	}
	return false
}

// contains checks if a substring exists in a string
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s[:len(substr)] == substr || contains(s[1:], substr))
}

// checkWSL checks if WSL is installed and validates WSL2 functionality
func checkWSL() {
	cmd := exec.Command("wsl", "--list", "--verbose")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the WSL command to list distributions
	err := cmd.Run()
	if err != nil {
		fmt.Printf("WSL is not installed or accessible: %v\n", stderr.String())
		return
	}

	output := out.String()
	fmt.Println("WSL Distributions:")
	fmt.Println(output)

	// Parse the output to determine WSL version
	if strings.Contains(output, "WSL 2") {
		fmt.Println("WSL2 is enabled and accessible.")
		validateWSL2()
	} else if strings.Contains(output, "WSL 1") {
		fmt.Println("WSL1 is enabled and accessible.")
	} else {
		fmt.Println("No active WSL distributions found.")
	}
}

// validateWSL2 runs a simple test within WSL2 to ensure it's working
func validateWSL2() {
	cmd := exec.Command("wsl", "-e", "uname", "-a")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error validating WSL2: %v\n", stderr.String())
		return
	}

	output := out.String()
	fmt.Println("WSL2 Test Output:")
	fmt.Println(output)
}
