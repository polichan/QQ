package wechat

import (
	"os"

	"github.com/polichan/qq/cache"
	"github.com/polichan/qq/miniprogram"
	miniConfig "github.com/polichan/qq/miniprogram/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// QQ struct
type QQ struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *QQ {
	return &QQ{}
}

// SetCache 设置cache
func (wc *QQ) SetCache(cahce cache.Cache) {
	wc.cache = cahce
}

// GetMiniProgram 获取小程序的实例
func (wc *QQ) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}
