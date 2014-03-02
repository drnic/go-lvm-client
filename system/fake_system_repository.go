package system

type FakeSystemRepository struct {
  PvsOutput string
  VgsOutput string
}

func (repo FakeSystemRepository) PhyiscalVolumes() (output string, delimiter string, err error) {
  output = repo.PvsOutput
  delimiter = ":"
  return
}

func (repo FakeSystemRepository) VolumeGroups() (output string, delimiter string, err error) {
  output = repo.VgsOutput
  delimiter = ":"
  return
}
