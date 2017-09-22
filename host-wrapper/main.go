package main

import (
	"log"
	"time"
)

const (
	PACKAGES_FOLDER = "packages"
	PACKAGES_DOWNLOADS_FOLDER = PACKAGES_FOLDER + "/" + "downloads" + "/"
	PACKAGES_SERVICES_FOLDER = PACKAGES_FOLDER + "/" + "services" + "/"
	LOGS_FOLDER = "logs"
)

var (
	hW HostWrapper
)

func main()  {
	log.Println("Starting")
	startup()

	hW = *new(HostWrapper)
	hW.init()

	err := hW.pM.startPackage("1234")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	defer hW.stop()
}

func startup()  {
	//build working directory
	create_fail := !createDir(PACKAGES_FOLDER)
	create_fail = (!createDir(PACKAGES_DOWNLOADS_FOLDER) || create_fail)
	create_fail = (!createDir(LOGS_FOLDER) || create_fail)

	if create_fail {
		failed(nil, "Failed to build working directory", 1)
	}
}
