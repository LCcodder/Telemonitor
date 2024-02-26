package monitoring

import (
	"fmt"

	"github.com/mackerelio/go-osstat/memory"
	//"github.com/mackerelio/go-osstat/network"
)

func GetMemoryLoad() string {
	memoryInfo, err := memory.Get()
	if err != nil {
		return "An <b>unknown error</b> was occured during metrics gaining...\nMake sure you running bot in <i>sudo</i>"
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

// func GetNetworkLoad() string {
// 	networkInfo, err := network.Get()
// 	if err != nil {
//     	return "*An unknown error was occured during metrics gaining...\nMake sure you running bot in* `sudo`"
//     }

//     return fmt.Sprintf(
//         networkInfo.
//     )

// }

// func GetDiskLoad() string {
// 	diskInfo, err := disk.Get()
// 	if err != nil {
// 		return "*An unknown error was occured during metrics gaining...\nMake sure you running bot in* `sudo`"
// 	}

// }
