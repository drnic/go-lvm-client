package lvm_client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/starkandwayne/go-lvm-client"
	"github.com/starkandwayne/go-lvm-client/system"
)

var _ = Describe("PhysicalVolume", func() {
	Describe("parse colon output", func() {
		It("new from colon output", func() {
			pv := NewPhysicalVolume()
			err := pv.ParseLine("  /dev/sda5:precise64:lvm2:a-:81672.00:0", ":")
			Expect(err).To(BeNil())
			Expect(pv.PVName).To(Equal("/dev/sda5"))
			Expect(pv.VGName).To(Equal("precise64"))
			Expect(pv.Format).To(Equal("lvm2"))
			Expect(pv.Attr).To(Equal("a-"))
			Expect(pv.PVSize).To(Equal(81672.0)) // Mb
			Expect(pv.FreePE).To(Equal(0.0))
		})

		It("invalid number of tokens", func() {
			pv := NewPhysicalVolume()
			err := pv.ParseLine("x:y:z", ":")
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("parse pvs output", func() {
		It("initial-vagrant", func() {
			systemRepo := &system.FakeSystemRepository{
				PvsOutput: "  /dev/sda5:precise64:lvm2:a-:81672.00:0\n  /dev/sda6:precise64:lvm2:a-:81672.00:0",
			}
			pvs, err := PhysicalVolumes(systemRepo)
			Expect(err).To(BeNil())
			Expect(len(pvs)).To(Equal(2))
			Expect(pvs[0].PVName).To(Equal("/dev/sda5"))
			Expect(pvs[1].PVName).To(Equal("/dev/sda6"))
		})
	})
})
