package main

import (
	"fmt"

	"github.com/ccp-clientlibrary-go/ccp"
)

type Config struct {
	Username string
	Password string
	Base_url string
}

func (c *Config) Client() *ccp.Client {

	client := ccp.NewClient(c.Username, c.Password, c.Base_url)

	err := client.Login(client)

	if err != nil {
		fmt.Println(err)
	}

	return client
}
