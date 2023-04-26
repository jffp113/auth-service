package data

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const DefaultDriverName = "postgres"

type Credentials struct {
	username string
	password string
}

type Client struct {
	ctx context.Context

	*gorm.DB

	host   string
	port   int
	dbname string

	credentials Credentials

	log *zap.SugaredLogger

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
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cli.host, cli.port, cli.credentials.username, cli.dbname, cli.credentials.password, "disable")

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: createLogger(cli.log),
	})

	if err != nil {
		return fmt.Errorf("openning connection to db: %w", err)
	}
	cli.DB = db

	if cli.debug {
		db.Debug()
	}

	return nil
}

func createLogger(log *zap.SugaredLogger) logger.Interface {
	if log == nil {
		return nil
	}

	return logger.New(zap.NewStdLog(log.Desugar()), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Error,
		Colorful:      false,
	})
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
	//return r.Client.Close()
	return nil
}
