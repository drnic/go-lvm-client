package lvm_client_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/starkandwayne/go-lvm-client"
  "github.com/starkandwayne/go-lvm-client/system"
)

var _ = Describe("VolumeGroup", func() {
  Describe("parse colon output", func() {
    It("new from colon output", func() {
      vg := NewVolumeGroup()
      err := vg.ParseLine("  precise64:1:2:0:wz--n-:81672.00:0", ":")
      Expect(err).To(BeNil())
      Expect(vg.VGName).To(Equal("precise64"))
      Expect(vg.PhyiscalVolumes).To(Equal(1))
      Expect(vg.LogicalVolumes).To(Equal(2))
      /* unknown '0' in token 4 */
      Expect(vg.Attrs).To(Equal("wz--n-"))
      /*
      The vg_attr bits are:
        1  Permissions: (w)riteable, (r)ead-only
        2  Resi(z)eable
        3  E(x)ported
        4  (p)artial: one or more physical volumes belonging to the volume group are missing from the system
        5  Allocation policy: (c)ontiguous, c(l)ing, (n)ormal, (a)nywhere, (i)nherited
        6  (c)lustered */
      Expect(vg.Writable).To(BeTrue())
      Expect(vg.Resizable).To(BeTrue())
      Expect(vg.Exported).ToNot(BeTrue())
      Expect(vg.Partial).ToNot(BeTrue())
      Expect(vg.AllocationPolicy).To(Equal("n"))
      Expect(vg.Clustered).ToNot(Equal(BeTrue()))

      Expect(vg.VSize).To(Equal(81672.0))
      Expect(vg.VFree).To(Equal(0.0))
    })


    It("invalid number of tokens", func() {
      pv := NewVolumeGroup()
      err := pv.ParseLine("x:y:z", ":")
      Expect(err).ToNot(BeNil())
    })
  })

  Describe("parse vgs output", func() {
    It("parses sample", func() {
      systemRepo := &system.FakeSystemRepository{
        VgsOutput: "  precise64:1:2:0:wz--n-:81672.00:0",
      }
      vgs, err := VolumeGroups(systemRepo)
      Expect(err).To(BeNil())
      Expect(len(vgs)).To(Equal(1))
      Expect(vgs[0].VGName).To(Equal("precise64"))
    })
  })

})
