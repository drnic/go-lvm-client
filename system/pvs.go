package system

import (
  "fmt"
  "os/exec"
)

type SystemRepository interface {
  PVS() (output string, delimiter string, err error)
}

type RealSystemRepository struct {
}

func (repo RealSystemRepository) PVS() (output string, delimiter string, err error) {
  delimiter = ":"
	cmd := exec.Command("pvs", "--units=m", "--separator=:", "--nosuffix", "--noheadings")
  out, err := cmd.Output()
  output = fmt.Sprintf("%s", out)
  if err != nil {
    return
  }
  return
}
