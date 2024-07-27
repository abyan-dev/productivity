package main

import "github.com/abyan-dev/productivity/pkg/server"

func main() {
	srv := server.Server{}
	app := srv.New()
	srv.Run(app)
}
