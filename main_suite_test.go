package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)

	cmd := exec.Command("go", "build", "-o", "test_user_cmd.exe")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	RunSpecs(t, "Main Suite")
}
