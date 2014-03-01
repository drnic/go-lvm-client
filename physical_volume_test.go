package lvm_client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/starkandwayne/go-lvm-client"
)

var _ = Describe("PhysicalVolume", func() {
	It("loads PVs from pvdisplay", func() {
		pvs, _ := PhysicalVolumes()
		Expect(len(pvs)).To(Equal(1))
	})
})
