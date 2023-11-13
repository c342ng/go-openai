package openai

import (
	"context"
	"fmt"
	"net/http"
)

type ThreadMessageRequest struct {
	Role     ThreadMessageRole `json:"role"`
	Content  string            `json:"content"`
	FileIDs  []string          `json:"file_ids,omitempty"`
	Metadata map[string]any    `json:"metadata,omitempty"`
}

// CreateThreadMessage creates a new thread message
func (c *Client) CreateThreadMessage(ctx context.Context, threadID string, request ThreadMessageRequest) (response Thread, err error) {
	urlSuffix := fmt.Sprint("%s/%s/messages", threadsSuffix, threadID)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix), withBody(request),
		withBetaAssistantV1())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
