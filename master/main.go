package main

import (
	"log"
	"io"
	"clustered/master-hwrapper"
)

func main() {
	s := master_hwrapper.InitAsServer("127.0.0.1:1234")
	s.S.OnConnect = onConnect
	log.Println("server started")
	select {}
}

func onConnect(remoteAddr string, rwc io.ReadWriteCloser) (io.ReadWriteCloser, error) {
	log.Println("Connected to", remoteAddr)
	return rwc, nil
}