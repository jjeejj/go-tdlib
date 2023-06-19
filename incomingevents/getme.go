package incomingevents

import "github.com/jjeejj/go-tdlib/entities"

type GetMe struct {
	Event

	*entities.User `json:"user"`
}
