package incomingevents

import "github.com/jjeejj/go-tdlib/entities"

type GetRemoteFile struct {
	Event

	*entities.File `json:"file"`
}
