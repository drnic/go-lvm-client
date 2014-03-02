package lvm_client

import (
  "errors"
  "strings"
  "strconv"
  "github.com/starkandwayne/go-lvm-client/system"
)

type LogicalVolumeType int

const (  // iota is reset to 0
  Unspecified LogicalVolumeType = iota
  Mirrored
  MirroredWithoutInitialSync
  Origin
  OriginWithMergingSnapshot
  Snapshot
  MergingSnapshot
  Pvmove
  Virtual
  MirrorImage
  MirrorImageOutOfSync
  UnderConversion
)

type LogicalVolume struct {
  LVName         string
  VGName         string
  Attrs          string
  VolumeType     LogicalVolumeType
  LVSize         float64
}

func NewLogicalVolume() LogicalVolume {
  return LogicalVolume{}
}

func (lv *LogicalVolume) ParseLine(lvsLine string, delimiter string) (err error) {
  tokens := strings.Split(strings.Trim(lvsLine, " "), delimiter)
  if (len(tokens) != 10) {
    err = errors.New("Expected 10 colon items from lvs. Perhaps an unsupported operating system.")
    return
  }
  lv.LVName = tokens[0]
  lv.VGName = tokens[1]
  lv.Attrs  = tokens[2]
  lv.parseAttr()

  lv.LVSize, err = strconv.ParseFloat(tokens[3], 32)
  if (err != nil) {
    return err
  }

  return
}

func (vg *LogicalVolume) parseAttr() {
  attrs := strings.Split(vg.Attrs, "")
  switch attrs[0] {
    case "-": vg.VolumeType = Unspecified
    case "m": vg.VolumeType = Mirrored
    case "M": vg.VolumeType = MirroredWithoutInitialSync
    case "o": vg.VolumeType = Origin
    case "O": vg.VolumeType = OriginWithMergingSnapshot
    case "s": vg.VolumeType = Snapshot
    case "S": vg.VolumeType = MergingSnapshot
    case "p": vg.VolumeType = Pvmove
    case "v": vg.VolumeType = Virtual
    case "i": vg.VolumeType = MirrorImage
    case "I": vg.VolumeType = MirrorImageOutOfSync
    case "c": vg.VolumeType = UnderConversion
  }
}

func LogicalVolumes(repo system.SystemRepository) (lvs []LogicalVolume, err error) {
  lvsOutput, delimiter, err := repo.LogicalVolumes()
  lvs = []LogicalVolume{}
  lvsLines := strings.Split(lvsOutput, "\n")
  for _, lvLine := range lvsLines {
    if len(lvLine) > 0 {
      lv := NewLogicalVolume()
      err = lv.ParseLine(lvLine, delimiter)
      if err != nil {
        return
      }
      lvs = append(lvs, lv)
    }
  }

  return
}
