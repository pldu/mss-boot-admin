package config

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/8/6 08:11:42
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/8/6 08:11:42
 */

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Mode string

const (
	ModeDev  Mode = "dev"
	ModeTest Mode = "test"
	ModeProd Mode = "prod"
)

type Application struct {
	Name       string            `yaml:"name" json:"name"`
	Mode       Mode              `yaml:"mode" json:"mode"`
	Origin     string            `yaml:"origin" json:"origin"`
	StaticPath map[string]string `yaml:"staticPath" json:"staticPath"`
	Labels     map[string]string `yaml:"labels" json:"labels"`
}

func (e *Application) Init(r gin.IRouter) {
	if e.Mode == "" {
		e.Mode = ModeDev
	}

	switch e.Mode {
	case ModeDev:
		// set gin mode
		gin.SetMode(gin.DebugMode)

		// set static path
		for k := range e.StaticPath {
			//if k == "404" {
			//	r.NoRoute(func(c *gin.Context) {
			//		c.File(e.StaticPath[k])
			//	})
			//	continue
			//}
			if filepath.Ext(k) != "" {
				r.StaticFile(k, e.StaticPath[k])
				continue
			}
			r.Static(k, e.StaticPath[k])
		}
		r.StaticFile("/swagger.json", "docs/swagger.json")
		r.StaticFile("/swagger.yaml", "docs/swagger.yaml")
	case ModeTest:
		// set gin mode
		gin.SetMode(gin.TestMode)
		// no static
	case ModeProd:
		// set gin mode
		gin.SetMode(gin.ReleaseMode)
		// no static
	}
}
