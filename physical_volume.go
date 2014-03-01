package lvm_client

type PhysicalVolume struct {
  PVName         string
  VGName         string
  PVSize         int
  Allocatable    bool
  PESize         int
  TotalPE        int
  FreePE         int
  AllocatedPE    int
  UUID           string
}

func PhysicalVolumes() (pvs []PhysicalVolume, err error) {
  pvs = []PhysicalVolume{
    {
      PVName: "foo",
    },
  }
  return
}
