package controller

type HWND uintptr

// A Handler gets the handle of a process using it's class name.
type Handler interface {
	Handle(processClassName string) (HWND, error)
}

// A MediaPlaybackControllers controls the playback for media.
type MediaPlaybackController interface {
	// Play sends a playback play instruction.
	Play(HWND) error

	// Pause sends a playback pause instruction.
	Pause(HWND) error

	// PlayPause send a playback play/pause instruction.
	PlayPause(HWND) error

	// Stop send a playback stop instruction.
	Stop(HWND) error

	// Previous send a back chapter instruction.
	Previous(HWND) error

	// Next send a forward chapter instruction.
	Next(HWND) error

	// Previous send a back entire track instruction.
	PreviousPlaylist(HWND) error

	// Next send a forward entire track instruction.
	NextPlaylist(HWND) error

	// SkipForward sends a jump forward instruction.
	SkipForward(HWND) error

	// SkipForward sends a jump back instruction.
	SkipBack(HWND) error
}
