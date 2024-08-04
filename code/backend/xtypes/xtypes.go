package xtypes

import (
	"io/fs"

	"github.com/gin-gonic/gin"
)

type App interface {
	Start(string) error
	Stop() error

	GetDatabase() any
	GetSockd() any
	GetEventBus() any
	GetSequencer() any
	GetFileStore() any
}

type AppType struct {
	Name string
	Slug string
	Info string
	Icon string

	// called when mounting api route group
	OnMount func(ctx *MountContext) error
	OnStart func() error
	OnStop  func(force bool) error

	OnPageRequest func(ctx *gin.Context) (bool, error)
	OnFileRequest func(ctx *gin.Context) (bool, error)
	// only called if this app is selected as root app
	OnRootMount func() *gin.Engine

	OnAppInit func(id int64) error

	// default template /z/apps/<app_slug>/<project_id>
	AppProjectURLTpl string

	EventTypes  []string
	Permissions []string

	AssetData fs.FS
	GlobalJS  string
	GlobalCSS string
}

type MountContext struct {
	App           App
	ApiRouteGroup *gin.RouterGroup
}

// /z/portal
// /z/apps/<app_slug>
// /z/api/apps/<app_slug>
// /z/api/portal
