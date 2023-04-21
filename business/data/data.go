package data

import (
	"com.cross-join.crossviewer.authservice/business/data/ent"
	"context"
	"fmt"
	_ "github.com/lib/pq"
)

const DefaultDriverName = "postgres"

type Credentials struct {
	username string
	password string
}

type Client struct {
	ctx context.Context
	*ent.Client

	host   string
	port   int
	dbname string

	credentials Credentials

	debug bool
}

func New(ctx context.Context, confs ...Config) (Client, error) {
	var c Client

	c.ctx = ctx
	if err := applyConfigs(&c, confs); err != nil {
		return c, err
	}

	if err := startClient(&c); err != nil {
		return c, err
	}

	return c, nil
}

func startClient(cli *Client) error {
	var opts []ent.Option

	if cli.debug {
		opts = append(opts, ent.Debug())
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cli.host, cli.port, cli.credentials.username, cli.dbname, cli.credentials.password, "disable")

	entCli, err := ent.Open(DefaultDriverName, connStr, opts...)
	if err != nil {
		return fmt.Errorf("openning connection to db: %w", err)
	}

	cli.Client = entCli

	return nil
}

func applyConfigs(c *Client, confs []Config) error {
	for _, conf := range confs {
		err := conf(c)
		if err != nil {
			return fmt.Errorf("applying configs: %w", err)
		}
	}

	return nil
}

func (r Client) Close() error {
	return r.Client.Close()
}
