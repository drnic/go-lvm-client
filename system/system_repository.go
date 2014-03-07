package system

import (
  "fmt"
  "os/exec"
)

type SystemRepository interface {
  PhysicalVolumes() (output string, delimiter string, err error)
  VolumeGroups() (output string, delimiter string, err error)
  LogicalVolumes() (output string, delimiter string, err error)
}

type RealSystemRepository struct {
}

func (repo RealSystemRepository) PhysicalVolumes() (output string, delimiter string, err error) {
  delimiter = ":"
  output, err = repo.runCommand("pvs", "--units=m", "--separator=:", "--nosuffix", "--noheadings")
  return
}

func (repo RealSystemRepository) VolumeGroups() (output string, delimiter string, err error) {
  delimiter = ":"
  output, err = repo.runCommand("vgs", "--units=m", "--separator=:", "--nosuffix", "--noheadings")
  return
}

func (repo RealSystemRepository) LogicalVolumes() (output string, delimiter string, err error) {
  delimiter = ":"
  output, err = repo.runCommand("lvs", "--units=m", "--separator=:", "--nosuffix", "--noheadings")
  return
}

func (repo RealSystemRepository) runCommand(name string, args ...string) (output string, err error) {
  cmd := exec.Command(name, args...)
  out, err := cmd.Output()
  output = fmt.Sprintf("%s", out)
  if err != nil {
    return
  }
  return
}
