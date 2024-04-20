package system_info

import (
	"fmt"
	"net"
	"os"

	"github.com/zcalusic/sysinfo"
)

var (
	errorMessage string = "<b>An error occured while trying to gather information about your system...</b>"
	pid          int    = os.Getpid()
	ipv4Address  []byte = *getOutboundIP()
)

// Must work, but can f up with unknown DNS settings
func getOutboundIP() *net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return new(net.IP)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return &localAddr.IP
}

func GetSystemInfo() string {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	var metrics string

	// Telebot instance info (host may not gather correctly)
	metrics += fmt.Sprintf(
		"Your <i>Telebot</i> instance is running on host: <b>%x</b> with PID: <b>%d</b>\n\n<b>System information:</b>\n",
		ipv4Address,
		pid,
	)

	// Tested only on Ubuntu 22.04 with WSL
	// WARN: May not work on non x86 or without superuser rights
	metrics += fmt.Sprintf(
		"<b>Hostname:</b> %s\n<b>Timezone:</b> %s\n<b>OS:</b> %s (%s - %s)\n<b>CPU:</b> %s (%d threads)\n\n",
		si.Node.Hostname,
		si.Node.Timezone,
		si.OS.Name,
		si.OS.Architecture,
		si.Kernel.Architecture,
		si.CPU.Model,
		si.CPU.Threads,
	)

	metrics += "<b>Network interfaces:</b>\n"
	networks := si.Network
	for _, inbound := range networks {
		metrics += fmt.Sprintf(
			"<b>Name: %s</b>\n<b>Driver:</b> %s\n<b>Mac:</b> %s\n\n",
			inbound.Driver,
			inbound.MACAddress,
			inbound.Name,
		)
	}

	return metrics
}
