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
