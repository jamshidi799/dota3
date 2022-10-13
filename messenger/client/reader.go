package client

func (c Clients) ReadFromConnection(connectionId int, schema any) {
	conn := c[connectionId] // todo: remove disconnected client
	conn.receive(schema)
}
