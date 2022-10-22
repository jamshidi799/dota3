package client

type Client interface {
	send(event any) error
	receive(schema any)
	GetId() int
	GetUsername() string
}

type BaseClient struct {
	id       int
	username string
}

type Clients map[int]Client
