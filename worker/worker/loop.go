package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/schidstorm/ffmpeg-jobs/worker/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Loop struct {
	config *config.Config
}

func NewLoop(config *config.Config) *Loop {
	return &Loop{config: config}
}

func (l *Loop) Run(applicationContext context.Context) error {
	claimUrl := fmt.Sprintf("%s/claim", l.config.ApiServerUrl)
	for {
		jobContext, cancel := context.WithCancel(context.Background())
		select {
		case <-applicationContext.Done():
			cancel()
			return nil
		case <-time.After(10 * time.Second):

		}

		resp, err := http.Post(claimUrl, "text/plain", bytes.NewBuffer([]byte("")))
		if err == nil && resp.StatusCode != 200 {
			logrus.Info(fmt.Errorf("no claims available (Status code %d)", resp.StatusCode))
			continue
		}
		if err != nil {
			logrus.Error(err)
			continue
		}

		bodyData, err := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			logrus.Error(err)
		}

		var job Job
		err = json.Unmarshal(bodyData, &job)
		if err != nil {
			logrus.Error(err)
		}

		logrus.Infof("claimed job %s", job.Data.InputFile)

		jobDone := make(chan error, 1)

		go func() {
			jobDone <- RunFfmpegJob(job, func(p float64) {
				updateProgress(l.config.ApiServerUrl, job, p)
			}, jobContext)
		}()

		select {
		case <-applicationContext.Done():
			cancel()
		case err = <-jobDone:
			cancel()
			if err == nil {
				if l.config.DeleteFinished {
					err = deleteJobInputFile(job)
				}
			} else {
				failJob(l.config.ApiServerUrl, job, err)
			}
		}

		if err != nil {
			logrus.Error(err)
		}
	}
}

func updateProgress(apiServerUrl string, job Job, progress float64) {
	putUrl := fmt.Sprintf("%s/job/%d", apiServerUrl, job.Data.ID)
	data := fmt.Sprintf("{\"Progress\": %f}", progress)
	req, err := http.NewRequest("PUT", putUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		logrus.Error(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error(err)
	}
}

func failJob(apiServerUrl string, job Job, err error) {
	putUrl := fmt.Sprintf("%s/job/%d", apiServerUrl, job.Data.ID)
	data := fmt.Sprintf("{\"Failed\": true, \"Claimed\": false, \"Error\": \"%s\"}", err.Error())
	req, err := http.NewRequest("PUT", putUrl, bytes.NewBuffer([]byte(data)))
	if err != nil {
		logrus.Error(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error(err)
	}
}

func deleteJobInputFile(job Job) error {
	return os.Remove(job.Data.InputFile)
}
