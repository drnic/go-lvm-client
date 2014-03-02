package lvm_client

import (
  "errors"
  "strings"
  "strconv"
)

type VolumeGroup struct {
  VGName           string
  PhyiscalVolumes  int
  LogicalVolumes   int
  Attr             string
  Writable         bool
  Resizable        bool
  Exported         bool
  Partial          bool
  AllocationPolicy bool
  Clustered        bool
  VSize            float64
  VFree            float64
}

func NewVolumeGroup() VolumeGroup {
  return VolumeGroup{}
}

func (pv *VolumeGroup) ParseLine(vgsLine string, delimiter string) (err error) {
  var uint64Value uint64

tokens := strings.Split(strings.Trim(vgsLine, " "), delimiter)
  if (len(tokens) != 7) {
    err = errors.New("Expected 7 colon items from vgs. Perhaps an unsupported operating system.")
    return
  }
  pv.VGName = tokens[0]

  uint64Value, err = strconv.ParseUint(tokens[1], 10, 32)
  if (err != nil) {
    return err
  }
  pv.PhyiscalVolumes = int(uint64Value)

  uint64Value, err = strconv.ParseUint(tokens[2], 10, 32)
  if (err != nil) {
    return err
  }
  pv.LogicalVolumes = int(uint64Value)

  pv.Attr   = tokens[4]

  pv.VSize, err = strconv.ParseFloat(tokens[5], 32)
  if (err != nil) {
    return err
  }

  pv.VFree, err = strconv.ParseFloat(tokens[6], 32)
  if (err != nil) {
    return
  }

  return
}
