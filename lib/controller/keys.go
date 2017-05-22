package controller

import (
	"errors"
)

var (
	procSendMessageW = user32.NewProc("SendMessageW")
)

const (
	WM_COMMAND = 0x0111

	// Commands taken from https://autohotkey.com/board/topic/73687-media-player-classic-home-cinema-function-libraries/
	mpchcPlay              = 887
	mpchcPause             = 888
	mpchcPlayPause         = 889
	mpchcStop              = 890
	mpchcJumpBackwardSmall = 899
	mpchcJumpForwardSmall  = 900
	mpchcPreviousPlaylist  = 919
	mpchcNextPlaylist      = 920
	mpchcPrevious          = 921
	mpchcNext              = 922
)

// MPCHCPlaybackController controls media playback through Windows API 'WM_COMMAND' messages to MPC-HC.
type MPCHCPlaybackController struct{}

// Play sends a playback play instruction.
func (c *MPCHCPlaybackController) Play(handle HWND) error {
	return c.sendCommand(handle, mpchcPlay)
}

// Pause sends a playback pause instruction.
func (c *MPCHCPlaybackController) Pause(handle HWND) error {
	return c.sendCommand(handle, mpchcPause)
}

// PlayPause send a playback play/pause instruction.
func (c *MPCHCPlaybackController) PlayPause(handle HWND) error {
	return c.sendCommand(handle, mpchcPlayPause)
}

// Stop send a playback stop instruction.
func (c *MPCHCPlaybackController) Stop(handle HWND) error {
	return c.sendCommand(handle, mpchcStop)
}

// Previous send a back chapter instruction.
func (c *MPCHCPlaybackController) Previous(handle HWND) error {
	return c.sendCommand(handle, mpchcPrevious)
}

// Next send a forward chapter instruction.
func (c *MPCHCPlaybackController) Next(handle HWND) error {
	return c.sendCommand(handle, mpchcNext)
}

// Previous send a back entire track instruction.
func (c *MPCHCPlaybackController) PreviousPlaylist(handle HWND) error {
	return c.sendCommand(handle, mpchcPreviousPlaylist)
}

// Next send a forward entire track instruction.
func (c *MPCHCPlaybackController) NextPlaylist(handle HWND) error {
	return c.sendCommand(handle, mpchcNextPlaylist)
}

// SkipForward sends a jump forward instruction.
func (c *MPCHCPlaybackController) SkipForward(handle HWND) error {
	return c.sendCommand(handle, mpchcJumpForwardSmall)
}

// SkipForward sends a jump back instruction.
func (c *MPCHCPlaybackController) SkipBack(handle HWND) error {
	return c.sendCommand(handle, mpchcJumpBackwardSmall)
}

// sendCommand sends an instruction
func (c *MPCHCPlaybackController) sendCommand(handle HWND, key uintptr) error {
	ret := SendMessageW(handle, WM_COMMAND, key, 0)
	if ret != 1 {
		return errors.New("Failed to send message to handle")
	}

	return nil
}

// SendMessageW wraps the call to the user32.dll SendMessageW.
// It's taken (rather than imported) from github.com/AllenDang/w32 to prevent gcc compilation.
func SendMessageW(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendMessageW.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}
