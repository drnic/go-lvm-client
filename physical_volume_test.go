package lvm_client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/starkandwayne/go-lvm-client"
)

var _ = Describe("PhysicalVolume", func() {
	Describe("parse colon output", func() {
		It("new from colon output", func() {
			pv := NewPhysicalVolume()
			err := pv.ParseDisplayWithColons("/dev/sda5:vg0:84254720:-1:8:8:-1:4096:10284:0:10284:IKGNO5-Dx7w-2UBv-rUzw-ekJg-e496-9RQ5cP")
			Expect(err).To(BeNil())
			Expect(pv.PVName).To(Equal("/dev/sda5"))
		})
		It("invalid number of tokens", func() {
			pv := NewPhysicalVolume()
			err := pv.ParseDisplayWithColons("x:y:z")
			Expect(err).ToNot(BeNil())
		})
	})
	It("loads PVs from pvdisplay", func() {
		pvs, _ := PhysicalVolumes()
		Expect(len(pvs)).To(Equal(1))
	})
})
