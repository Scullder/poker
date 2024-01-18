package websocket

func GetHandlers() *map[string]EventHandler {
	handlers := make(map[string]EventHandler)

	handlers["joinToGame"] = func(event Event, c *Client) error {
		return nil
	}

	return &handlers
}
