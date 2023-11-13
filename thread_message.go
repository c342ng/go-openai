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

type ThreadMessageResp struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	CreatedAt int    `json:"created_at"`
	ThreadId  string `json:"thread_id"`
	Role      string `json:"role"`
	Content   []struct {
		Type string `json:"type"`
		Text struct {
			Value       string        `json:"value"`
			Annotations []interface{} `json:"annotations"`
		} `json:"text"`
	} `json:"content"`
	FileIds     []interface{} `json:"file_ids"`
	AssistantId interface{}   `json:"assistant_id"`
	RunId       interface{}   `json:"run_id"`
	Metadata    struct {
	} `json:"metadata"`
}

type ThreadMessageResponse struct {
	ThreadMessageResp
	httpHeader
}

type ThreadMessagesResponse struct {
	Object string              `json:"object"`
	Data   []ThreadMessageResp `json:"data"`
	httpHeader
}

// CreateThreadMessage creates a new thread message
func (c *Client) CreateThreadMessage(ctx context.Context, threadID string, request ThreadMessageRequest) (response ThreadMessageResponse, err error) {
	urlSuffix := fmt.Sprint("%s/%s/messages", threadsSuffix, threadID)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix), withBody(request),
		withBetaAssistantV1())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// CreateThreadMessage creates a new thread message
func (c *Client) ListThreadMessage(ctx context.Context, threadID string) (response ThreadMessagesResponse, err error) {
	urlSuffix := fmt.Sprint("%s/%s/messages", threadsSuffix, threadID)
	req, err := c.newRequest(ctx, http.MethodGet, c.fullURL(urlSuffix), withBetaAssistantV1())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
