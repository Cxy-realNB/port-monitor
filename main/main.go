package main

import (
	"main/argsparse"
	"main/clog"
	"main/devices"
	"main/web"
	"net/http"
)

func init() {
	devices.GetDevices()
}

func main() {
	argsparse.ParseArgs()
	temp := argsparse.Args
	clog.InitLog(temp.LogPath)
	clog.InfoLogger.Println(devices.GetDevicesJson())
	go devices.ListenAll()
	clog.InfoLogger.Println("Initialize successfully!")
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	web.InitServer(temp.IP, temp.Port, temp.LogReq)

}
