package whitelist

import (
	"slices"
	"telemonitor/config"
)

func IsInWhitelist(username *string) bool {
	return slices.Contains(config.Cfg.AuthParams.Whitelist, *username)
}

func IsAdmin(username *string) bool {
	return config.Cfg.AuthParams.Admin == *username
}

func AddToWhitelist(username *string) bool {
	if IsInWhitelist(username) {
		return false
	}

	config.Cfg.AuthParams.Whitelist = append(config.Cfg.AuthParams.Whitelist, *username)
	status, _ := config.WriteConfig(config.Cfg)
	return status
}

func RemoveFromWhitelist(username *string) bool {
	// based, whitelist already sorted
	index, isFound := slices.BinarySearch(config.Cfg.AuthParams.Whitelist, *username)
	if !isFound {
		return false
	}

	// a bit slow, but it totaly lost on me :P
	config.Cfg.AuthParams.Whitelist = append(config.Cfg.AuthParams.Whitelist[:index], config.Cfg.AuthParams.Whitelist[index+1:]...)

	status, _ := config.WriteConfig(config.Cfg)
	return status
}
