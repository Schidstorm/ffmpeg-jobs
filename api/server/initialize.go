package server

import "github.com/schidstorm/ffmpeg-jobs/api/controller"

func InitializeServer(server *Server) {
	for _, c := range controller.Index() {
		server.AddController(c)
	}
}
