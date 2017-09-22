package main

import (
	"log"
	"fmt"
	"clustered/master-hwrapper"
	"clustered/communicator"
)

type MasterCommunicator struct {
	c *communicator.Client
}

func (this *MasterCommunicator) init() {
	log.Println("Master Communicator init")
	c := master_hwrapper.InitAsClient("127.0.0.1:1234")
	this.c = c
}

func (this *MasterCommunicator) stop() {
	this.c.Stop()
}

func (this *MasterCommunicator) getPackageList() (packages []*Package, error error) {
	res, err := this.c.Dc.Call(master_hwrapper.GET_PACKAGE_LIST, &master_hwrapper.PackageList{})
	if err != nil {
		error = fmt.Errorf("Package list call failed", err)
		return
	}

	for _, p := range res.(master_hwrapper.PackageList) {
		packages = append(packages, &Package{Id: p.Id, Name: p.Name, Version: p.Version})
	}
	return
}
