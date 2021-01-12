package exe

import (
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestJavaIsInstalled(t *testing.T) {
	isInstalled := JavaIsInstalled()
	if !isInstalled {
		t.Fatalf("Java not installed.")
	}
}

func TestFileExists(t *testing.T) {
	exists := fileexists("go.mod")
	if !exists {
		t.Fatalf("go.mod not exists.")
	}
}

func TestCmd(t *testing.T) {
	output, err := cmd("meu-software", "args")
	fmt.Println("output: " + output)

	if err == nil {
		t.Fatalf("meu-software existe")
	}
}
