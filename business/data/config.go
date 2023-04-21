package data

const (
	DefaultHost = "localhost:9080"
)

type Config func(*Client) error

func WithCredentials(username string, password string) Config {
	return func(client *Client) error {
		client.credentials = Credentials{
			username: username,
			password: password,
		}

		return nil
	}
}

func WithDbName(dbname string) Config {
	return func(client *Client) error {
		client.dbname = dbname
		return nil
	}
}

func WithHostAndPort(host string, port int) Config {
	return func(client *Client) error {
		client.host = host
		client.port = port
		return nil
	}
}

func WithDebug(debug bool) Config {
	return func(client *Client) error {
		client.debug = debug
		return nil
	}
}
