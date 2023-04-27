package main

import (
	"context"
)

func (c *Client) CreateOnlineSale(ctx context.Context, litleReq LitleOnlineRequest) (*LitleOnlineResponse, error) {
	req, err := c.NewRequest(ctx, litleReq)

	response := &LitleOnlineResponse{}

	if err != nil {
		return response, err
	}

	err = c.Send(req, response)

	return response, err
}
