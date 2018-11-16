package app

import (
	"github.com/emicklei/go-restful"
	"github.com/twfx7758/megatron/cmd/mega-apiserver/app/routers"
)

type Router struct {
	host string
	port int
}

func NewRouter(host string, port int) *Router {
	return &Router{
		host: host,
		port: port,
	}
}

func (r *Router) Start() {
	container := restful.NewContainer()
	container.Router(restful.CurlyRouter{})

	user := new(routers.UserController)
	user.Register(container)
}
