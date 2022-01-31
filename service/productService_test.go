package service

import (
	"os"
	"os/exec"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	products := GetProducts("products.json")

	if len(products) == 0 {
		t.Errorf("Expected length higher than zero, but not %v", len(products))
	}

	for _, product := range products {
		if product.Price <= 0 {
			t.Errorf("Expected price greater than 0, but price is %v", product.Price)
		}
	}
}

func TestReadFromFileNotExistFile(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		GetProducts("no_file.json")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
