package main

type HostWrapper struct {
	mC MasterCommunicator
	pM PackageManager
}

func (this *HostWrapper) init()  {
	this.mC = *new(MasterCommunicator)
	this.mC.init()

	this.pM = *new(PackageManager)
	err := this.pM.init()
	if err != nil {
		panic(err)
	}


}

func (this HostWrapper) stop()  {
	this.mC.stop()
}
