package main

import (
	"fmt"
	"os"
	"log"
	"os/exec"
	"bufio"
	"context"
)

type Package struct {
	Name        string
	Id          string
	Version     string
	folder      string
	startable   bool
	packedFile  string
	cmd         *exec.Cmd
	cmdOut      *bufio.Reader
	running     bool
	cancel      context.CancelFunc
	process		*os.Process
	processInfo *ProcessInfo
}

type ProcessInfo struct {
	stdout string

}

func (this *Package) Init() (error error) {
	log.Println("Init package:", this.Name, "version:", this.Version)
	this.folder = this.Name + "/" + this.Version

	//build service
	err := this.build()
	if err != nil {
		return fmt.Errorf("Build failed", err)
	}

	this.startable = true;
	this.running = false;
	return
}

func (this *Package) Start() (error error) {
	log.Println("Starting package:", this.Name, "Version:", this.Version)
	if !this.startable {
		return fmt.Errorf("This package is not startable")
	}

	ctx, cancel := context.WithCancel(context.Background())
	this.cancel = cancel
	this.cmd = exec.CommandContext(ctx, "ping", "bremersklaver.nl")
	stdout, err := this.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Failed to get stdout", err)
	}

	this.cmd.Stdout = os.Stdout
	this.cmdOut = bufio.NewReader(stdout)

	err = this.cmd.Start()
	if err != nil {
		return fmt.Errorf("Failed to start service", err)
	}

	this.process, err = os.FindProcess(this.cmd.Process.Pid)
	if err != nil {
		return fmt.Errorf("failed to get packet process with pid", err)
	}

	this.wrapper()

	this.running = true;
	return
}

func (this *Package) wrapper() (error error) {
	this.processInfo = new(ProcessInfo)
	for {
		fmt.Println(this.process)
		bytes := make([]byte, this.cmdOut.Buffered())
		n, err := this.cmdOut.Read(bytes)

		if err != nil {
			return fmt.Errorf("Reading from stdout of package process", err)
		}
		fmt.Print(string(bytes[:n]))
		this.processInfo.stdout += string(bytes[:n])
		//this.processInfo = append(this.processInfo, ProcessInfo{
		//
		//})
	}
	this.running = false;
	return
}

func (this *Package) kill() {
	this.cmd.Process.Kill()
}

func (this *Package) build() (error error) {
	//check if file exist
	this.packedFile = PACKAGES_DOWNLOADS_FOLDER + this.Name + "-" + this.Version + ".zip"
	_, err := os.Stat(this.packedFile)
	if os.IsNotExist(err) {
		//download file if not exist
		err = this.download()
		if err != nil {
			return fmt.Errorf("Download failed", err)
		}
	}

	//unpack file
	err = this.unpack()
	if err != nil {
		return fmt.Errorf("unpacking failed", err)
	}

	return
}

func (this *Package) unpack() (error error) {
	if !createDir(PACKAGES_SERVICES_FOLDER + this.folder) {
		fmt.Println(PACKAGES_SERVICES_FOLDER + this.folder)
		return fmt.Errorf("Package folder could not be created")
	}

	err := unzip(this.packedFile, PACKAGES_SERVICES_FOLDER+this.folder)
	if err != nil {
		return fmt.Errorf("Failed to unzip service", err)
	}

	return
}

func (this *Package) download() (err error) {
	return fmt.Errorf("Had to download shit")
}
