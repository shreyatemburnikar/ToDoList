package main

import (
	"context"

	"github.com/urfave/cli/v2"
)

func Sync(c *cli.Context) error {
	client := GetClient(c)

	err := client.Sync(context.Background())
	if err != nil {
		return err
	}
	return WriteCache(cachePath, client.Store)
}
