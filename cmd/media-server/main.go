package main

import (
	"net/http"

	"flag"

	"github.com/emicklei/go-restful"
	"github.com/jonathroth/media-server/lib/controller"
	"github.com/jonathroth/media-server/lib/serving"
)

var host = flag.String("h", ":8080", "The server host & port, leave host empty to use all interfaces")

func main() {
	flag.Parse()

	handler := &controller.WinAPIHandler{}
	keyboard := &controller.MPCHCPlaybackController{}
	api := serving.NewMediaPlayerAPI(handler, keyboard)

	container := restful.NewContainer()
	api.Register(container)
	err := http.ListenAndServe(*host, container)
	if err != nil {
		panic(err)
	}
}
