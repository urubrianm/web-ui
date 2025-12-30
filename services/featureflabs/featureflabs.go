package featureflags

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

const (
	DisableDownloadsEnv = "WEBUI_DISABLE_DOWNLOADS"
	DisableAdsEnv       = "WEBUI_DISABLE_ADS"
	DisableLogsEnv      = "WEBUI_DISABLE_LOGS"

	DisableDownloadsFlag = "disable-downloads"
	DisableAdsFlag       = "disable-ads"
	DisableLogsFlag      = "disable-logs"

	ctxKey = "featureflags"
)

type Flags struct {
	DisableDownloads bool
	DisableAds       bool
	DisableLogs      bool
}

func RegisterFlags(f []cli.Flag) []cli.Flag {
	return append(f,
		cli.BoolFlag{
			Name:   DisableDownloadsFlag,
			Usage:  "Disable downloads (hide UI + block endpoints)",
			EnvVar: DisableDownloadsEnv,
		},
		cli.BoolFlag{
			Name:   DisableAdsFlag,
			Usage:  "Disable ads (do not render ads step)",
			EnvVar: DisableAdsEnv,
		},
		cli.BoolFlag{
			Name:   DisableLogsFlag,
			Usage:  "Disable progress logs UI (still keeps template renders working)",
			EnvVar: DisableLogsEnv,
		},
	)
}

func FromCLI(c *cli.Context) Flags {
	return Flags{
		DisableDownloads: c.Bool(DisableDownloadsFlag),
		DisableAds:       c.Bool(DisableAdsFlag),
		DisableLogs:      c.Bool(DisableLogsFlag),
	}
}

func Middleware(flags Flags) gin.HandlerFunc {
	return func(c *gin.Context) {
		// attach to gin context (so web.NewContext can pick it up)
		c.Set(ctxKey, flags)

		// hard-block downloads endpoints when disabled
		if flags.DisableDownloads {
			p := strings.ToLower(c.Request.URL.Path)
			if p == "/download-file" || p == "/download-dir" || p == "/download-torrent" {
				c.AbortWithStatus(403)
				return
			}
		}

		c.Next()
	}
}

func Get(c *gin.Context) Flags {
	if c == nil {
		return Flags{}
	}
	if v, ok := c.Get(ctxKey); ok {
		if ff, ok2 := v.(Flags); ok2 {
			return ff
		}
	}
	return Flags{}
}
