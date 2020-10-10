package constant

import "github.com/google/gopacket/pcap"

var Handles = make(map[string]*pcap.Handle)

var DevicesList *[]pcap.Interface
