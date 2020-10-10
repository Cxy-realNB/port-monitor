package devices

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"golang.org/x/net/websocket"
	"main/clog"
	"main/constant"
	"time"
)

func ListenAll() {
	for _, device := range *constant.DevicesList {
		go ListenDev(device)
	}
}

func ListenDev(dev pcap.Interface) {
	clog.InfoLogger.Println("Listening " + dev.Name)
	handle, err := pcap.OpenLive(dev.Name, 1024, false, 30*time.Second)
	if err != nil {
		clog.ErrorLogger.Println(err)
		handle.Close()
	} else {
		constant.Handles[dev.Name] = handle
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			for _, v := range constant.Socks {
				if v.DevName == dev.Name {
					// todo
					websocket.Message.Send(v.Conn, packet.String())
				}
			}
			clog.InfoLogger.Println(packet.String())
		}
	}
}
