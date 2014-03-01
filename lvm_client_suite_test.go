package lvm_client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLvmClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LVM Client Suite")
}
