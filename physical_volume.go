package lvm_client

import (
  "errors"
  "strings"
  "strconv"
)

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
  var value uint64

  tokens := strings.Split(pvdisplayWithColons, ":")
  if (len(tokens) != 12) {
    err = errors.New("Expected 12 colon items from pvdisplay")
    return
  }
  pv.PVName = tokens[0]
  pv.VGName = tokens[1]

  value, err = strconv.ParseUint(tokens[2], 10, 32)
  if (err != nil) {
    return
  }
  pv.PVSize = int(value)
  pv.UUID = tokens[11]
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
