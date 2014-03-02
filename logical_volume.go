package lvm_client

import (
  "errors"
  "strings"
  "strconv"
  "github.com/starkandwayne/go-lvm-client/system"
)

type LogicalVolume struct {
  LVName         string
  VGName         string
  Attr           string
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
  lv.Attr   = tokens[2]

  lv.LVSize, err = strconv.ParseFloat(tokens[3], 32)
  if (err != nil) {
    return err
  }

  return
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
