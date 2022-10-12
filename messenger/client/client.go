package client

type Client interface {
	write(event any) error
	read(schema any)
	GetId() int
}

type BaseClient struct {
	id       int
	username string
}

type Clients map[int]Client
