package main

import (
	"fmt"
	"os"
	"log"
)

type PackageManager struct {
	packages map[string]*Package
}

func (this *PackageManager) init() (error error) {
	log.Println("Package Manager init")
	//empty services dir
	os.Remove(PACKAGES_SERVICES_FOLDER)
	createDir(PACKAGES_SERVICES_FOLDER)

	packages, err := this.checkForPackages()
	this.packages = packages
	if err != nil {
		return fmt.Errorf("Failed to check for packages", err)
	}

	//init packages
	for _, p := range this.packages {
		err := p.Init()
		if err != nil {
			return fmt.Errorf("Package init Failed", err)
		}
	}

	return
}

func (this *PackageManager) startPackage(id string) (error error) {
	if p, ok := this.packages[id]; ok {
		err := p.Start()
		if err != nil {
			return fmt.Errorf("Starting package failed", err)
		}
		return
	} else {
		return fmt.Errorf("There is no package with id (tried to start it)", id)
	}
}

func (this *PackageManager) checkForPackages() (packages map[string]*Package, error error) {
	//check if package folder exist
	if !exist(PACKAGES_FOLDER) {
		error = fmt.Errorf("Packages folders does not exist")
		return
	}

	//scan for packages
	packageList, err := hW.mC.getPackageList()

	if err != nil {
		return packages, fmt.Errorf("Oops", err)
	}

	packages = make(map[string]*Package)
	for _, p := range packageList {
		packages[p.Id] = p
	}

	log.Println("Found", len(packages), "Package(s)")

	return
}

func (this *PackageManager) getNeededPackages() {

}
