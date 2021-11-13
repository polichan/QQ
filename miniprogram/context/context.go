package context

import (
	"github.com/polichan/qq/credential"
	"github.com/polichan/qq/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
