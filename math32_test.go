package math32

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestVetMath32(t *testing.T) {
	linuxArches := []string{"amd64", "arm", "arm64", "s390x", "ppc64le", "riscv64", "386", "mips", "mips64", "mipsle", "mips64le"}

	// Linux architectures
	for _, GOARCH := range linuxArches {
		GOOS := "linux"
		t.Run(fmt.Sprintf("GOOS=%s GOARCH=%s", GOOS, GOARCH), func(t *testing.T) {
			goVet(t, GOOS, GOARCH)
		})
	}
	// WASM
	t.Run("GOOS=js GOARCH=wasm", func(t *testing.T) {
		goVet(t, "js", "wasm")
	})

	// AIX
	t.Run("GOOS=aix GOARCH=ppc64", func(t *testing.T) {
		goVet(t, "aix", "ppc64")
	})
}

func goVet(t *testing.T, GOOS, GOARCH string) {
	cmd := exec.Command("go", "vet", ".")
	cmd.Env = append(os.Environ(), "GOOS="+GOOS, "GOARCH="+GOARCH)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(string(output), err)
	}
}
