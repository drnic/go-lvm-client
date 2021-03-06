package main

import (
	lvm "github.com/starkandwayne/go-lvm-client"
	"github.com/starkandwayne/go-lvm-client/system"
	"fmt"
)

func main() {
  repo := system.RealSystemRepository{}

	pvs, err := lvm.PhysicalVolumes(repo)
  if err != nil {
    fmt.Println(err.Error());
    return
  }
  fmt.Println("Physical Volumes:");
  fmt.Printf("%v\n", pvs)

	vgs, err := lvm.VolumeGroups(repo)
	if err != nil {
		fmt.Println(err.Error());
		return
	}
	fmt.Println("Volume Groups:");
	fmt.Printf("%v\n", vgs)

	lvs, err := lvm.LogicalVolumes(repo)
	if err != nil {
		fmt.Println(err.Error());
		return
	}
	fmt.Println("Logical Volumes:");
	fmt.Printf("%v\n", lvs)
}
