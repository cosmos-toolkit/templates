package commands

import (
	"bytes"
	"os"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = old }()

	VersionCmd.SetArgs([]string{})
	err := VersionCmd.Execute()
	if err != nil {
		t.Fatalf("version: %v", err)
	}
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	out := buf.String()
	if out != "0.1.0\n" {
		t.Errorf("version output: got %q", out)
	}
}

func TestRunCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = old }()

	RunCmd.SetArgs([]string{"--name", "test"})
	err := RunCmd.Execute()
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	out := buf.String()
	if out != "Hello, test!\n" {
		t.Errorf("run output: got %q", out)
	}
}
