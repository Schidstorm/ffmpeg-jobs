package main

import (
	"github.com/schidstorm/ffmpeg-jobs/api/dependencies"
	"github.com/schidstorm/ffmpeg-jobs/api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	dependencies.InitCollection()
	srv := server.NewServer()
	server.InitializeServer(srv)
	logrus.Error(srv.Serve("0.0.0.0:8080"))
}
