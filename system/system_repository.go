package system

import (
  "fmt"
  "os/exec"
)

type SystemRepository interface {
  PhyiscalVolumes() (output string, delimiter string, err error)
  VolumeGroups() (output string, delimiter string, err error)
  LogicalVolumes() (output string, delimiter string, err error)
}

type RealSystemRepository struct {
}

func (repo RealSystemRepository) PhyiscalVolumes() (output string, delimiter string, err error) {
  delimiter = ":"
  cmd := exec.Command("pvs", "--units=m", "--separator=:", "--nosuffix", "--noheadings")
  out, err := cmd.Output()
  output = fmt.Sprintf("%s", out)
  if err != nil {
    return
  }
  return
}

func (repo RealSystemRepository) VolumeGroups() (output string, delimiter string, err error) {
  delimiter = ":"
  cmd := exec.Command("vgs", "--units=m", "--separator=:", "--nosuffix", "--noheadings")
  out, err := cmd.Output()
  output = fmt.Sprintf("%s", out)
  if err != nil {
    return
  }
  return
}

func (repo RealSystemRepository) LogicalVolumes() (output string, delimiter string, err error) {
  delimiter = ":"
  cmd := exec.Command("lvs", "--units=m", "--separator=:", "--nosuffix", "--noheadings")
  out, err := cmd.Output()
  output = fmt.Sprintf("%s", out)
  if err != nil {
    return
  }
  return
}
