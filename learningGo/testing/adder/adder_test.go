package adder

import (
	"os"
	"testing"
)

func Test_add(t *testing.T) {
    result := add(1, 2)
    if result != 3 {
        t.Errorf("add(1, 2) failed. Got %d, expected 3.", result)
    }
}

func createTempFile(t *testing.T) (string, error) {
    f, err := os.Create("tempfile")
    if err != nil {
        return "", err
    }
    
    t.Cleanup(func() {
        os.Remove(f.Name())
        t.Log("Removed a temporary file.")
    })
    return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
    filename, err := createTempFile(t)
    if err != nil {
        t.Fatalf("Error: %s", err)
    }
    t.Logf("Name of temporary file: %s", filename)
    // Don't worry about cleanup
}
    
