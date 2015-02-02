package main_test

import (
	"github.com/cloudfoundry/cli/testhelpers/plugin_builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)

	plugin_builder.BuildTestBinary(".", "test_user_cmd")

	RunSpecs(t, "Main Suite")
}
