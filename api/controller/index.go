package controller

import "github.com/schidstorm/ffmpeg-jobs/api/lib"

func Index() []lib.Controller {
	return []lib.Controller{
		&Job{},
		&Claim{},
	}
}
