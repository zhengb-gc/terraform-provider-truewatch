package api

import (
	"github.com/TrueWatchTech/terraform-provider-truewatch/internal/consts"
)

func init() {
	apiURLs[consts.TypeNameCustomRegion] = map[string]string{
		ResourceCreate: "/dialing_region/regist",
		ResourceDelete: "/dialing_region/%s/delete",
		ResourceRead:   "/dialing_region/%s/info",
	}
}
