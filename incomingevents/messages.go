package incomingevents

import "github.com/jjeejj/go-tdlib/entities"

type GetChatHistoryResponse struct {
	Event

	TotalCount int64              `json:"total_count"`
	Messages   []entities.Message `json:"messages"`
}
