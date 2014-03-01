package lvm_client

import (
  "errors"
  "strings"
  "strconv"
)

type PhysicalVolume struct {
  PVName         string
  VGName         string
  Format         string
  Attr           string
  PVSize         float64
  FreePE         float64
}

func NewPhysicalVolume() PhysicalVolume {
	return PhysicalVolume{}
}

func (pv *PhysicalVolume) ParseLine(pvdisplayWithColons string, delimiter string) (err error) {
  tokens := strings.Split(strings.Trim(pvdisplayWithColons, " "), delimiter)
  if (len(tokens) != 6) {
    err = errors.New("Expected 6 colon items from pvs")
    return
  }
  pv.PVName = tokens[0]
  pv.VGName = tokens[1]
  pv.Format = tokens[2]
  pv.Attr   = tokens[3]

  pv.PVSize, err = strconv.ParseFloat(tokens[4], 32)
  if (err != nil) {
    return err
  }

  pv.FreePE, err = strconv.ParseFloat(tokens[5], 32)
  if (err != nil) {
    return
  }

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
