// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 84830579f0
// Version Date: Thu Dec 31 05:37:46 UTC 2020

package main

import (
	"flag"

	// This Service
	"videoWeb/verify-service/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	server.Run(server.DefaultConfig)
}
