package system

type FakeSystemRepository struct {
  PvsOutput string
}

func (repo FakeSystemRepository) PVS() (output string, delimiter string, err error) {
  output = repo.PvsOutput
  delimiter = ":"
  return
}
