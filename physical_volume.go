package lvm_client

import "strings"

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

func NewPhysicalVolume() PhysicalVolume {
	return PhysicalVolume{}
}

func (pv *PhysicalVolume) ParseDisplayWithColons(pvdisplayWithColons string) (err error) {
  tokens := strings.Split(pvdisplayWithColons, ":")
  pv.PVName = tokens[0]
/*  VGName: tokens[1],
  PVSize: 0,
  Allocatable: true,
  PESize: 0,
  TotalPE: 0,
  FreePE: 0,
  AllocatedPE: 0,
  UUID: "xxx",*/
  return
}

func PhysicalVolumes() (pvs []PhysicalVolume, err error) {
  pvs = []PhysicalVolume{
    {
      PVName: "foo",
    },
  }
  return
}
