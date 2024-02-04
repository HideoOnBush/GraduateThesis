package base

import (
	"GraduateThesis/conf"
	"GraduateThesis/platform"
)

type Base struct {
	*platform.Platform
}

func (s *Base) Init(c *conf.Config) {
	platform.Init(c)
	s.Platform = &platform.P
}
