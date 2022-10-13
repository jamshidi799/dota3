package client

func (c Clients) BroadcastEvent(event any) {
	for _, client := range c {
		if err := client.send(event); err != nil {
			// todo: remove disconnected client
		}
	}
}

func (c Clients) BroadcastEventToOther(exceptionPlayerId int, event any) {
	for id, client := range c {
		if id != exceptionPlayerId {
			if err := client.send(event); err != nil {
				// todo: remove disconnected client
			}
		}
	}
}

func (c Clients) SendEventToConnection(connectionId int, event any) error {
	conn := c[connectionId] // todo: remove disconnected client
	return conn.send(event)
}
