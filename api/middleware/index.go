package middleware

import "github.com/schidstorm/ffmpeg-jobs/api/lib"

func Index() []lib.Middleware {
	return []lib.Middleware{
		Cors{},
	}
}
