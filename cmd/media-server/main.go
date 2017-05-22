package main

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jonathroth/media-server/lib/controller"
	"github.com/jonathroth/media-server/lib/serving"
)

func main() {
	handler := &controller.WinAPIHandler{}
	keyboard := &controller.MPCHCPlaybackController{}
	api := serving.NewMediaPlayerAPI(handler, keyboard)

	container := restful.NewContainer()
	api.Register(container)
	err := http.ListenAndServe(":8080", container)
	if err != nil {
		panic(err)
	}
}
