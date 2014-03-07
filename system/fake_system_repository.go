package system

type FakeSystemRepository struct {
  PvsOutput string
  VgsOutput string
  LvsOutput string
}

func (repo FakeSystemRepository) PhysicalVolumes() (output string, delimiter string, err error) {
  output = repo.PvsOutput
  delimiter = ":"
  return
}

func (repo FakeSystemRepository) VolumeGroups() (output string, delimiter string, err error) {
  output = repo.VgsOutput
  delimiter = ":"
  return
}

func (repo FakeSystemRepository) LogicalVolumes() (output string, delimiter string, err error) {
  output = repo.LvsOutput
  delimiter = ":"
  return
}
