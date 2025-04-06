package main

import (
	"encoding/json"
	"errors"
)

type Client struct {
	Endpoint string `json:"endpoint"`
	Id       string `json:"id"`
}

type Response interface {
	Serialize() ([]byte, error)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (response *ErrorResponse) Serialize() ([]byte, error) {
	return json.Marshal(response)
}

var clients = make(map[string]*Client)

func getClient(Endpoint string) (*Client, error) {
	original, exists := clients[Endpoint]
	if !exists {
		return nil, errors.New("could not find Client")
	}

	return original, nil
}

func addClient(c *Client) error {
	original, exists := clients[c.Endpoint]
	if exists && original.Id != c.Id {
		return errors.New("endpoint already has a registered client")
	}

	clients[c.Endpoint] = c
	return nil
}
