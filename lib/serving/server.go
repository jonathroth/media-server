package serving

import (
	"log"
	"os"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/jonathroth/media-server/lib/controller"
	"github.com/jonathroth/media-server/lib/handling"
)

var logger *log.Logger = log.New(os.Stdout, "", 0)

// RequestLogger logs HTTP requests to Stdout.
func RequestLogger(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	chain.ProcessFilter(req, resp)
	logger.Printf("%s - \"%s %s %s\" %d",
		time.Now().Format(time.Stamp),
		req.Request.Method,
		req.Request.RequestURI,
		req.Request.Proto,
		resp.StatusCode(),
	)
}

// MediaPlayerAPI defines all routes for the media player API.
type MediaPlayerAPI struct {
	Handler  controller.Handler
	Playback controller.MediaPlaybackController
}

// NewMediaPlayerAPI initializes a new media player API.
func NewMediaPlayerAPI(focus controller.Handler, keyboard controller.MediaPlaybackController) *MediaPlayerAPI {
	return &MediaPlayerAPI{
		Handler:  focus,
		Playback: keyboard,
	}
}

// Register adds all the media player API routes to the given container.
func (a *MediaPlayerAPI) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Filter(RequestLogger)

	ws.Route(ws.GET("/play").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.Play).Handle).
		Doc("Plays video playback"))

	ws.Route(ws.GET("/pause").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.Pause).Handle).
		Doc("Pauses video playback"))

	ws.Route(ws.GET("/playpause").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.PlayPause).Handle).
		Doc("Plays/Pauses video playback"))

	ws.Route(ws.GET("/stop").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.Stop).Handle).
		Doc("Stops video playback"))

	ws.Route(ws.GET("/next").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.Next).Handle).
		Doc("Goes to the next chapter"))

	ws.Route(ws.GET("/previous").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.Previous).Handle).
		Doc("Goes to the previous chapter"))

	ws.Route(ws.GET("/next-playlist").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.NextPlaylist).Handle).
		Doc("Goes to the next track"))

	ws.Route(ws.GET("/previous-playlist").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.PreviousPlaylist).Handle).
		Doc("Goes to the previous track"))

	ws.Route(ws.GET("/skip-forward").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.SkipForward).Handle).
		Doc("playback"))

	ws.Route(ws.GET("/skip-backward").
		To(handling.NewMPCHCHandler(a.Handler, a.Playback.SkipBack).Handle).
		Doc("Rewinds playback"))

	container.Add(ws)
}
