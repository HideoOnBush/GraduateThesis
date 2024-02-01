package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"GraduateThesis/biz/handler"
	"GraduateThesis/biz/router"
)

// Register registers all routers.
func Register(r *server.Hertz) {
	router.GeneratedRegister(r)
	customizedRegister(r)
}

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
}
