package controller

import (
	"github.com/schidstorm/ffmpeg-jobs/api/dependencies"
	"github.com/schidstorm/ffmpeg-jobs/api/domain"
	"github.com/schidstorm/ffmpeg-jobs/api/lib"
	"net/url"
)

type Claim struct {
}

func (c Claim) Name() string {
	return "claim"
}

func (c Claim) GetHandler() lib.GetHandlerFunc {
	return nil
}

func (c Claim) PutHandler() (lib.PutHandlerFunc, interface{}) {
	return nil, nil
}

func (c Claim) PostHandler() (lib.PostHandlerFunc, interface{}) {
	return func(values url.Values, _ interface{}) (interface{}, error) {
		db := dependencies.Current.Database.DB()
		var job domain.Job
		err := db.First(&job, "started = ?", false).Update("started", true).Error
		if err != nil {
			return nil, err
		}

		return job, nil
	}, nil
}

func (c Claim) ListHandler() lib.ListHandlerFunc {
	return nil
}
