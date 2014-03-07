package lvm_client

import (
  "errors"
  "strings"
  "strconv"
  "github.com/starkandwayne/go-lvm-client/system"
)

type VolumeGroup struct {
  VGName           string
  PhysicalVolumes  int
  LogicalVolumes   int
  Attrs            string
  Writable         bool
  Resizable        bool
  Exported         bool
  Partial          bool
  AllocationPolicy string
  Clustered        bool
  VSize            float64
  VFree            float64
}

func NewVolumeGroup() VolumeGroup {
  return VolumeGroup{}
}

func (vg *VolumeGroup) ParseLine(vgsLine string, delimiter string) (err error) {
  var uint64Value uint64

  tokens := strings.Split(strings.Trim(vgsLine, " "), delimiter)
  if (len(tokens) != 7) {
    err = errors.New("Expected 7 colon items from vgs. Perhaps an unsupported operating system.")
    return
  }
  vg.VGName = tokens[0]

  uint64Value, err = strconv.ParseUint(tokens[1], 10, 32)
  if (err != nil) {
    return err
  }
  vg.PhysicalVolumes = int(uint64Value)

  uint64Value, err = strconv.ParseUint(tokens[2], 10, 32)
  if (err != nil) {
    return err
  }
  vg.LogicalVolumes = int(uint64Value)

  vg.Attrs = tokens[4]
  vg.parseAttr()

  vg.VSize, err = strconv.ParseFloat(tokens[5], 32)
  if (err != nil) {
    return err
  }

  vg.VFree, err = strconv.ParseFloat(tokens[6], 32)
  if (err != nil) {
    return
  }


  return
}

func (vg *VolumeGroup) parseAttr() {
  attrs := strings.Split(vg.Attrs, "")
  vg.Writable = attrs[0] == "w"
  vg.Resizable = attrs[1] == "z"
  vg.Exported = attrs[2] == "x"
  vg.Partial = attrs[3] == "p"
  vg.AllocationPolicy = attrs[4]
  vg.Clustered = attrs[5] == "c"
}

func VolumeGroups(repo system.SystemRepository) (vgs []VolumeGroup, err error) {
  vgsOutput, delimiter, err := repo.VolumeGroups()
  vgs = []VolumeGroup{}
  vgsLines := strings.Split(vgsOutput, "\n")
  for _, vgLine := range vgsLines {
    if len(vgLine) > 0 {
      vg := NewVolumeGroup()
      err = vg.ParseLine(vgLine, delimiter)
      if err != nil {
        return
      }
      vgs = append(vgs, vg)
    }
  }

  return
}
