package system

type SystemRepository interface {
  PVS() (output string, delimiter string, err error)
}

type RealSystemRepository struct {
}

func (repo RealSystemRepository) PVS() (output string, delimiter string, err error) {
  output = "  /dev/sda5:precise64:lvm2:a-:81672.00:0"
  delimiter = ":"
  return
}
