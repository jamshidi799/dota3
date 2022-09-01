package cmd

import "game/transport/http"

func Run() error {
	return http.StartServer()
}
