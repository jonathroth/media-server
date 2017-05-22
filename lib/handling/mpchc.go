package handling

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jonathroth/media-server/lib/controller"
)

const mpchcClassName = "MediaPlayerClassicW"

// PlaybackControlFunc is a function that changes the playback.
type PlaybackControlFunc func(controller.HWND) error

// MPCHCHandler handles HTTP requests by setting MPC-HC to focus and sending the required key press.
type MPCHCHandler struct {
	// Handler gets the MPC-HC window's handle to direct the commands to it.
	Handler controller.Handler

	// PlaybackControlFunc is the function that changes the playback as required.
	PlaybackControlFunc PlaybackControlFunc
}

// NewMPCHCHandler initializes a new MPC-HC handler.
func NewMPCHCHandler(focus controller.Handler, playbackControlFunc PlaybackControlFunc) *MPCHCHandler {
	return &MPCHCHandler{
		Handler:             focus,
		PlaybackControlFunc: playbackControlFunc,
	}
}

// Handle sets the MPC-HC window to focus and sends the required key.
func (h *MPCHCHandler) Handle(req *restful.Request, resp *restful.Response) {
	handle, err := h.Handler.Handle(mpchcClassName)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("Couldn't get MPCHC handle: %v", err))
		return
	}

	err = h.PlaybackControlFunc(handle)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("Couldn't send command to MPCHC: %v", err))
		return
	}
}
