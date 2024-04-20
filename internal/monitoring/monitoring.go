package monitoring

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
)

const (
	errorMessage string = "An <b>unknown error</b> was occured during metrics gaining...\nMake sure you running bot in <i>sudo</i>"
)

func GetMemoryLoad() string {
	memoryInfo, err := memory.Get()
	if err != nil {
		return errorMessage
	}

	return fmt.Sprintf(
		"<b>RAM load is</b>: %d <i>percents</i>\n<b>RAM used:</b> %.2f <i>gb</i>\n<b>RAM free:</b> %.2f <i>gb</i>\n<b>Swap load is:</b> %d <i>percents</i>\n<b>Swap used:</b> %.2f <i>gb</i>\n<b>Swap free:</b> %.2f <i>gb</i>",
		uint32(float64(memoryInfo.Used)/float64(memoryInfo.Total)*100),
		float64(memoryInfo.Used)/1_073_741_824.0,
		float64(memoryInfo.Inactive)/1_073_741_824.0,
		uint32(float64(memoryInfo.SwapUsed)/float64(memoryInfo.SwapTotal)*100),
		float64(memoryInfo.SwapUsed)/1_073_741_824.0,
		float64(memoryInfo.SwapFree)/1_073_741_824.0,
	)
}

func GetNetworkLoad() string {
	networksInfo, err := network.Get()
	if err != nil {
		return errorMessage
	}

	var metrics string
	for _, inbound := range networksInfo {
		metrics += fmt.Sprintf(
			"<b>Showing info for interface</b> <i>%s</i>\n<b>Outcome:</b> %d <i>bytes</i>\n<b>Income:</b> %d <i>bytes</i>\n\n",
			inbound.Name,
			inbound.RxBytes,
			inbound.TxBytes,
		)
	}
	if metrics == "" {
		return errorMessage
	}
	return metrics
}

func GetDiskLoad() string {
	disksInfo, err := disk.Get()
	if err != nil {
		return errorMessage
	}

	var metrics string
	for _, dir := range disksInfo {
		if dir.ReadsCompleted == 0 || dir.WritesCompleted == 0 {
			continue
		}
		metrics += fmt.Sprintf(
			"<b>Showing info for disk</b> <i>%s</i>\n<b>Writes:</b> <i>%d</i>\n<b>Reads:</b> <i>%d</i>\n\n",
			dir.Name,
			dir.WritesCompleted,
			dir.ReadsCompleted,
		)
	}
	if metrics == "" {
		return errorMessage
	}
	return metrics
}

func GetCpuLoad() string {
	cpuInfo, err := cpu.Get()
	if err != nil {
		return errorMessage
	}

	pp.Print(cpuInfo)

	return "Currently not implemented"
}
