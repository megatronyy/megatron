package app

import (
	"github.com/emicklei/go-restful"
	"github.com/twfx7758/megatron/cmd/mega-apiserver/app/routers"
	"net/http"
	"fmt"
	"dubbo-mesh/log"
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

	addr := fmt.Sprintf("%s:%d", r.host, r.port)

	fmt.Printf("mega-apiserver start listening at %s", addr)
	ser := &http.Server{Addr: addr, Handler: container}
	log.Fatal(ser.ListenAndServe())
}
