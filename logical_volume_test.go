package lvm_client_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/starkandwayne/go-lvm-client"
  "github.com/starkandwayne/go-lvm-client/system"
)

var _ = Describe("LogicalVolume", func() {
  Describe("parse colon output", func() {
    It("new from colon output", func() {
      lv := NewLogicalVolume()
      err := lv.ParseLine("  root:precise64:owi-ao:80904.00::::::", ":")
      Expect(err).To(BeNil())
      Expect(lv.LVName).To(Equal("root"))
      Expect(lv.VGName).To(Equal("precise64"))
      Expect(lv.Attrs).To(Equal("owi-ao"))
      Expect(lv.VolumeType).To(Equal(LVTOrigin))
      Expect(lv.Writable).To(BeTrue())
      Expect(lv.AllocationPolicy).To(Equal(LVATInherited))
      Expect(lv.Locked).ToNot(BeTrue())
      Expect(lv.FixedMinor).ToNot(BeTrue())
      Expect(lv.State).To(Equal(LVStateActive))
      Expect(lv.DeviceOpen).To(BeTrue())

      Expect(lv.LVSize).To(Equal(80904.0)) // Mb
    })

    It("invalid number of tokens", func() {
      lv := NewLogicalVolume()
      err := lv.ParseLine("x:y:z", ":")
      Expect(err).ToNot(BeNil())
    })
  })

  Describe("parse lvs output", func() {
    It("parses sample", func() {
      systemRepo := &system.FakeSystemRepository{
        LvsOutput: "  root:precise64:vwi-ao:80904.00::::::\n  swap_1:precise64:Iwi-ao:768.00::::::\n",
      }
      lvs, err := LogicalVolumes(systemRepo)
      Expect(err).To(BeNil())
      Expect(len(lvs)).To(Equal(2))
      Expect(lvs[0].LVName).To(Equal("root"))
      Expect(lvs[0].VolumeType).To(Equal(LVTVirtual))
      Expect(lvs[1].LVName).To(Equal("swap_1"))
      Expect(lvs[1].VolumeType).To(Equal(LVTMirrorImageOutOfSync))
    })
  })
})
