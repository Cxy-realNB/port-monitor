package devices

import (
	"encoding/json"
	"github.com/google/gopacket/pcap"
	"main/clog"
	"main/constant"
)

func GetDevices() (ifs *[]pcap.Interface) {
	if constant.DevicesList == nil {
		devices, err := pcap.FindAllDevs()
		if err != nil {
			clog.ErrorLogger.Println(err)
		}
		constant.DevicesList = &devices
	}
	return constant.DevicesList
}

func GetDevicesJson() (result string) {
	devices := GetDevices()
	jsonByte, err := json.Marshal(devices)
	if err != nil {
		clog.ErrorLogger.Println(err)
	}
	return string(jsonByte)
}

func GetDevicesByName(name string) (dev pcap.Interface, err string) {
	var devR *pcap.Interface
	var errR string
	for _, d := range *constant.DevicesList {
		if d.Name == name {
			devR = &d
			break
		}
	}

	if devR == nil {
		errR = "Do not find the device: " + name
		devR = &pcap.Interface{}
	}
	return *devR, errR
}
