package controller

import (
	"github.com/schidstorm/ffmpeg-jobs/api/dependencies"
	"github.com/schidstorm/ffmpeg-jobs/api/domain"
	"github.com/schidstorm/ffmpeg-jobs/api/lib"
	"net/url"
)

type Job struct {
}

func (j Job) Name() string {
	return "job"
}

func (j Job) GetHandler() lib.GetHandlerFunc {
	return func(id int64, values url.Values) (interface{}, error) {
		db := dependencies.Current.Database.DB()
		job := &domain.Job{}
		result := db.First(&job, id)
		return job, result.Error
	}
}

func (j Job) PutHandler() (lib.PutHandlerFunc, interface{}) {
	requestData := &domain.Job{}
	return func(id int64, values url.Values, data interface{}) (interface{}, error) {
		db := dependencies.Current.Database.DB()
		job := &domain.Job{}
		result := db.First(&job, id).Updates(requestData)
		return job, result.Error
	}, requestData
}

func (j Job) PostHandler() (lib.PostHandlerFunc, interface{}) {
	requestData := &domain.Job{}
	return func(values url.Values, data interface{}) (interface{}, error) {
		db := dependencies.Current.Database.DB()
		result := db.Create(requestData)
		return requestData, result.Error
	}, requestData
}

func (j Job) ListHandler() lib.ListHandlerFunc {
	return func(values url.Values) (interface{}, error) {
		db := dependencies.Current.Database.DB()
		var result []domain.Job
		db.Find(&result)

		return result, db.Error
	}
}
